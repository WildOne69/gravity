package tsdb

import (
	"context"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"beryju.io/gravity/pkg/extconfig"
	"beryju.io/gravity/pkg/roles"
	debugTypes "beryju.io/gravity/pkg/roles/debug/types"
	"beryju.io/gravity/pkg/roles/tsdb/types"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Role struct {
	log *log.Entry
	i   roles.Instance
	ctx context.Context
	m   map[string]types.Metric
	ms  sync.RWMutex
}

func New(instance roles.Instance) *Role {
	r := &Role{
		log: instance.Log(),
		i:   instance,
		m:   make(map[string]types.Metric),
		ms:  sync.RWMutex{},
	}
	r.i.AddEventListener(debugTypes.EventTopicDebugMuxSetup, func(ev *roles.Event) {
		mux := ev.Payload.Data["mux"].(*mux.Router)
		mux.HandleFunc("/debug/tsdb/write", func(w http.ResponseWriter, re *http.Request) {
			r.i.DispatchEvent(types.EventTopicTSDBWrite, roles.NewEvent(
				re.Context(),
				map[string]interface{}{},
			))
		})
	})
	r.i.AddEventListener(types.EventTopicTSDBBeforeWrite, func(ev *roles.Event) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		r.setMetric(
			r.i.KV().Key(
				types.KeySystem,
				"memory",
			).String(),
			types.Metric{
				Value: int(m.Sys),
			},
		)
	})
	return r
}

func (r *Role) setMetric(key string, value types.Metric) {
	r.ms.Lock()
	defer r.ms.Unlock()
	r.log.WithField("key", key).Trace("tsdb set")
	r.m[key] = value
}

func (r *Role) Start(ctx context.Context, config []byte) error {
	r.ctx = ctx
	r.i.AddEventListener(types.EventTopicTSDBWrite, func(ev *roles.Event) {
		r.write()
	})
	r.i.AddEventListener(types.EventTopicTSDBSet, func(ev *roles.Event) {
		key := ev.Payload.Data["key"].(string)
		r.log.WithField("key", key).Trace("tsdb set")
		r.setMetric(key, ev.Payload.Data["value"].(types.Metric))
	})
	r.i.AddEventListener(types.EventTopicTSDBInc, func(ev *roles.Event) {
		r.ms.Lock()
		defer r.ms.Unlock()
		key := ev.Payload.Data["key"].(string)
		r.log.WithField("key", key).Trace("tsdb inc")
		val, ok := r.m[key]
		if !ok {
			val = ev.Payload.Data["default"].(types.Metric)
		}
		val.Value += 1
		r.m[key] = val
	})
	go func() {
		for {
			select {
			case <-r.ctx.Done():
				return
			default:
				r.write()
				time.Sleep(30 * time.Second)
			}
		}
	}()
	return nil
}

func (r *Role) write() {
	r.log.Trace("writing metrics")
	r.i.DispatchEvent(types.EventTopicTSDBBeforeWrite, roles.NewEvent(r.ctx, map[string]interface{}{}))
	r.ms.RLock()
	defer r.ms.RUnlock()
	// Don't bother granting a lease if we don't have any metrics
	if len(r.m) < 1 {
		return
	}
	lease, err := r.i.KV().Grant(r.ctx, 60*30)
	if err != nil {
		r.log.WithError(err).Warning("failed to grant lease, skipping write")
		return
	}
	for rkey, value := range r.m {
		key := r.i.KV().Key(
			types.KeyRole,
			rkey,
			extconfig.Get().Instance.Identifier,
			strconv.FormatInt(time.Now().Unix(), 10),
		).String()
		_, err := r.i.KV().Put(
			r.ctx,
			key,
			strconv.Itoa(value.Value),
			clientv3.WithLease(lease.ID),
		)
		if err != nil {
			r.log.WithError(err).WithField("key", key).Warning("failed to put value")
		}
		if value.ResetOnWrite {
			value.Value = 0
			r.m[rkey] = value
		}
	}
}

func (r *Role) Stop() {
}
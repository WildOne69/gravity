package dhcp

import (
	"time"

	"beryju.io/gravity/pkg/roles/dhcp/types"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (r *DHCPRole) startWatchLeases() {
	evtHandler := func(ev *clientv3.Event) {
		r.leasesSync.Lock()
		defer r.leasesSync.Unlock()
		if ev.Type == clientv3.EventTypeDelete {
			delete(r.leases, string(ev.Kv.Key))
		} else {
			rec, err := r.leaseFromKV(ev.Kv)
			if err != nil {
				r.log.WithError(err).Warning("failed to parse lease")
				return
			}
			r.leases[string(ev.Kv.Key)] = rec
		}
	}

	prefix := r.i.KV().Key(
		types.KeyRole,
		types.KeyLeases,
	).Prefix(true).String()

	leases, err := r.i.KV().Get(r.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		r.log.WithError(err).Warning("failed to list initial leases")
		time.Sleep(5 * time.Second)
		r.startWatchLeases()
		return
	}
	for _, lease := range leases.Kvs {
		evtHandler(&clientv3.Event{
			Type: mvccpb.PUT,
			Kv:   lease,
		})
	}

	watchChan := r.i.KV().Watch(
		r.ctx,
		prefix,
		clientv3.WithPrefix(),
		clientv3.WithProgressNotify(),
	)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			go evtHandler(event)
		}
	}
}

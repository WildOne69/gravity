package api

import (
	"context"
	"encoding/base64"
	"encoding/json"

	instanceTypes "beryju.io/gravity/pkg/instance/types"
	"beryju.io/gravity/pkg/roles/api/types"
	"github.com/gorilla/securecookie"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

type RoleConfig struct {
	Port         int32             `json:"port"`
	CookieSecret string            `json:"cookieSecret"`
	OIDC         *types.OIDCConfig `json:"oidc"`
}

func (r *Role) checkCookieSecret(cfg *RoleConfig, fallback string) {
	if cfg.CookieSecret != "" {
		return
	}
	cfg.CookieSecret = fallback
	r.log.Info("cookie secret not in config, generating one")
	go func(cfg *RoleConfig) {
		jc, err := json.Marshal(cfg)
		if err != nil {
			r.log.WithError(err).Warning("failed to json parse config")
			return
		}
		_, err = r.i.KV().Put(
			context.Background(),
			r.i.KV().Key(
				instanceTypes.KeyInstance,
				instanceTypes.KeyRole,
				types.KeyRole,
			).String(),
			string(jc),
		)
		if err != nil {
			r.log.WithError(err).Warning("failed to save config")
			return
		}
	}(cfg)
}

func (r *Role) decodeRoleConfig(raw []byte) *RoleConfig {
	fallbackSecret := base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32))
	def := RoleConfig{
		Port: 8008,
	}
	if len(raw) < 1 {
		r.checkCookieSecret(&def, fallbackSecret)
		return &def
	}
	err := json.Unmarshal(raw, &def)
	if err != nil {
		r.log.WithError(err).Warning("failed to decode role config")
	}
	r.checkCookieSecret(&def, fallbackSecret)
	return &def
}

type APIRoleAPIConfigOutput struct {
	Config RoleConfig `json:"config" required:"true"`
}

func (r *Role) APIRoleConfigGet() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *APIRoleAPIConfigOutput) error {
		output.Config = *r.cfg
		return nil
	})
	u.SetName("api.get_role_config")
	u.SetTitle("API role config")
	u.SetTags("roles/api")
	return u
}

type APIRoleAPIConfigInput struct {
	Config RoleConfig `json:"config" required:"true"`
}

func (r *Role) APIRoleConfigPut() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input APIRoleAPIConfigInput, output *struct{}) error {
		jc, err := json.Marshal(input.Config)
		if err != nil {
			return status.Wrap(err, status.InvalidArgument)
		}
		_, err = r.i.KV().Put(
			ctx,
			r.i.KV().Key(
				instanceTypes.KeyInstance,
				instanceTypes.KeyRole,
				types.KeyRole,
			).String(),
			string(jc),
		)
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		return nil
	})
	u.SetName("api.put_role_config")
	u.SetTitle("API role config")
	u.SetTags("roles/api")
	u.SetExpectedErrors(status.InvalidArgument, status.Internal)
	return u
}
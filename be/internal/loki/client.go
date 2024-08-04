package loki

import "github.com/grafana/loki/v3/pkg/logcli/client"

func NewClient(config *Config) client.Client {
	return &client.DefaultClient{
		Username: config.User,
		Password: config.Password,
		Address:  config.Endpoint,
		Retries:  0,
		BackoffConfig: client.BackoffConfig{
			MaxBackoff: 60,
			MinBackoff: 10,
		},
	}
}

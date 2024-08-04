package loki

import (
	"fmt"

	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
	"github.com/prometheus/common/model"
)

// Loki Metrics-queries (https://grafana.com/docs/loki/latest/query/metric_queries/) use Prometheus models

func mapSampleToLoginAttempt(sample model.Sample) (*stats.LoginAttempt, error) {
	if user, ok := sample.Metric["user"]; ok {
		return &stats.LoginAttempt{
			Username: string(user),
			Count:    int(sample.Value),
		}, nil
	}

	return nil, fmt.Errorf("loki: query result did not have an 'user' label. metrics: %v", sample.Metric)
}

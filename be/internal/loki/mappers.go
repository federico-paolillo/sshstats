package loki

import (
	"fmt"

	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
	"github.com/prometheus/common/model"
)

func mapPrometheusSampleToOurs(promSample model.Sample) *RawSample {
	labels := make(SampleLabels, 0)

	for k, v := range promSample.Metric {
		labels[string(k)] = string(v)
	}

	return &RawSample{
		Labels: labels,
		Value:  float64(promSample.Value),
	}
}

func mapRawSampleToLoginAttempt(sample *RawSample) (*stats.LoginAttempt, error) {
	if user, ok := sample.Labels["user"]; ok {
		return &stats.LoginAttempt{
			Username: string(user),
			Count:    int(sample.Value),
		}, nil
	}

	return nil, fmt.Errorf("loki: sample did not have an 'user' label. sample: %v", sample)
}

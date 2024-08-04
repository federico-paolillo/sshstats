package loki

import (
	"fmt"
	"time"

	"github.com/grafana/loki/v3/pkg/logproto"
	"github.com/grafana/loki/v3/pkg/logcli/client"
	"github.com/grafana/loki/v3/pkg/loghttp"
	"github.com/prometheus/common/model"
)

type Connector interface {
	Query() (loghttp.Vector, error)
}

type Labels = map[string]string

type RawSample struct {
	labels Labels
	value  float64
}

type connector struct {
	client client.Client
}

func (c *connector) MetricQuery(logql string) ([]RawSample, error) {
	resp, err := c.client.Query(
		logql,
		100,
		time.Now(),
		logproto.BACKWARD,
		false,
	)

	if err != nil {
		return nil, fmt.Errorf("loki: could not run query '%s'. %v", logql, err)
	}

	if resp.Status != "success" {
		return nil, fmt.Errorf("loki: query was not successfull. status: %s", resp.Status)
	}

	respType := resp.Data.ResultType

	if respType != loghttp.ResultTypeVector {
		return nil, fmt.Errorf("loki: query returned unexpected result type. type: %s", respType)
	}

	querySamplesRetrieved := resp.Data.Result.(loghttp.Vector)

	rawSamples := make([]RawSample, 0, len(querySamplesRetrieved))

	for _, querySample := range querySamplesRetrieved {
		rawSamples = append(rawSamples, *mapPrometheusSampleToOurs(querySample))
	}

	return rawSamples, nil
}

func mapPrometheusSampleToOurs(promSample model.Sample) *RawSample {
	labels := make(Labels, 0)

	for k, v := range promSample.Metric {
		labels[string(k)] = string(v)
	}

	return &RawSample{
		labels: labels,
		value:  float64(promSample.Value),
	}
}

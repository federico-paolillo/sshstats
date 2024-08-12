package loki

import (
	"fmt"
	"time"

	"github.com/grafana/loki/v3/pkg/logcli/client"
	"github.com/grafana/loki/v3/pkg/loghttp"
	"github.com/grafana/loki/v3/pkg/logproto"
)

type LogcliConnector struct {
	client client.Client
}

func NewLokiConnector(client client.Client) *LogcliConnector {
	return &LogcliConnector{client}
}

func (c *LogcliConnector) MetricQuery(logql string) ([]*RawSample, error) {
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

var _ Connector = (*LogcliConnector)(nil)

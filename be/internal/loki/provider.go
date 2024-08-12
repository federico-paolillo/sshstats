package loki

import (
	"fmt"
	"log"
	"slices"

	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
)

const top15QueryTemplate = "topk(15, sum by(user) (count_over_time({node=\"%s\"} | json [24h])))"

type Connector interface {
	MetricQuery(logql string) ([]*RawSample, error)
}

type Provider struct {
	connector Connector
	log       *log.Logger
}

func NewProvider(log *log.Logger, connector Connector) *Provider {
	return &Provider{connector, log}
}

func (p *Provider) Top15LoginAttempts(nodeName string) (stats.Attempts, error) {
	top15QueryWithParams := fmt.Sprintf(top15QueryTemplate, nodeName)

	p.log.Printf("querying loki for node '%s'. query: '%s'", nodeName, top15QueryWithParams)

	rawSamples, err := p.connector.MetricQuery(top15QueryWithParams)

	if err != nil {
		return nil, fmt.Errorf("loki: could not query node %s. %w", nodeName, err)
	}

	p.log.Printf("got back from loki %d samples for node '%s'", len(rawSamples), nodeName)

	result := make(stats.Attempts, 0, len(rawSamples))

	for _, rawSample := range rawSamples {
		loginAttempt, err := mapRawSampleToLoginAttempt(rawSample)

		if err != nil {
			return nil, fmt.Errorf("loki: could not map sample %v. %w", rawSample, err)
		}

		result = append(result, loginAttempt)
	}

	slices.SortFunc(
		result,
		func(a *stats.LoginAttempt, b *stats.LoginAttempt) int {
			return a.Count - b.Count
		})

	return result, nil
}

func (p *Provider) Last10Attackers() (stats.Attackers, error) {
	panic("not implemented")
}

var _ stats.Provider = (*Provider)(nil)

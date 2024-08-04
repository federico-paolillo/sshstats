package loki

import (
	"errors"
	"fmt"
	"log"

	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
	"github.com/grafana/loki/v3/pkg/logcli/client"
)

var top15Query = "topk(15, sum by(user) (count_over_time({node=\"%s\"} | json [24h])))"
var last10Query = "TODO"

type statistician struct {
	log    *log.Logger
	client client.Client
}

func NewStatistician(log *log.Logger, client client.Client) stats.Statistician {
	return &statistician{
		log,
		client,
	}
}

func (s *statistician) Top15LoginAttempts(nodeName string) (stats.Attempts, error) {
	query := fmt.Sprintf(top15Query, nodeName)

	attemptsMapped := make(stats.Attempts, 0, len(lokiQueryResults))

	var mappingErrors []error = make([]error, 0)

	for _, sample := range lokiQueryResults {
		attemptMapped, err := mapSampleToLoginAttempt(sample)

		if err != nil {
			mappingErrors = append(mappingErrors, err)
			continue
		}

		attemptsMapped = append(attemptsMapped, attemptMapped)
	}

	if len(mappingErrors) > 0 {
		return nil, fmt.Errorf("loki: could not map some query result entries. %w", errors.Join(mappingErrors...))
	}

	return attemptsMapped, nil
}

func (s *statistician) Last10Attackers() (stats.Attackers, error) {
	panic("unimplemented")
}

var (
	_ stats.Statistician = (*statistician)(nil)
)

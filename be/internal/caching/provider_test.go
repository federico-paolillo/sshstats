package caching_test

import (
	"log"
	"testing"

	"github.com/federico-paolillo/ssh-attempts/internal/caching"
	"github.com/federico-paolillo/ssh-attempts/internal/fakes"
	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
)

func TestCachingShortCircuitsCallsToUnderlyingProvider(t *testing.T) {
	log := log.Default()

	fp := &fakes.Provider{
		Data: map[string]stats.Attempts{
			"nodename-1": {
				&stats.LoginAttempt{
					Username: "pippo",
					Count:    12,
				},
			},
		},
		T:     t,
		Calls: 0,
	}

	cp := caching.NewProvider(log, fp)

	//Warmup plus extra calls

	_, _ = cp.Top15LoginAttempts("nodename-1")
	_, _ = cp.Top15LoginAttempts("nodename-1")
	_, _ = cp.Top15LoginAttempts("nodename-1")

	if fp.Calls != 1 {
		t.Fatalf(
			"expected provider to be called just once. was called %d",
			fp.Calls,
		)
	}
}

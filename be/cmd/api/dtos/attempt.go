package dtos

import "github.com/federico-paolillo/ssh-attempts/pkg/stats"

type LoginAttempt struct {
	Username string `json:"username"`
	Count    int    `json:"count"`
}

type LoginAttempts struct {
	Attempts []*LoginAttempt `json:"attempts"`
}

func MapAttemptsToDto(attemptsToMap []*stats.LoginAttempt) *LoginAttempts {
	attemptsMapped := make([]*LoginAttempt, 0, len(attemptsToMap))

	for _, attemptToMap := range attemptsToMap {
		attemptMapped := &LoginAttempt{
			Username: attemptToMap.Username,
			Count:    attemptToMap.Count,
		}

		attemptsMapped = append(attemptsMapped, attemptMapped)
	}

	return &LoginAttempts{
		Attempts: attemptsMapped,
	}
}

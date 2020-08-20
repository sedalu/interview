package store

import (
	"context"

	"github.com/sedalu/interview/internal/model"
)

type ChallengeList struct {
	Challenges []model.Challenge `json:"challenges"`
	Count      int               `json:"count"` // The number of challenges in the list.
	Total      int               `json:"total"` // The total number of challenges.
	Limit      int               `json:"limit"`
	Offset     int               `json:"offset"`
}

type Store interface {
	ListChallenges(ctx context.Context, limit, offset int) (*ChallengeList, error)
}

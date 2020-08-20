package pgsql

import (
	"context"

	"github.com/sedalu/interview/internal/model"
	"github.com/sedalu/interview/internal/store"
)

// ListChallenges returns a ChallengeList containing a list of up to limit Challenges, starting at offset.
func (s *Store) ListChallenges(ctx context.Context, limit, offset int) (*store.ChallengeList, error) {
	// TODO: Write SQL query to count the total challenges in the table.
	// TODO: Write SQL query to get up to limit challenges, starting at offset.
	const (
		countSQL = ""
		getSQL   = ""
	)

	list := store.ChallengeList{
		Challenges: make([]model.Challenge, 0, limit),
		Limit:      limit,
		Offset:     offset,
	}

	// Get the total count of challenges.
	row := s.QueryRowContext(ctx, countSQL)
	if err := row.Scan(&list.Total); err != nil {
		return nil, err
	}

	// Get the challenges.
	rows, err := s.QueryContext(ctx, getSQL, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c model.Challenge

		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return nil, err
		}

		list.Challenges = append(list.Challenges, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Set the count of included challenges.
	list.Count = len(list.Challenges)

	return &list, nil
}

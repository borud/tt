package model

import ttv1 "github.com/borud/tt/pkg/tt/v1"

// Snippet is a short note that says what has been done on a given day.
type Snippet struct {
	ID       IDType `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	TS       int64  `json:"ts" db:"ts"`
	Contents string `json:"contents" db:"contents"`
}

// Proto returns proto representation.
func (s Snippet) Proto() *ttv1.Snippet {
	return &ttv1.Snippet{
		Id:       uint64(s.ID),
		Username: s.Username,
		Ts:       s.TS,
		Contents: s.Contents,
	}
}

// SnippetFromProto returns Snippet corresponding to ttv1.Snippet
func SnippetFromProto(s *ttv1.Snippet) Snippet {
	return Snippet{
		ID:       IDType(s.Id),
		Username: s.Username,
		TS:       s.Ts,
		Contents: s.Contents,
	}
}

package model

import (
	ttv1 "github.com/borud/tt/pkg/tt/v1"
)

// Work represents work that has been done on a project.
type Work struct {
	ID       IDType `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Project  string `json:"project" db:"project"`
	TS       int64  `json:"ts" db:"ts"`
	Duration int64  `json:"duration" db:"duration"`
	Descr    string `json:"descr" db:"descr"`
}

// Proto returns proto representation.
func (w Work) Proto() *ttv1.Work {
	return &ttv1.Work{
		Id:       uint64(w.ID),
		Username: w.Username,
		Project:  w.Project,
		Ts:       w.TS,
		Duration: w.Duration,
		Descr:    w.Descr,
	}
}

// WorkFromProto returns Work corresponding to ttv1.Work
func WorkFromProto(w *ttv1.Work) Work {
	return Work{
		ID:       IDType(w.Id),
		Username: w.Username,
		Project:  w.Project,
		TS:       w.Ts,
		Duration: w.Duration,
		Descr:    w.Descr,
	}
}

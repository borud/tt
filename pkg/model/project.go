package model

import ttv1 "github.com/borud/tt/pkg/tt/v1"

// Project represents a project.
type Project struct {
	Name  string `json:"name" db:"name"`
	Descr string `json:"descr" db:"descr"`
}

// Proto returns proto representation.
func (p Project) Proto() *ttv1.Project {
	return &ttv1.Project{
		Name:  p.Name,
		Descr: p.Descr,
	}
}

// ProjectFromProto returns Project corresponding to ttv1.Project
func ProjectFromProto(p *ttv1.Project) Project {
	return Project{
		Name:  p.Name,
		Descr: p.Descr,
	}
}

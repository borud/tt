package store

import "github.com/borud/tt/pkg/model"

// Store interface for tt
type Store interface {
	Close() error

	AddUser(user model.User) error
	GetUser(username string) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(username string) error
	ListUsers() ([]model.User, error)

	AddProject(project model.Project) error
	GetProject(name string) (model.Project, error)
	UpdateProject(project model.Project) error
	DeleteProject(name string) error
	ListProjects() ([]model.Project, error)

	AddWork(work model.Work) error
	GetWork(id model.IDType) (model.Work, error)
	UpdateWork(work model.Work) error
	DeleteWork(id model.IDType) error
	ListWork(from int64, until int64) ([]model.Work, error)
	ListWorkForUser(username string, from int64, until int64) ([]model.Work, error)
	ListWorkForProject(project string, from int64, until int64) ([]model.Work, error)

	AddSnippet(snippet model.Snippet) error
	GetSnippet(id model.IDType) (model.Snippet, error)
	UpdateSnippet(snippet model.Snippet) error
	DeleteSnippet(id model.IDType) error
	ListSnippets(from int64, until int64) ([]model.Snippet, error)
	ListSnippetsForUser(username string, from int64, until int64) ([]model.Snippet, error)
}

package sqlitestore

import "github.com/borud/tt/pkg/model"

func (s *sqliteStore) AddProject(project model.Project) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec("INSERT INTO projects(name,descr) VALUES(:name,:descr)", project)
	return err
}

func (s *sqliteStore) GetProject(name string) (model.Project, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var project model.Project
	return project, s.db.QueryRowx("SELECT * FROM projects WHERE name = ?", name).StructScan(&project)
}

func (s *sqliteStore) UpdateProject(project model.Project) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(`
		UPDATE projects
		SET 
			descr = :descr,
		WHERE
			name = :name`, project)
	return err
}

func (s *sqliteStore) DeleteProject(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec("DELETE FROM projects WHERE name = ?", name)
	return err
}

func (s *sqliteStore) ListProjects() ([]model.Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var projects []model.Project
	return projects, s.db.Select(&projects, "SELECT * FROM projects ORDER BY name")
}

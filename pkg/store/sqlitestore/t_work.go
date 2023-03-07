package sqlitestore

import "github.com/borud/tt/pkg/model"

func (s *sqliteStore) AddWork(work model.Work) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec(`
		INSERT INTO work (username,project,ts,duration,descr)
		VALUES (:username,:project,:ts,:duration,:descr)`))
}

func (s *sqliteStore) GetWork(id model.IDType) (model.Work, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var work model.Work
	return work, s.db.QueryRowx("SELECT * FROM work WHERE id = ?", id).StructScan(&work)
}

func (s *sqliteStore) UpdateWork(work model.Work) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec(`
		UPDATE work
		SET
			project = :project,
			ts = :ts,
			duration = :duration,
			descr = :descr
		WHERE
			id = :id`))
}
func (s *sqliteStore) DeleteWork(id model.IDType) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec("DELETE FROM work WHERE id = ?", id))
}

func (s *sqliteStore) ListWork(from int64, until int64) ([]model.Work, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var work []model.Work
	return work, s.db.Select(&work, `
		SELECT * FROM work 
		WHERE ts >= ? AND ts < ?
		ORDER BY ts`, from, until)
}

func (s *sqliteStore) ListWorkForUser(username string, from int64, until int64) ([]model.Work, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var work []model.Work
	return work, s.db.Select(&work, `
		SELECT * FROM work 
		WHERE username = ? 
			AND ts >= ?
			AND ts < ?
		ORDER BY ts`, from, until, username)
}
func (s *sqliteStore) ListWorkForProject(project string, from int64, until int64) ([]model.Work, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var work []model.Work
	return work, s.db.Select(&work, `
		SELECT * FROM work 
		WHERE project = ? 
			AND ts >= ?
			AND ts < ?
		ORDER BY ts`, from, until, project)
}

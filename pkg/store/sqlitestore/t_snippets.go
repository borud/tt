package sqlitestore

import "github.com/borud/tt/pkg/model"

func (s *sqliteStore) AddSnippet(snippet model.Snippet) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec(`
		INSERT INTO snippets (username, ts, contents)
		VALUES(:username, :ts, :contents)`))
}

func (s *sqliteStore) GetSnippet(id model.IDType) (model.Snippet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var snippet model.Snippet
	return snippet, s.db.QueryRowx("SELECT * FROM snippets WHERE id = ?", id).StructScan(&snippet)
}

func (s *sqliteStore) UpdateSnippet(snippet model.Snippet) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec(`
		UPDATE snippets
		SET
			ts = :ts,
			contents = :contents
		WHERE
			id = :id`))
}

func (s *sqliteStore) DeleteSnippet(id model.IDType) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec("DELETE FROM snippets WHERE id = ?", id))
}

func (s *sqliteStore) ListSnippets(from int64, until int64) ([]model.Snippet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var snippets []model.Snippet
	return snippets, s.db.Select(&snippets, `
		SELECT * FROM snippets
		WHERE ts >= ? AND ts < ?
		ORDER BY ts`, from, until)
}

func (s *sqliteStore) ListSnippetsForUser(username string, from int64, until int64) ([]model.Snippet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var snippets []model.Snippet
	return snippets, s.db.Select(&snippets, `
		SELECT * FROM snippets
		WHERE 
			username = ?
			AND ts >= ? 
			AND ts < ?
		ORDER BY ts`, username, from, until)
}

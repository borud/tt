package sqlitestore

import "github.com/borud/tt/pkg/model"

func (s *sqliteStore) AddUser(user model.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec(`
		INSERT INTO users(username,password,email,phone) 
		VALUES(:username,:password,:email,:phone)`, user))
}

func (s *sqliteStore) GetUser(username string) (model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var user model.User
	return user, s.db.QueryRowx("SELECT * FROM users WHERE username = ?", username).StructScan(&user)
}

func (s *sqliteStore) UpdateUser(user model.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec(`
		UPDATE users
		SET 
			password = :password,
			email = :email,
			phone = :phone
		WHERE
			username = :username`, user))

}

func (s *sqliteStore) DeleteUser(username string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return CheckForZeroRowsAffected(s.db.Exec("DELETE FROM users WHERE username = ?", username))
}

func (s *sqliteStore) ListUsers() ([]model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var users []model.User
	return users, s.db.Select(&users, "SELECT * FROM users ORDER BY username")
}

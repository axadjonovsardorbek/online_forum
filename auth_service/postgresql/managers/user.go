package managers

import (
	"auth-service/models"
	"database/sql"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Conn: db}
}

func (m *UserManager) Register(req models.RegisterReq) error {
	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	var id string
	err := m.Conn.QueryRow(query, req.ID, req.Username, req.Email, req.Password).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserManager) GetByID(req models.GetProfileByIdReq) (*models.GetProfileByIdResp, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	user := &models.GetProfileByIdResp{}
	err := m.Conn.QueryRow(query, req.ID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserManager) Profile(req models.GetProfileReq) (*models.GetProfileResp, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	user := &models.GetProfileResp{}
	err := m.Conn.QueryRow(query, req.Email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserManager) EmailExists(email string) (bool, error) {
	var exists bool
	err := m.Conn.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	return exists, err
}

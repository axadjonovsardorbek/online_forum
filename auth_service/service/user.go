package service

import (
	"auth-service/models"
	"auth-service/postgresql/managers"
	"database/sql"

	"github.com/google/uuid"
)

type UserService struct {
	UM managers.UserManager
}

func NewUserService(conn *sql.DB) *UserService {
	return &UserService{UM: *managers.NewUserManager(conn)}
}

func (u *UserService) Register(req *models.RegisterReq) error {
	req.ID = uuid.NewString()
	return u.UM.Register(*req)
}

// func (u *UserService) Login(req *models.LoginReq) bool {
// 	return false

// }

func (u *UserService) GetProfile(req *models.GetProfileReq) (*models.GetProfileResp, error) {
	return u.UM.Profile(*req)
}

func (u *UserService) GetByID(req *models.GetProfileByIdReq) (*models.GetProfileByIdResp, error) {
	return u.UM.GetByID(*req)
}

func (u *UserService) EmailExists(email string) (bool, error) {
	return u.UM.EmailExists(email)
}

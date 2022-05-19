package service

import (
	"fmt"
)

type AuthGetQuery struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (s *Service) CheckAuth(param *AuthGetQuery) (uint, error) {
	auth, err := s.Dao.GetAuth(param.Username, param.Password)
	if err != nil {
		return 0, err
	}
	if auth.ID > 0 {
		return auth.ID, nil
	}
	return 0, fmt.Errorf("%v", "auth info does not exist")
}

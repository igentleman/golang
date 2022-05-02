package service

import (
	"fmt"
)

type AuthGetQuery struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (s *Service) CheckAuth(param *AuthGetQuery) error {
	auth, err := s.Dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return fmt.Errorf("%v", "auth info does not exist")
}

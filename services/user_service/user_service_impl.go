package user_service

import (
	"fmt"
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/generic_utils/validate_utils"
	"github.com/three-kinds/user-center/utils/service_utils/password_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"strings"
	"time"
)

type UserServiceImpl struct {
	dao daos.IUserDAO
}

func (s *UserServiceImpl) checkUserIsActive(user *bo.UserBO) error {
	if !user.IsActive {
		return se.AccountDisabledError()
	}
	return nil
}

func (s *UserServiceImpl) loginWithPassword(user *bo.UserBO, password string) (*bo.UserBO, error) {
	ok := password_utils.IsSamePassword(password, user.Password)
	if !ok {
		return nil, se.ValidationError("incorrect account or password", se.Cause(fmt.Sprintf("username %s incorrect password", user.Username)))
	}

	err := s.checkUserIsActive(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) LoginByAccountAndPassword(account string, password string) (*bo.UserBO, error) {
	if validate_utils.IsEmail(account) {
		user, err := s.dao.GetUserByEmail(account)
		if err == nil {
			return s.loginWithPassword(user, password)
		}
	}

	if validate_utils.IsPhoneNumber(account) {
		user, err := s.dao.GetUserByPhoneNumber(account)
		if err == nil {
			return s.loginWithPassword(user, password)
		}
	}

	user, err := s.dao.GetUserByUsername(account)
	if err == nil {
		return s.loginWithPassword(user, password)
	}

	return nil, se.ValidationError("incorrect account or password", se.Cause(fmt.Sprintf("account %s mismatch", account)))
}

func (s *UserServiceImpl) GetActiveUserByID(userID int64) (*bo.UserBO, error) {
	user, err := s.dao.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	err = s.checkUserIsActive(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetActiveUserByEmail(email string) (*bo.UserBO, error) {
	user, err := s.dao.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = s.checkUserIsActive(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) RegisterUserByEmailPassword(email string, password string) error {
	oldUser, err := s.dao.GetUserByEmail(email)
	if err == nil && oldUser != nil {
		return se.ValidationError("this email has been registered, please login directly")
	}

	hashedPassword, err := password_utils.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = s.dao.CreateUser(&bo.CreateUserBO{
		Email:       email,
		Password:    hashedPassword,
		Username:    strings.Split(email, "@")[0],
		IsSuperuser: false,
	}, int64(initializers.SnowflakeNode.Generate()), time.Now())
	return err
}

func (s *UserServiceImpl) ResetPassword(id int64, password string) error {
	hashedPassword, err := password_utils.HashPassword(password)
	if err != nil {
		return err
	}

	return s.dao.UpdatePassword(id, hashedPassword)
}

func NewUserServiceImpl(dao daos.IUserDAO) *UserServiceImpl {
	return &UserServiceImpl{dao: dao}
}

package daos

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services/user_management_service"
	"github.com/three-kinds/user-center/utils/frame_utils/test_utils"
	"github.com/three-kinds/user-center/utils/generic_utils/gorm_addons"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
	"time"
)

func init() {
	test_utils.InitOnTestDAO(&models.User{})
}

func TestUserDAOImpl_CreateUser(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})

	dao := NewUserDAOImpl(initializers.DB)
	createUserBo := &user_management_service.CreateUserBO{
		Email:       "xx@xx.cpm",
		Username:    "xx",
		Password:    "password",
		IsSuperuser: false,
	}
	_, err := dao.CreateUser(createUserBo, 1, time.Now())
	assert.Equal(t, nil, err)

	total, err := dao.Count(nil, nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), total)

	// duplicate error
	_, err = dao.CreateUser(createUserBo, 2, time.Now())
	assert.NotNil(t, err)
	assert.Regexp(t, "duplicate key", err.Error())
}

func TestUserDAOImpl_WithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := gorm_addons.NewMockIDB(ctrl)
	dao := NewUserDAOImpl(dbMock)
	errDB := &gorm.DB{Error: errors.New("mock error")}

	// create
	dbMock.EXPECT().Create(gomock.Any()).Return(errDB)
	createUserBo := &user_management_service.CreateUserBO{
		Email:       "xx@xx.cpm",
		Username:    "xx",
		Password:    "password",
		IsSuperuser: false,
	}
	_, err := dao.CreateUser(createUserBo, 1, time.Now())
	assert.NotNil(t, err)
	assert.Regexp(t, "create user error", err.Error())
}

func TestUserDAOImpl_FailedWithShit(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	dao := NewUserDAOImpl(initializers.DB)
	newUser, err := dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		Password:    "password",
		IsSuperuser: true,
	}, 1, time.Now())

	_ = initializers.DB.AddError(errors.New("shit error"))
	defer func() {
		initializers.DB.Error = nil
	}()

	// Count
	_, err = dao.Count(nil, nil)
	assert.NotNil(t, err)
	assert.Regexp(t, "query user count error", err.Error())
	// UpdatePassword
	err = dao.UpdatePassword(newUser.ID, "new-password")
	assert.NotNil(t, err)
	assert.Regexp(t, "update user password error", err.Error())
	// UpdateUser
	username := "new-username"
	err = dao.UpdateUser(newUser.ID, &user_management_service.UpdateUserBO{Username: &username})
	assert.NotNil(t, err)
	assert.Regexp(t, "update user error", err.Error())
	// DeleteUserByID
	err = dao.DeleteUserByID(0)
	assert.NotNil(t, err)
	assert.Regexp(t, "delete user error", err.Error())
	// getRawUserByUniqueField
	_, err = dao.getRawUserByUniqueField("id", "0")
	assert.NotNil(t, err)
	assert.Regexp(t, "not found user", err.Error())
	// getUserByUniqueField
	_, err = dao.getUserByUniqueField("id", "0")
	assert.NotNil(t, err)
	assert.Regexp(t, "not found user", err.Error())
	// ListUsers
	_, err = dao.ListUsers(1, 10, nil, nil)
	assert.NotNil(t, err)
	assert.Regexp(t, "find users error", err.Error())
	// CheckPassword
	_, err = dao.CheckPassword(newUser.ID, "new-password")
	assert.NotNil(t, err)
}

func TestUserDAOImpl_Count(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})

	dao := NewUserDAOImpl(initializers.DB)
	isActive := true
	isSuperuser := true
	total, err := dao.Count(&isActive, &isSuperuser)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), total)

	_, _ = dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		IsSuperuser: true,
	}, 1, time.Now())
	total, err = dao.Count(&isActive, &isSuperuser)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), total)

	_, _ = dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "2@xx.com",
		Username:    "2",
		IsSuperuser: false,
	}, 2, time.Now())
	total, err = dao.Count(&isActive, nil)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), total)
}

func TestUserDAOImpl_UpdatePassword(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	dao := NewUserDAOImpl(initializers.DB)
	newUser, _ := dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		Password:    "password",
		IsSuperuser: true,
	}, 1, time.Now())

	newPassword := "new-password"
	err := dao.UpdatePassword(newUser.ID, newPassword)
	assert.Nil(t, err)

	isSame, err := dao.CheckPassword(newUser.ID, newPassword)
	assert.Nil(t, err)
	assert.True(t, isSame)
}

func TestUserDAOImpl_UpdateUser(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	dao := NewUserDAOImpl(initializers.DB)
	newUser, _ := dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		Password:    "password",
		IsSuperuser: true,
	}, 1, time.Now())

	err := dao.UpdateUser(newUser.ID, &user_management_service.UpdateUserBO{})
	assert.Nil(t, err)

	nickname := "nickname"
	err = dao.UpdateUser(newUser.ID, &user_management_service.UpdateUserBO{Nickname: &nickname})
	assert.Nil(t, err)
}

func TestUserDAOImpl_ListUsers(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	dao := NewUserDAOImpl(initializers.DB)
	newUser, _ := dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		Password:    "password",
		IsSuperuser: true,
	}, 1, time.Now())

	isActive := true
	isSuperuser := true
	userList, err := dao.ListUsers(1, 10, &isActive, &isSuperuser)
	assert.Nil(t, err)

	assert.Equal(t, newUser.Email, userList[0].Email)
}

func TestUserDAOImpl_GetUser(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	dao := NewUserDAOImpl(initializers.DB)
	newUser, _ := dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		Password:    "password",
		IsSuperuser: true,
	}, 1, time.Now())

	user, err := dao.GetUserByID(newUser.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.Email, newUser.Email)

	user, err = dao.GetUserByEmail(newUser.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, newUser.ID)

	user, err = dao.GetUserByUsername(newUser.Username)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, newUser.ID)

	user, err = dao.GetUserByPhoneNumber(newUser.Username)
	assert.NotNil(t, err)
	assert.Regexp(t, "not found user", err.Error())
}

func TestUserDAOImpl_DeleteUserByID(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	dao := NewUserDAOImpl(initializers.DB)
	newUser, _ := dao.CreateUser(&user_management_service.CreateUserBO{
		Email:       "1@xx.com",
		Username:    "1",
		Password:    "password",
		IsSuperuser: true,
	}, 1, time.Now())

	err := dao.DeleteUserByID(newUser.ID)
	assert.Nil(t, err)
}

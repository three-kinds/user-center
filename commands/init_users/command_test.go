package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/di"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services"
	"github.com/three-kinds/user-center/utils/frame_utils/test_utils"
	"github.com/three-kinds/user-center/utils/generic_utils/testify_addons"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestInitUsers_Success(t *testing.T) {
	test_utils.ClearTables(initializers.DB, &models.User{})
	main()
	userManagementService := di.NewUserManagementService()
	newUserCount := len(userList)

	total, _, err := userManagementService.ListUsers(1, newUserCount, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, total, int64(newUserCount))
	// has existed
	main()
}

func TestMustCreateUser_WithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := services.NewMockIUserManagementService(ctrl)
	serviceMock.EXPECT().CreateUser(gomock.Any()).Return(nil, errors.New("mock error"))
	testify_addons.PanicsWithValueMatch(t, "create user failed", func() {
		mustCreateUser(serviceMock, &userList[0])
	})
}
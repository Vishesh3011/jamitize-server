package service

import (
	"example/internal/models"
	"example/internal/test"
	"example/internal/utils"
	assertions "github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService(t *testing.T) {
	appTest := test.NewAppTest(nil)
	defer appTest.Cancel()
	
	userService := NewUserService(appTest.Context(), appTest.Logger(), appTest.DB(), appTest.Config().JwtConfig().Secret())
	t.Cleanup(func() {
		if err := utils.CleanupCollections(appTest.Context(), appTest.DB()); err != nil {
			t.Fatalf("Error cleaning up collections: %v", err)
		} else {
			t.Log("Collections cleaned up successfully")
		}
	})

	assert := assertions.New(t)

	t.Run("TestCreateUser001", func(t *testing.T) {
		request := &models.CreateUserRequest{
			Name:        "test user",
			Email:       "testuser@jamitize.com",
			Password:    "Abcdefgh@12345678#",
			Instruments: []string{"guitar", "drums"},
			Genres:      []string{"pop", "rock"},
			City:        "Canberra",
			Experience:  "5",
			Bio:         "",
			Socials:     nil,
		}
		user, err := userService.CreateUser(request)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}
		t.Logf("User created: %v", user)
		assert.NotNil(user)
		assert.Equal(request.Name, user.Name)
		assert.Equal(request.Email, user.Email)
		assert.Equal(request.Instruments, user.Instruments)
		assert.Equal(request.Genres, user.Genres)
	})

	t.Run("TestLoginUser005", func(t *testing.T) {
		request := &models.CreateUserRequest{
			Name:        "test user",
			Email:       "testuser2@jamitize.com",
			Password:    "Abcdefgh@12345678#",
			Instruments: []string{"guitar", "drums"},
			Genres:      []string{"pop", "rock"},
			City:        "Canberra",
			Experience:  "5",
			Bio:         "",
			Socials:     nil,
		}
		user, err := userService.CreateUser(request)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}
		t.Logf("User created: %v", user)

		loginUserReq := &models.UserLoginRequest{
			Email:    user.Email,
			Password: request.Password,
		}

		loggedInUser, err := userService.LoginUser(loginUserReq)
		if err != nil {
			t.Fatalf("Error logging in user: %v", err)
		}
		assert.Equal(loginUserReq.Email, loggedInUser.Email)
		assert.Equal(user.Name, loggedInUser.Name)
		assert.Equal(user.Instruments, loggedInUser.Instruments)
		assert.Equal(user.Genres, loggedInUser.Genres)
	})

	t.Run("TestFindUserByID010", func(t *testing.T) {
		request := &models.CreateUserRequest{
			Name:        "test user",
			Email:       "testuser3@jamitize.com",
			Password:    "Abcdefgh@12345678#",
			Instruments: []string{"guitar", "drums"},
			Genres:      []string{"pop", "rock"},
			City:        "Canberra",
			Experience:  "5",
			Bio:         "",
			Socials:     nil,
		}
		user, err := userService.CreateUser(request)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}

		t.Logf("User created: %v", user)
		foundUser, err := userService.FindUserByEmail(user.Email)
		if err != nil {
			t.Fatalf("Error finding user by ID: %v", err)
		}

		assert.NotNil(foundUser)
		assert.Equal(user.ID, foundUser.ID)
		assert.Equal(user.Name, foundUser.Name)
		assert.Equal(user.Email, foundUser.Email)
		assert.Equal(user.Instruments, foundUser.Instruments)
		assert.Equal(user.Genres, foundUser.Genres)
	})
}

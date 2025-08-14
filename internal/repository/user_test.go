package repository

import (
	"example/internal/models"
	"example/internal/test"
	"example/internal/types"
	"example/internal/utils"
	assertions "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var collections = []string{
	"users",
}

func TestUserRepository(t *testing.T) {
	appTest := test.NewAppTest(nil)
	userRepo := NewUserRepository(appTest.Context(), appTest.Logger(), appTest.DB())
	assert := assertions.New(t)

	t.Cleanup(func() {
		if err := utils.CleanupCollections(appTest.Context(), appTest.DB()); err != nil {
			t.Fatalf("Error cleaning up collections: %v", err)
		} else {
			t.Log("Collections cleaned up successfully")
		}
	})

	t.Run("TestCreateUser001", func(t *testing.T) {
		user := &models.User{
			ID:          primitive.NewObjectID(),
			Name:        "test user",
			Email:       "testuser@jamitize.com",
			Password:    "Abcdefgh@12345678#",
			Instruments: nil,
			Genres:      nil,
			City:        "Canberra",
			Experience:  "5",
			Bio:         nil,
			Socials:     nil,
			BandIds:     nil,
		}
		createdUser, err := userRepo.CreateUser(user)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}
		t.Logf("User created: %v", createdUser)
		assert.Equal(user.ID.Hex(), createdUser.ID.Hex())
		assert.Equal(user.Name, createdUser.Name)
		assert.Equal(user.Email, createdUser.Email)
	})

	t.Run("TestCreateUserWithExistingEmail002", func(t *testing.T) {
		user := &models.User{
			ID:         primitive.NewObjectID(),
			Name:       "test user 1",
			Email:      "testuser@jamitize.com",
			Password:   "Abcdefgh@12345678#",
			City:       "Canberra",
			Experience: "5",
		}
		createdUser, err := userRepo.CreateUser(user)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}

		userWithSameEmail := &models.User{
			ID:         primitive.NewObjectID(),
			Name:       "test user 1 duplicate",
			Email:      createdUser.Email,
			Password:   createdUser.Password,
			City:       "Canberra",
			Experience: "5",
		}
		_, err = userRepo.CreateUser(userWithSameEmail)
		assert.Equal(types.Conflict, err.Status())
	})

	t.Run("TestFindUserByEmail005", func(t *testing.T) {
		user := &models.User{
			ID:         primitive.NewObjectID(),
			Name:       "test user 2",
			Email:      "testuser2@jamitize.com",
			Password:   "Abcdefgh@12345678#",
			City:       "Canberra",
			Experience: "5",
		}
		createdUser, err := userRepo.CreateUser(user)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}
		t.Logf("User created: %v", createdUser)

		foundUser, err := userRepo.FindUserByEmail(createdUser.Email)
		if err != nil {
			t.Fatalf("Error finding user by email: %v", err)
		}
		t.Logf("User found by email: %v", foundUser)
		assert.Equal(createdUser.ID.Hex(), foundUser.ID.Hex())
		assert.Equal(createdUser.Name, foundUser.Name)
		assert.Equal(createdUser.Email, foundUser.Email)
		assert.Equal(createdUser.City, foundUser.City)
	})

	t.Run("TestFindUserByID010", func(t *testing.T) {
		user := &models.User{
			ID:         primitive.NewObjectID(),
			Name:       "test user 3",
			Email:      "testuser3@jamitize.com",
			Password:   "Abcdefgh@12345678#",
			City:       "Canberra",
			Experience: "5",
		}
		createdUser, err := userRepo.CreateUser(user)
		if err != nil {
			t.Fatalf("Error creating user: %v", err)
		}

		t.Logf("User created: %v", createdUser)
		foundUser, err := userRepo.FindUserByID(createdUser.ID)
		if err != nil {
			t.Fatalf("Error finding user by ID: %v", err)
		}
		t.Logf("User found by ID: %v", foundUser)
		assert.Equal(createdUser.ID.Hex(), foundUser.ID.Hex())
		assert.Equal(createdUser.Name, foundUser.Name)
		assert.Equal(createdUser.Email, foundUser.Email)
		assert.Equal(createdUser.City, foundUser.City)
		assert.Equal(createdUser.Experience, foundUser.Experience)
	})
}

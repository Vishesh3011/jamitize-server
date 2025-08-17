package models

import (
	"example/internal/types"
	"example/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id"`
	Name         string               `bson:"name"`
	Email        string               `bson:"email"`
	Password     string               `bson:"password"`
	Instruments  []string             `bson:"instruments"`
	Genres       []string             `bson:"genres"`
	City         string               `bson:"city"`
	Experience   string               `bson:"experience"`
	Bio          *string              `bson:"bio,omitempty"`
	Socials      []string             `bson:"socials"`
	BandIds      []primitive.ObjectID `bson:"band_ids,omitempty"`
	RefreshToken string               `bson:"refresh_token,omitempty"`
	State        types.State          `bson:"state,omitempty"`
	CreatedAt    time.Time            `bson:"created_at,omitempty"`
	UpdatedAt    time.Time            `bson:"updated_at,omitempty"`
}

func (u User) ToUserResponse(token *string) *UserResponse {
	return &UserResponse{
		Name:         u.Name,
		Email:        u.Email,
		Instruments:  u.Instruments,
		Genres:       u.Genres,
		City:         u.City,
		Experience:   u.Experience,
		Bio:          u.Bio,
		Socials:      u.Socials,
		Token:        token,
		RefreshToken: u.RefreshToken,
		State:        u.State,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

type UserResponse struct {
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Instruments  []string    `json:"instruments"`
	Genres       []string    `json:"genres"`
	City         string      `json:"city"`
	Experience   string      `json:"experience"`
	Bio          *string     `json:"bio,omitempty"`
	Socials      []string    `json:"socials"`
	Token        *string     `json:"token,omitempty"`
	RefreshToken string      `json:"refreshToken"`
	State        types.State `json:"state"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

type CreateUserRequest struct {
	Name        string   `json:"name" validate:"required"`
	Email       string   `json:"email" validate:"required,email"`
	Password    string   `json:"password" validate:"required"`
	Instruments []string `json:"instruments"`
	Genres      []string `json:"genres"`
	City        string   `json:"city"`
	Experience  string   `json:"experience"`
	Bio         string   `json:"bio,omitempty"`
	Socials     []string `json:"socials"`
}

func (r CreateUserRequest) ToUser() *User {
	return &User{
		ID:          primitive.NewObjectID(),
		Name:        r.Name,
		Email:       r.Email,
		Password:    utils.HashPassword(r.Password),
		Instruments: r.Instruments,
		Genres:      r.Genres,
		City:        r.City,
		Experience:  r.Experience,
		Bio:         &r.Bio,
		Socials:     r.Socials,
		State:       types.ACTIVE,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

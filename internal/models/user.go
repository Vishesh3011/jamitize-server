package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Instruments []string             `bson:"instruments"`
	Genres      []string             `bson:"genres"`
	City        string               `bson:"city"`
	Experience  string               `bson:"experience"`
	Bio         *string              `bson:"bio,omitempty"`
	Socials     []string             `bson:"socials"`
	BandIds     []primitive.ObjectID `bson:"band_ids,omitempty"`
}

func (u User) ToUserResponse() *UserResponse {
	return &UserResponse{
		Name:        u.Name,
		Instruments: u.Instruments,
		Genres:      u.Genres,
		City:        u.City,
		Experience:  u.Experience,
		Bio:         u.Bio,
		Socials:     u.Socials,
	}
}

type UserResponse struct {
	Name        string   `json:"name"`
	Instruments []string `json:"instruments"`
	Genres      []string `json:"genres"`
	City        string   `json:"city"`
	Experience  string   `json:"experience"`
	Bio         *string  `json:"bio,omitempty"`
	Socials     []string `json:"socials"`
}

type CreateUserRequest struct {
	Name        string   `json:"name"`
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
		Instruments: r.Instruments,
		Genres:      r.Genres,
		City:        r.City,
		Experience:  r.Experience,
		Bio:         &r.Bio,
		Socials:     r.Socials,
	}
}

package entities

import "time"

type UserEntity struct {
	Id        string    `json:"id" bson:"_id" validate:"string"`
	Name      string    `json:"name" bson:"name" validate:"required"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password" bson:"password" validate:"required"`
	Active    bool      `json:"active" bson:"active" validate:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" validate:"datetime"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at" validate:"datetime"`
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required,min=1,max=50"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required,min=1,max=50"`
	Age       *int               `json:"age,omitempty" bson:"age,omitempty" validate:"omitempty,min=0,max=150"`
	IsActive  bool               `json:"isActive" bson:"isActive"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,min=1,max=50"`
	LastName  string `json:"lastName" validate:"required,min=1,max=50"`
	Age       *int   `json:"age,omitempty" validate:"omitempty,min=0,max=150"`
}

type UpdateUserRequest struct {
	Email     *string `json:"email,omitempty" validate:"omitempty,email"`
	FirstName *string `json:"firstName,omitempty" validate:"omitempty,min=1,max=50"`
	LastName  *string `json:"lastName,omitempty" validate:"omitempty,min=1,max=50"`
	Age       *int    `json:"age,omitempty" validate:"omitempty,min=0,max=150"`
	IsActive  *bool   `json:"isActive,omitempty"`
}

type UsersResponse struct {
	Users  []User `json:"users"`
	Total  int64  `json:"total"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Database  string    `json:"database"`
	Version   string    `json:"version,omitempty"`
}

type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

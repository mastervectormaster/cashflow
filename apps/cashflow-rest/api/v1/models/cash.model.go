package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Cash struct {
	// Id       primitive.ObjectID `json:"_id,omitempty"`
	Name     string  `json:"name,omitempty" validate:"required"`
	Amount   float64 `json:"amount,omitempty" validate:"required"`
	Platform string  `json:"platform,omitempty" validate:"required"`
	Url      string  `json:"url,omitempty"`
	// Img      primitive.Binary   `json:"img,omitempty"`
	Notes string `json:"notes,omitempty"`
}

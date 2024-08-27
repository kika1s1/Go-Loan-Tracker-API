package domain

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Loan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Amount      float64            `bson:"amount"`
	Status      string             `bson:"status"` // pending, approved, rejected
	CreatedAt   string             `bson:"created_at"`
	UpdatedAt   string             `bson:"updated_at"`
}


type LoanStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

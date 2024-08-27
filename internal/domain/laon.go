package domain

import "time"

type Loan struct {
    ID         string    `bson:"_id,omitempty" json:"_id,omitempty"`
    UserID     string    `bson:"user_id" json:"user_id"`
    Amount     float64   `bson:"amount" json:"amount"`
    Status     string    `bson:"status" json:"status"`
    CreatedAt  time.Time `bson:"created_at" json:"created_at"`
    UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}

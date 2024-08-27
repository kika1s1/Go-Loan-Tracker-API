package domain
import "time"

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	IsAdmin   bool      `bson:"is_admin" json:"is_admin"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

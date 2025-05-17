package types

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

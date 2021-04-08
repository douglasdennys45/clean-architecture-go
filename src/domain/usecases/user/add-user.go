package user

type AddUserParam struct {
	Name     string `json:"name" bson:"name" validate:"required,string"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type AddUser interface {
	Add(addUser *AddUserParam) error
}

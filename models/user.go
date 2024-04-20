package models

type User struct {
	ID        int64
	FirstName string
	LastName  string `binding:"required"`
	Email     string `binding:"required"`
}

func NewUser(id int64, fname string, lname string, email string) *User {
	return &User{
		ID:        id,
		FirstName: fname,
		LastName:  lname,
		Email:     email,
	}

}

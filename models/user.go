package models

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	//Wonder if i should make this a private property and not export it?
}

func NewUser(id int64, fname string, lname string, email string, password string) *User {
	return &User{
		ID:        id,
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  password,
	}

}

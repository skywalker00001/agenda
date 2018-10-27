package entity

// User an entity with Name, Password, Email and Phone
type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

// NewUser constuct an user according to the params
// params: a string indicate the name of the user
// params: a string indicate the password of the user
// params: a string indicate the email of the user
// params: a string indicate the phone of the user
func NewUser(name, password, email, phone string) *User {
	return &User{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
}

// GetName get the name of the user
// return: a string indicate the name of the user
func (u User) GetName() string {
	return u.Name
}

// SetName set a new name for the user
// param: a string indicate the new name of the user
func (u *User) SetName(newName string) {
	u.Name = newName
}

// GetPassword get the password of the user
// return: a string indicate the password of the user
func (u User) GetPassword() string {
	return u.Password
}

// SetPassword set a new password for the user
// param: a string indicate the new password of the user
func (u *User) SetPassword(newPassword string) {
	u.Password = newPassword
}

// GetEmail get the email of the user
// return: a string indicate the email of the user
func (u User) GetEmail() string {
	return u.Email
}

// SetEmail set a new email for the user
// param: a string indicate the new email of the user
func (u *User) SetEmail(newEmail string) {
	u.Email = newEmail
}

// GetPhone get the phone of the user
// return: a string indicate the phone of the user
func (u User) GetPhone() string {
	return u.Phone
}

// SetPhone set a new phone for the user
// param: a string indicate the new phone of the user
func (u *User) SetPhone(newPhone string) {
	u.Phone = newPhone
}

func (u *User) assign(user User) {
	u.Name = user.Name
	u.Password = user.Password
	u.Email = user.Email
	u.Phone = user.Phone
}

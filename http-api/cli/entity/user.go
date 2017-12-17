package entity

type User struct {
	username string
	password string
	email    string
	phone    string
}

func (u *User) SetUsername(s string) {
	u.username = s
}
func (u *User) SetPassword(s string) {
	u.password = s
}
func (u *User) SetEmail(s string) {
	u.email = s
}
func (u *User) SetPhone(s string) {
	u.phone = s
}
func (u *User) SetUser(username, password, email, phone string) {
	u.username = username
	u.password = password
	u.email = email
	u.phone = phone
}
func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetEmail() string {
	return u.email
}
func (u *User) GetPhone() string {
	return u.phone
}
func (u *User) GetUser() (string, string, string, string) {
	return u.username, u.password, u.email, u.phone
}

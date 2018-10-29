package entity

// User with
type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

// InitUser init the user information with the argument
func (m *User) InitUser(tUserName, tUserPassword, tUserEmail, tUserPhone string) {
	m.Name = tUserName
	m.Password = tUserPassword
	m.Email = tUserEmail
	m.Phone = tUserPhone
}

func (m *User) copyUser(t User) {
	m.Name = t.Name
	m.Password = t.Password
	m.Email = t.Email
	m.Phone = t.Phone
}

func (m User) getName() string {
	return m.Name
}

func (m *User) setName(n string) {
	m.Name = n
}

func (m *User) getPassword() string {
	return m.Password
}

func (m *User) setPassword(p string) {
	m.Password = p
}
func (m User) getEmail() string {
	return m.Email
}

func (m *User) setEmail(e string) {
	m.Email = e
}
func (m User) getPhone() string {
	return m.Phone
}

func (m *User) setPhone(p string) {
	m.Phone = p
}

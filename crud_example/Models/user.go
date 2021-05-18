package Models

type User struct {
	firstname string
	lastname  string
	email     string
	age       int
}

func (self *User) SetFristname(name string) {
	self.firstname = name

}

func (self *User) SetLastname(lastname string) {
	self.lastname = lastname

}

func (self *User) SetEmail(email string) {
	self.email = email

}

func (self *User) SetAge(age int) {
	self.age = age

}

func (self *User) GetFristname() string {
	return self.firstname

}

func (self *User) GetLastname() string {
	return self.lastname

}

func (self *User) GetEmail() string {
	return self.email

}

func (self *User) GetAge() int {
	return self.age

}

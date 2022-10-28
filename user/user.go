package user

type UserBuilder struct {
	name     string
	email    string
	age      int
	salary   float64
	isActive bool
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}
func (userBuilder *UserBuilder) WithName(name string) *UserBuilder {
	userBuilder.name = name
	return userBuilder
}
func (userBuilder *UserBuilder) WithEmail(email string) *UserBuilder {
	userBuilder.email = email
	return userBuilder
}
func (userBuilder *UserBuilder) WithAge(age int) *UserBuilder {
	userBuilder.age = age
	return userBuilder
}
func (userBuilder *UserBuilder) WithSalary(salary float64) *UserBuilder {
	userBuilder.salary = salary
	return userBuilder
}
func (userBuilder *UserBuilder) WithIsActive(active bool) *UserBuilder {
	userBuilder.isActive = active
	return userBuilder
}
func (userBuilder *UserBuilder) Build() *User {
	user := &User{}
	user.name = userBuilder.name
	user.email = userBuilder.email
	user.age = userBuilder.age
	user.salary = userBuilder.salary
	user.isActive = userBuilder.isActive
	return user
}

type User struct {
	name     string
	email    string
	age      int
	salary   float64
	isActive bool
}

func (user *User) GetName() string {
	return user.name
}
func (user *User) GetEmail() string {
	return user.email
}
func (user *User) GetAge() int {
	return user.age
}
func (user *User) GetSalary() float64 {
	return user.salary
}
func (user *User) GetIsActive() bool {
	return user.isActive
}

package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

type UserDate struct {
	users map[int]User
}

type MapDate interface {
	SaveDate(User) (User, error)
	GetDate(int) (User, error)
}

func (u *UserDate) SaveDate(su User) error {
	user, ok := u.users[su.id]
	if ok {
		return fmt.Errorf("%#v is already exit ", user)
	}
	u.users[su.id] = su
	return nil
}

func (u *UserDate) GetDate(id int) (User, error) {
	user, ok := u.users[id]
	if ok != true {
		return User{}, fmt.Errorf("%#v does't find ", id)
	}
	return user, nil
}

// use service cover all the XxxDate
type Service struct {
	UserDate
}

func main() {
	us := &Service{
		UserDate{
			make(map[int]User),
		},
	}
	u1 := User{
		id:   1,
		name: "kaka",
	}
	us.SaveDate(u1)
	uget, err := us.GetDate(1)

	fmt.Println("-------------------")
	fmt.Println(uget, err)
	fmt.Println(us.UserDate.users)
}

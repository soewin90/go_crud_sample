package models // supportive package name , package name match with folder name
import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// variable block
var (
	users  []*User // slice type
	nextID = 1     // package level no need := implicit initialization or nextID  int32= 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		// User{} mean empty user
		return User{}, errors.New("New user ID must be greater than 0!")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u) // take address because of Pointer type *User
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			// return user pointer and nil(no error)
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			// splice opearatio over slice,
			// users[:i], get user update to i, but not included i
			// then users[i+1:]... , get user from i+1 to end of users
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}

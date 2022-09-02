package main

type user struct {
	id int
	username int
	password int
}

type userService struct {
	userList []user
}

func (usr userService) getAllUsers() []user{
	return usr.userList
}

func (usr userService) getUserById(id int) user{
	for _, currentUser := range usr.userList {
		if id == currentUser.id {
			return currentUser
		}
	}
	return user{}
}
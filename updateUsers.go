package main


func (u User) sendMessage(message string, messageCharLimit int) (string, bool) {
	if len(message) <= u.messageCharLimit{
		return message, true
	}
	return "",false


}

type User struct {
	Name string
	membership
}

func newUser(name string, membershipType string) User {
	user := User{}
	user.Name = name
	user.Type = membershipType
	if membershipType== "Premium"{
		user.messageCharLimit=1000
	}else{
		user.messageCharLimit=100
	}

	return user
}

type membership struct {
	Type string
	messageCharLimit int
}
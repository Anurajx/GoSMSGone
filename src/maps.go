package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers){
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{name: names[i], number: phoneNumbers[i]}
	}
	return userMap, nil
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	_, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !users[name].scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

// type user struct {
// 	name        string
// 	number int
// }

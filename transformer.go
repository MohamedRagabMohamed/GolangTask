package main

func UserTransformer(users map[int]User) []User {
	result := make([]User, 0)
	for _, value := range users {
		result = append(result, value)
	}
	return result
}

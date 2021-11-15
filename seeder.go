package main

import "github.com/icrowley/fake"

func Seeder() {
	UserSeeder(10)
}

func UserSeeder(count int) {
	for i := 1; i <= count; i++ {
		Users[i] = User{
			Name:        fake.FullName(),
			Email:       fake.EmailAddress(),
			Password:    fake.SimplePassword(),
			PhoneNumber: fake.Phone(),
		}
	}
}

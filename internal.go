package main

func createUser(email, password string) (*User, error) {
	user := &User{Email: email, Password: password}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func getUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

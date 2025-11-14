package solidprinciples

import "gorm.io/gorm"

//Single Responsibility Principle (SRP)
//This principle states that every class or struct instance should have one type of responsibility
//Bad Example as here the user struct's methods are tied to return the fullname and also saving the data to db
type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Save() error {
	//Saving data to user logic here
	var err error
	return err
}

//This is the right example as here the UserRight struct has only one responsibility...to get the full name
//And USerRepo has one that is to save the user to the DB like we have done in the Gin_Project
type UserRight struct {
	FirstName string
	LastName  string
}

func (u *UserRight) GetFullNameRight() string {
	return u.FirstName + " " + u.LastName
}

type UserRepository struct {
	DB gorm.DB
	// Database connection or other storage-related fields
}

func (r *UserRepository) Save(u *User) error {
	// Save user to the database logic
	var err error
	return err
}

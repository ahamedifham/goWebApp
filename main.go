package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string 
	Age int 
	Birthdate time.Time
}

func (u *User) printFirstName() string {
	return u.FirstName
}

func main(){
	user := User{
		FirstName: "Ahamed",
		LastName: "Ifham",
		PhoneNumber: "0777290914",
	}

	var user2 User

	user2.FirstName = "AhamedI"

	log.Println(user.Birthdate)
	log.Println(user2.printFirstName())
}

func saySomething(s string) (string, string){
	log.Println("s from the saySomething is ", s)
	return s,"Word"
}

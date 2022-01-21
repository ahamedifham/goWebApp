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

	myMap := make(map[string]string)
	myMap["dog"] = user.printFirstName()
	log.Println(myMap["dog"])

	var mySlice []string
	log.Println("mySlice before appending", &mySlice)

	mySlice = append(mySlice, "ai")

	log.Println("mySlice after first appending", &mySlice)
	mySlice2 := append(mySlice, "hj")

	log.Println(mySlice[len(mySlice)-1])
	log.Println(mySlice2)


	numbers := []int{1,2,4,5,12,58,79}
	log.Println(numbers)
	log.Println(numbers[3:len(numbers)])

	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue")
	} else {
		log.Println("Is not true")
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("cat is ", cat)
	}
}

func saySomething(s string) (string, string){
	log.Println("s from the saySomething is ", s)
	return s,"Word"
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	helpers "github.com/ahamedi/myniceprogram/helper"
)

var s = "seven"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

type User struct {
	FirstName string
	LastName string
	PhoneNumber string 
	Age int 
	Birthdate time.Time
}

type Person struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
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

	for i :=0; i < 10; i++{
		log.Print(i)
	}

	animals := []string{"dog", "fish", "horse", "cow"}

	for _, animal := range animals {
		log.Println(animal)
	}

	houses := make(map[string]string)
	houses["one"] = "home"
	houses["two"] = "kalubowila"

	for houseKey,house := range houses {
		log.Println(houseKey, house)
	}

	var firstLine string = "Once upon a midnight"

	for _, letter := range firstLine{
		log.Println(letter)
	}

	dog := Dog{"Samson", "GermanShepeherd"}

	PrintInfo(&dog)

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)

	myJson := `
	[
		{
			"first_name":"Ahamed",
			"last_name":"Ifham"
		},
		{
			"first_name":"Ian",
			"last_name":"Bell"
		}

	]`

	var unmarshled []Person

	err := json.Unmarshal([]byte(myJson), unmarshled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshelled: %v", unmarshled)

	var persons []Person

	m1 := Person{"Ahamed", "Ifham"}
	m2 := Person{"John", "Cat"}


	persons = append(persons, m1)
	persons = append(persons, m2)

	newJson, err := json.MarshalIndent(persons, "", " ")

	if err != nil {
		log.Println("Error marshalling", err)
	}

	fmt.Print(string(newJson))

	result, error := divide(100.00, 10.00)

	if error != nil {
		log.Println(error)
		return
	}

	log.Println(result)


}

func CalculateValue(intchan chan int){
	randomNumber := helpers.RandomNumber(10)
	intchan <- randomNumber
}

func saySomething(s string) (string, string){
	log.Println("s from the saySomething is ", s)
	return s,"Word"
}

func (d *Dog) Says() string{
	return "Woof My Name is "  + d.Name
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal){
	log.Println("This animal says", a.Says() , "and has", a.NumberOfLegs(), "legs")
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y==0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x/y
	return result, nil
}
package main

import "fmt"

func main() {
  fmt.Println("This is an struct video")

  myUser := User{"Rajat", "bisrajat123@gmail.com", true, 23, 40}
  fmt.Println(myUser)
  fmt.Printf("%+v\n", myUser)
  myUser.GetStatus()
  myUser.NewMail()
}

type User struct {
  Name   string
  Email  string
  Status bool
  Age    int
  oneAge int // Cannot be exported
}


func (u User) GetStatus(){
  fmt.Println("IS user active: ", u.Status)
}

func (u User) NewMail(){
  u.Email = "test@go.dev"
  fmt.Println("Eamil of this user is: ", u.Email)
}
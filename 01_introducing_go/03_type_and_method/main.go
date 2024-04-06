package main 

import (
  "fmt"
)

type Person struct {
  FirstName string
  LastName string
  Age int
}

func (p Person) Name() string {
  return p.FirstName + " " + p.LastName
}

func (p Person) Describe() string {
  return fmt.Sprintf("%s is %d years old.\n", p.Name(), p.Age)
}

func main() {
  rajesh := Person{
    FirstName: "Rajesh", 
    LastName: "Pillai",
    Age: 999999,
  }
  fmt.Println(rajesh.Describe())
}







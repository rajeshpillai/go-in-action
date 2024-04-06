package main 

import (
  "fmt"
)

type Describer interface {
  Describe() string 
}

func PrintDescription(d Describer) {
  fmt.Println(d.Describe()) 
}

type Person struct {
  FirstName string
  LastName string 
  Age int 
}

func (p Person) Name() string {
  return p.FirstName + " " + p.LastName 
}

func (p Person) Describe() string {
  return fmt.Sprintf("%s is %d year old human", p.Name(), p.Age)
}

type Cat struct {
  Name string 
  Attitude string
  Color string 
}

func (c Cat) Describe() string {
  return fmt.Sprintf("%s is %s cat (%s) ", c.Name, c.Color, c.Attitude)
}

func main () {
  rajesh := Person{
    FirstName: "Rajesh",
    LastName: "Pillai",
    Age: 99999999,
  }

  tom := Cat{
    Name: "Tom",
    Attitude: "the best boss",
    Color: "black and white",
  }

  gini := Cat{
    Name: "Gini",
    Attitude: "kind of sneaky",
    Color: "orange",
  }

  PrintDescription(rajesh) 
  PrintDescription(tom) 
  PrintDescription(gini)

}

















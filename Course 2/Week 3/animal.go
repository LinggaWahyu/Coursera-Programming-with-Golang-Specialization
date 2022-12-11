package main

import (
	"fmt"
	"os"
)

type Animal struct {
	foodEaten, locomotionMethod, spokenSound string
}

func (a Animal) Eat() {
	fmt.Println(a.foodEaten)
}

func (a Animal) Locomotion() {
	fmt.Println(a.locomotionMethod)
}

func (a Animal) Speak() {
	fmt.Println(a.spokenSound)
}

func main() {
	data := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		fmt.Print(" => ")
		animalSubject := "0"
		actionOfAnimal := "0"
		_, err := fmt.Scan(&animalSubject, &actionOfAnimal)
		if err != nil {
			fmt.Println("Invalid input = ", err)
			os.Exit(0)
		}
		if actionOfAnimal == "eat" {
			data[animalSubject].Eat()
		} else if actionOfAnimal == "move" {
			data[animalSubject].Locomotion()
		} else if actionOfAnimal == "speak" {
			data[animalSubject].Speak()
		}
	}
}

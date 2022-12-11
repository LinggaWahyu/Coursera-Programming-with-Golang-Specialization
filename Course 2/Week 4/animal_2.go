package main

import (
	"fmt"
)

type Cow struct {
	foodEaten, locomotionMethod, spokenSound, name string
}

type Snake struct {
	foodEaten, locomotionMethod, spokenSound, name string
}

type Bird struct {
	foodEaten, locomotionMethod, spokenSound, name string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
	getName() string
}

func (c Cow) Eat() {
	fmt.Println(c.foodEaten)
	return
}

func (c Cow) Move() {
	fmt.Println(c.locomotionMethod)
	return
}

func (c Cow) Speak() {
	fmt.Println(c.spokenSound)
	return
}

func (s Snake) Eat() {
	fmt.Println(s.foodEaten)
	return
}

func (s Snake) Move() {
	fmt.Println(s.locomotionMethod)
	return
}

func (s Snake) Speak() {
	fmt.Println(s.spokenSound)
	return
}

func (b Bird) Eat() {
	fmt.Println(b.foodEaten)
	return
}

func (b Bird) Move() {
	fmt.Println(b.locomotionMethod)
	return
}

func (b Bird) Speak() {
	fmt.Println(b.spokenSound)
	return
}

func (c Cow) getName() string {
	return c.name
}

func (s Snake) getName() string {
	return s.name
}

func (b Bird) getName() string {
	return b.name
}

func main() {
	cow := Cow{
		foodEaten:        "grass",
		locomotionMethod: "walk",
		spokenSound:      "moo",
		name:             "cow",
	}

	bird := Bird{
		foodEaten:        "worms",
		locomotionMethod: "fly",
		spokenSound:      "peep",
		name:             "bird",
	}

	snake := Snake{
		foodEaten:        "mice",
		locomotionMethod: "slither",
		spokenSound:      "hsss",
		name:             "snake",
	}

	sliceOfAnimals := []animalInterface{}
	sliceOfAnimals = append(sliceOfAnimals, cow)
	sliceOfAnimals = append(sliceOfAnimals, bird)
	sliceOfAnimals = append(sliceOfAnimals, snake)

	for {
		var command, animalRequest, typeRequest string
		fmt.Print(" > ")
		_, err := fmt.Scan(&command, &animalRequest, &typeRequest)
		if err != nil {
			fmt.Println("Error Occurred while reading input= ", err)
		}
		if command == "query" {
			for _, animal := range sliceOfAnimals {
				if animal.getName() == animalRequest {
					if typeRequest == "move" {
						animal.Move()
					} else if typeRequest == "eat" {
						animal.Eat()
					} else if typeRequest == "speak" {
						animal.Speak()
					}
				}
			}
		}
		if command == "newanimal" {
			if typeRequest == "cow" {
				sliceOfAnimals = append(sliceOfAnimals, Cow{
					foodEaten:        cow.foodEaten,
					locomotionMethod: cow.locomotionMethod,
					spokenSound:      cow.spokenSound,
					name:             animalRequest,
				})
				fmt.Println("Created it!")
			} else if typeRequest == "snake" {
				sliceOfAnimals = append(sliceOfAnimals, Snake{
					foodEaten:        snake.foodEaten,
					locomotionMethod: snake.locomotionMethod,
					spokenSound:      snake.spokenSound,
					name:             animalRequest,
				})
				fmt.Println("Created it!")
			} else if typeRequest == "bird" {
				sliceOfAnimals = append(sliceOfAnimals, Bird{
					foodEaten:        bird.foodEaten,
					locomotionMethod: bird.locomotionMethod,
					spokenSound:      bird.spokenSound,
					name:             animalRequest,
				})
				fmt.Println("Created it!")
			} else {
				fmt.Println("Invalid Animal")
			}
		}
	}
}

package main

import "fmt"

func main() {
	var a, v0, s0 float64
	var time float64
	fmt.Printf("Enter value of acceleration(a) \t\t\t = ")
	fmt.Scanln(&a)

	fmt.Printf("Enter value of Initial Velocity(v0) \t\t = ")
	fmt.Scanln(&v0)

	fmt.Printf("Enter value of Initial Displacement(s0) \t = ")
	fmt.Scanln(&s0)

	fmt.Printf("Enter value of Time(time) \t\t\t\t = ")
	fmt.Scanln(&time)
	fmt.Println(" ===================================")

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %v Seconds \t\t\t = %v\n", time, fn(time))
	fmt.Println(" ===================================")
}

func GenDisplaceFn(acc, iniVelocity, iniDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (1/2)*acc*(time*time) + iniVelocity*time + iniDisplacement
	}
}

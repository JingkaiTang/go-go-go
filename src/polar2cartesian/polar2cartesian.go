package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const result = "Polar radius=%.02f 0=%.02f° → Cartesian x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, or q to quit."

type polar struct {
	radius float64
	θ float64
}

type cartesian struct {
	x float64
	y float64
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ) 
			y := polarCoord.radius * math.Sin(θ)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius adn angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, θ float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
			if strings.HasPrefix(line, "q") {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, θ}
		coord := <-answers
		fmt.Printf(result, radius, θ, coord.x, coord.y)
	}
	fmt.Println()
}

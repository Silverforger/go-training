package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Handler function handles the HTTP response writer and checks if provided URL path is correct, returns an error if not.
func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/solve" {
			message := "Invalid URL. " + r.URL.Path + " does not exist."
			http.Error(w, message, http.StatusBadRequest)
		} else {
			calculate(w, r)
		}
	}
}

// Calculate function handles the input parameters in the URL and runs the three unknowns solver if inputs are valid.
// This will return an error if the parameter is incorrect, and if the coefficient values are not correct.
// This will also default string values (unparseable string to float64 values) to 0.
func calculate(w http.ResponseWriter, r *http.Request) {
	queryVal := r.URL.Query()
	for paramsInput := range queryVal {
		if paramsInput != "coef" {
			fmt.Println("Invalid parameter, " + paramsInput + " does not exist.")
		} else if paramsInput == "coef" {
			paramsInputCoef := r.URL.Query()["coef"]
			vals := strings.Split(paramsInputCoef[0], ",")
			coefficients := make([]float64, 0)
			if len(vals) != 12 {
				fmt.Println("Invalid input. 12 coefficient values are needed.")
			} else {
				for _, stringInput := range vals {
					floatVar, err := strconv.ParseFloat(stringInput, 32)
					if err != nil {
						fmt.Println("Parse error. One of your inputs was not a number. Defaulting that value to 0.")
						coefficients = append(coefficients, 0)
						continue
					} else if err == nil {
						coefficients = append(coefficients, floatVar)
						continue
					}
				}
				fmt.Fprintf(w, threeUnknownsSolver(coefficients, w))
			}
		}
	}
}

// determinantSolver function returns the determinant of a given matrix.
func determinantSolver(a float64, b float64, c float64, d float64) float64 {
	return a*d - b*c
}

// threeUnknownsSolver function uses the Cramer's Rule to solve for the unknown variables given three equations.
// This calls the determinantSolver function for each calculation of determinant matrix.
// This will return one of three outputs, depending on the values of D, Dx, Dy, and Dz.
func threeUnknownsSolver(coefArray []float64, w http.ResponseWriter) string {
	var conclusion string

	x1 := coefArray[0]
	y1 := coefArray[1]
	z1 := coefArray[2]
	c1 := coefArray[3]
	x2 := coefArray[4]
	y2 := coefArray[5]
	z2 := coefArray[6]
	c2 := coefArray[7]
	x3 := coefArray[8]
	y3 := coefArray[9]
	z3 := coefArray[10]
	c3 := coefArray[11]

	D := (x1 * determinantSolver(y2, z2, y3, z3)) - (y1 * determinantSolver(x2, z2, x3, z3)) + (z1 * determinantSolver(x2, y2, x3, y3))
	Dx := (c1 * determinantSolver(y2, z2, y3, z3)) - (y1 * determinantSolver(c2, z2, c3, z3)) + (z1 * determinantSolver(c2, y2, c3, y3))
	Dy := (x1 * determinantSolver(c2, z2, c3, z3)) - (c1 * determinantSolver(x2, z2, x3, z3)) + (z1 * determinantSolver(x2, c2, x3, y3))
	Dz := (x1 * determinantSolver(y2, c2, y3, c3)) - (y1 * determinantSolver(x2, c2, x3, c3)) + (c1 * determinantSolver(x2, y2, x3, y3))

	x := Dx / D
	y := Dy / D
	z := Dz / D

	fmt.Fprintf(w, "System: \n%vx + %vy + %vz = %v\n%vx + %vy + %vz = %v\n%vx + %vy + %vz = %v\n", x1, y1, z1, c1, x2, y2, z2, c2, x3, y3, z3, c3)

	if D == 0 {
		if Dx == 0 && Dy == 0 && Dz == 0 {
			conclusion = "dependent - with multiple solutions"
		} else {
			conclusion = "inconsistent - no solution"
		}
	} else {
		conclusion = fmt.Sprintf("Solution:\nx = %v, y = %v, z = %v", x, y, z)
	}

	return conclusion
}

func main() {
	http.HandleFunc("/", handler())
	http.ListenAndServe(":8080", nil)
}

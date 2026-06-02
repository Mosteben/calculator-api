package services

import (
	"errors"
	"math"

	"github.com/Mosteben/calculator-api/db"
)

func Calculate(a int, b int, op string) (float64, error) {

	var result float64

	switch op {

	case "add":
		result = float64(a + b)

	case "subtract":
		result = float64(a - b)

	case "multiply":
		result = float64(a * b)

	case "divide":
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		result = float64(a) / float64(b)

	case "modulo":
		result = float64(a % b)

	case "power":
		result = math.Pow(float64(a), float64(b))

	case "sqrt":
		result = math.Sqrt(float64(a))

	case "average":
		result = float64(a+b) / 2

	default:
		return 0, errors.New("invalid operation")
	}

	
	_, err := db.DB.Exec(
		"INSERT INTO history (a, b, operation, result) VALUES ($1, $2, $3, $4)",
		a, b, op, result,
	)

	if err != nil {
		return 0, err
	}

	return result, nil
}
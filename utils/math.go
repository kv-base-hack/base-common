package utils

import (
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(val*ratio) / ratio
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func RoundingString(number string, precision int) string {
	pos := strings.Index(number, ".")
	if pos == -1 {
		return number
	}
	return number[:pos+1+Min(len(number)-1-pos, precision)]
}

// Multiply return a * b
func MultiplyString(a, b string) (string, error) {
	af, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	bf, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	return af.Mul(bf).String(), nil
}

// DivideString return a / b
func DivideString(a, b string) (string, error) {
	af, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	bf, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	return af.Div(bf).String(), nil
}

// AddString return a + b
func AddString(a, b string) (string, error) {
	af, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	bf, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	return af.Add(bf).String(), nil
}

// SubString return a - b
func SubString(a, b string) (string, error) {
	af, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	bf, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	return af.Sub(bf).String(), nil
}

// AbsString return abs(a)
func AbsString(a string) (string, error) {
	af, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	return af.Abs().String(), nil
}

// SubString return abs(a - b)
func SubAbsString(a, b string) (string, error) {
	af, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	bf, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	return af.Sub(bf).Abs().String(), nil
}

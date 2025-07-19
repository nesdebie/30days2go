package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UnitCategory string

const (
	Length     UnitCategory = "length"
	Mass       UnitCategory = "mass"
	Temperature UnitCategory = "temperature"
)

type Unit struct {
	Category UnitCategory
	ToBase   func(float64) float64
	FromBase func(float64) float64
}

var units = map[string]Unit{
	"mm": {Length, func(v float64) float64 { return v * 0.001 }, func(v float64) float64 { return v / 0.001 }},
	"cm": {Length, func(v float64) float64 { return v * 0.01 }, func(v float64) float64 { return v / 0.01 }},
	"m":  {Length, func(v float64) float64 { return v }, func(v float64) float64 { return v }},
	"km": {Length, func(v float64) float64 { return v * 1000 }, func(v float64) float64 { return v / 1000 }},
	"in": {Length, func(v float64) float64 { return v * 0.0254 }, func(v float64) float64 { return v / 0.0254 }},
	"ft": {Length, func(v float64) float64 { return v * 0.3048 }, func(v float64) float64 { return v / 0.3048 }},
	"yd": {Length, func(v float64) float64 { return v * 0.9144 }, func(v float64) float64 { return v / 0.9144 }},
	"mi": {Length, func(v float64) float64 { return v * 1609.344 }, func(v float64) float64 { return v / 1609.344 }},

	"g":  {Mass, func(v float64) float64 { return v / 1000 }, func(v float64) float64 { return v * 1000 }},
	"kg": {Mass, func(v float64) float64 { return v }, func(v float64) float64 { return v }},
	"lb": {Mass, func(v float64) float64 { return v * 0.453592 }, func(v float64) float64 { return v / 0.453592 }},
	"oz": {Mass, func(v float64) float64 { return v * 0.0283495 }, func(v float64) float64 { return v / 0.0283495 }},

	"c":  {Temperature, func(v float64) float64 { return v }, func(v float64) float64 { return v }},
	"f":  {Temperature, func(v float64) float64 { return (v - 32) * 5 / 9 }, func(v float64) float64 { return v*9/5 + 32 }},
	"k":  {Temperature, func(v float64) float64 { return v - 273.15 }, func(v float64) float64 { return v + 273.15 }},
}

func convert(value float64, from, to string) (float64, error) {
	fromUnit, ok1 := units[from]
	toUnit, ok2 := units[to]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("unknown unit(s): %s or %s", from, to)
	}
	if fromUnit.Category != toUnit.Category {
		return 0, fmt.Errorf("incompatible units: %s (%s) → %s (%s)",
			from, fromUnit.Category, to, toUnit.Category)
	}

	baseValue := fromUnit.ToBase(value)
	result := toUnit.FromBase(baseValue)
	return result, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: " + os.Args[0] + " <value> <from_unit> <to_unit>")
		fmt.Println("Example: " + os.Args[0] + " 100 f c")
		return
	}

	val, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number:", os.Args[1])
		return
	}

	from := strings.ToLower(os.Args[2])
	to := strings.ToLower(os.Args[3])

	result, err := convert(val, from, to)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", val, from, result, to)
}

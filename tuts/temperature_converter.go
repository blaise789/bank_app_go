package main

import (
	"errors"
	"fmt"
)

func temperatureConverter(temperature float64, currentUnit string, targetUnit string) (float64,error ){

	switch currentUnit {
	case "C":
		switch targetUnit {
		case "F":
			temperature = (temperature * 9 / 5) + 32
		case "K":
			temperature += 273.15
		default:
			fmt.Println("invalid target unit")

		}
		return temperature,nil

	case "F":
		switch targetUnit {
		case "C":
			temperature = (temperature - 32) * 5 / 9
		case "K":
			temperature = (temperature - 273.15)

		}
		return temperature,nil
	case "K":
		switch targetUnit {
		case "C":
			temperature -= 273.15
		case "F":
			temperature = (temperature * 9 / 5) - 459.67
		}
		return temperature,nil
	default:
		fmt.Println("invalid temperature unit")
        return 0,errors.New("invalid target unit")

	}

}

package calculate

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Calculate(number int64, isLogNeed bool) int64 {
	if isLogNeed {
		log.Printf("Start calculations...\nCalculate <%d>!", number)
		if number <= 0 {
			log.Fatal("Number less then 0!")
		}

		result := factorial(number)

		if result <= 0 {
			log.Fatal("Too big number!")
		}

		log.Printf("Calculations complete!")
		return factorial(number)
	} else {
		if number <= 0 {
			fmt.Print("Number less then 0!")
			return -1
		}
		result := factorial(number)
		if result <= 0 {
			fmt.Print("Too big number!")
			return -1
		}
		return result
	}
}

func factorial(number int64) int64 {
	var result int64 = 1
	for i := int64(1); i < number+1; i++ {
		result *= i
	}
	return result
}

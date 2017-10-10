package main

import (
	"fmt"
)

func main() {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tenDigits := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	digitsCharCount := 0
	for _, v := range digits {
		digitsCharCount += len(v)
	}

	fmt.Println("digitsCharCount:", digitsCharCount)

	tenDigitsCharCount := 0
	for _, v := range tenDigits {
		tenDigitsCharCount += len(v)
	}

	fmt.Println("tenDigitsCharCount:", tenDigitsCharCount)

	oneHundredCharCount := digitsCharCount + tenDigitsCharCount
	for _, v := range tens {
		oneHundredCharCount += len(v)*10 + digitsCharCount
	}

	fmt.Println("oneHundredCharCount:", oneHundredCharCount)

	sumCharCount := oneHundredCharCount
	for _, v := range digits {
		sumCharCount += len(v+" hundred") - 1
		sumCharCount += 99*(len(v+" hundred and")-2) + sumCharCount
	}
	fmt.Println(sumCharCount)
	sumCharCount += len("one thousand") - 1

	fmt.Println(sumCharCount)
}

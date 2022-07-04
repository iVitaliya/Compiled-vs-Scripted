package main

import (
	"fmt"
	"math"
	"time"

	"github.com/iVitaliya/colors-go"
)

func Print(message []string) {
	for i := 0; i < len(message); i++ {
		fmt.Print(message[i])
	}

	fmt.Print("\n")
}

func ReverseArray(array []int) []int {
	var arr []int = []int{}

	for i := len(array) - 1; i != 0; i-- {
		var value int = array[i]
		arr = append(arr, value)
	}

	return arr
}

func LoopThru(array []string) {
	for _, v := range array {
		Print([]string{
			colors.Gray("["), colors.BrightBlue("FOR-OF"), " ", colors.Dim(colors.BrightYellow("Loop")), colors.Gray("]"),
			" Item: ", v,
		})
	}

	Print([]string{""})

	for i := 0; i < len(array); i++ {
		Print([]string{
			colors.Gray("["), colors.BrightBlue("FOR-INDEX"), " ", colors.Dim(colors.BrightYellow("Loop")), colors.Gray("]"),
			" Item: ", array[i],
		})
	}

	Print([]string{""})

	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Print([]string{
		fmt.Sprint(ReverseArray(m)),
		"\n",
	})
}

func TimeDevider(timer_one int, timer_two int) int {
	return int(math.Abs((float64(timer_one) - float64(timer_two))))
}

func main() {
	timer_one := time.Now().Local().UnixMilli()

	LoopThru([]string{
		"This",
		"is",
		"that",
		"and",
		"that",
		"is",
		"this",
		"and",
		"I",
		"am",
		"not",
		"possible",
	})

	timer_two := time.Now().Local().UnixMilli()

	Print([]string{
		colors.Gray("["), colors.Red("DEVISION"), colors.Gray("]"),
		" ", fmt.Sprint(TimeDevider(int(timer_one), int(timer_two))), "ms",
	})
}

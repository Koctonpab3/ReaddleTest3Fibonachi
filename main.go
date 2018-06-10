package main


import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	Timer := time.NewTimer(10 * time.Second)
	var (
		errCounter  int = 0
		victCounter int = 0
		STEP        int = 0
		text        int = -1
	)

	for {
		go timerProcess(Timer, &STEP)
		fmt.Println("Enter ", STEP, " fibonachi number: ")
		n, err := fmt.Scanf("%d\n", &text)
		if err != nil || n != 1 {
			// handle invalid input
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Error input ", n, err)
			continue
		}

		if text != fibonachi(STEP) {
			fmt.Println("Wrong input :(")
			mapA, _ := json.Marshal(map[int]int{STEP: fibonachi(STEP)})
			fmt.Println(string(mapA))
			Timer.Reset(10 * time.Second)
			errCounter++
			victCounter = 0
		}
		if text == fibonachi(STEP) {
			fmt.Println("Right input :)")
			Timer.Reset(10 * time.Second)
			victCounter++
			errCounter = 0
		}

		fmt.Println("\n---STATUS---")
		fmt.Println("Errors:", errCounter)
		fmt.Println("Good:", victCounter, "\n")

		STEP++

		if errCounter > 2 {
			fmt.Println("You are loose. Wrong trying:", errCounter)
			break
		}
		if victCounter > 9 {
			fmt.Println("You are WIN. Right trying:", victCounter)
			break
		}

	}

}
func timerProcess(Timer *time.Timer, step *int) {

	<-Timer.C
	var stepTmp int
	stepTmp = *step
	mapA, _ := json.Marshal(map[int]int{stepTmp: fibonachi(stepTmp)})
	fmt.Println(string(mapA))
	//fmt.Println("{\"", stepTmp, "\":\"", fibonachi(stepTmp), "\"}")
	*step++
	Timer.Reset(10 * time.Second)
	go timerProcess(Timer, step)
	fmt.Println("Enter ", *step, " fibonachi number: ")
}

func fibonachi(number int) int {
	if number == 0 || number == 1 {
		return number
	}

	return fibonachi(number-2) + fibonachi(number-1)
}
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}

func main() {
	for {
		var workNumber int
		fmt.Print("\033[H\033[2J")
		fmt.Print("Work time: ")
		_, err := fmt.Scanf("%d", &workNumber)
		check(err)

		workDuration := workNumber * 60

		for i := range workDuration {
			fmt.Print("\033[H\033[2J")
			fmt.Print(i)
			time.Sleep(1 * time.Second)
		}

		workDoneNotif := strconv.Itoa(workNumber) + " minute/s is over"

		err = beeep.Notify("Time to rest", workDoneNotif, "")
		check(err)

		var restNumber int
		fmt.Print("\033[H\033[2J")
		fmt.Print("Resting time: ")
		_, err = fmt.Scanf("%d", &restNumber)
		check(err)

		restDuration := restNumber * 60

		for i := range restDuration {
			fmt.Print("\033[H\033[2J")
			fmt.Print(i)
			time.Sleep(1 * time.Second)
		}

		restDoneNotif := strconv.Itoa(restNumber) + " minute/s is over"

		err = beeep.Notify("Time to work", restDoneNotif, "")
		check(err)
	}
}

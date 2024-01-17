package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var csvFileName string = "quiz_game/problems.csv" // default
	var quizPartArgs string
	var partNo uint8
	var positivePoints, negativePoints uint8
	const TOTAL_PROBLEMS uint8 = 12

	fmt.Sscanf(strings.Join(os.Args[1:], " "), "%s %v", &quizPartArgs, &partNo)
	if quizPartArgs != "--part" || partNo < 1 {
		fmt.Println("Enter a part as argument to continue")
		os.Exit(69)
	}

	fmt.Print("Enter csv filename [default -> quiz_game/problems.csv, press enter to continue]: ")
	fmt.Scanln(&csvFileName)
	if strings.TrimSpace(csvFileName) == "" {
		fmt.Println("A filename is required")
		os.Exit(69)
	}

	file, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(69)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error: ", err.Error())
		os.Exit(69)
	}

	if partNo == 1 {
		RunQuizPart1(*scanner, &positivePoints, &negativePoints)
		printScore(TOTAL_PROBLEMS, positivePoints, negativePoints)

	} else if partNo == 2 {
		waitingSeconds := 10
		timer := time.NewTimer(time.Second * time.Duration(waitingSeconds))
		go RunQuizPart2(*scanner, &positivePoints, &negativePoints)
		<-timer.C
		printScore(TOTAL_PROBLEMS, positivePoints, negativePoints)
	} else {
		fmt.Println("Invalid part number")
		os.Exit(69)
	}
}

func RunQuizPart1(scanner bufio.Scanner, positivePoints, negativePoints *uint8) {
	var problem string
	var givenAnswer uint8

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s", &problem)

		splitProblem := strings.Split(problem, ",")
		question := splitProblem[0]
		actualAnswer, err := strconv.Atoi(splitProblem[1])

		if err != nil {
			*negativePoints++
			continue
		}

		fmt.Printf("What does %s evaluate to?: ", question)
		fmt.Scanln(&givenAnswer)

		if uint8(actualAnswer) == givenAnswer {
			*positivePoints++
		} else {
			*negativePoints++
		}
	}
}

func RunQuizPart2(scanner bufio.Scanner, positivePoints, negativePoints *uint8) {
	var problem string
	var givenAnswer uint8

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s", &problem)

		splitProblem := strings.Split(problem, ",")
		question := splitProblem[0]
		actualAnswer, err := strconv.Atoi(splitProblem[1])

		if err != nil {
			*negativePoints++
			continue
		}

		fmt.Printf("What does %s evaluate to?: ", question)
		fmt.Scanln(&givenAnswer)

		if uint8(actualAnswer) == givenAnswer {
			*positivePoints++
		} else {
			*negativePoints++
		}
	}
}

func printScore(totalPoints, positivePoints, negativePoints uint8) {
	fmt.Println("\n\nTIME UPP :)")
	fmt.Printf("There were %d total questions, you got %d right and %d wrong answer(s).\n", totalPoints, positivePoints, negativePoints)
}

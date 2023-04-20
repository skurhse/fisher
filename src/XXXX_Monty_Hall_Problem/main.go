package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type messages struct {
	GameIntro     string
	FirstChoice   string
	DoorRemoval   string
	SecondChoice  string
	GameOutcome   string
	ReplayInquiry string
}

var msgs = messages{
	GameIntro:     "{{ .welcome }}In front of you are three doors.\n",
	FirstChoice:   "Which door would you like to choose? (1-3)\n",
	DoorRemoval:   "Well, before you make a decision, door #{{ .door }} doesn’t hold a prize!\n",
	SecondChoice:  "Now, would you like to stay on door #{{ .chosenDoor }}, or swap to door #{{ .survivingDoor }}? (stay/swap)\n",
	GameOutcome:   "You’ve {{ .outcome }}! The door that held the prize was door #{{ .door }}!\n",
	ReplayInquiry: "Would you like to play again? (y/n)\n",
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var err error

	for letsPlay, welcome, playAgain := true, true, false; letsPlay; letsPlay = playAgain {
		err = playGame(welcome)
		if err != nil {
			err = fmt.Errorf("playing game: %w", err)
			break
		}
		welcome = false

		playAgain, err = confirmChoice(msgs.ReplayInquiry, "y", "n")
		if err != nil {
			err = fmt.Errorf("inquiring replay: %w", err)
			break
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func playGame(welcome bool) (err error) {
	rand.Seed(time.Now().UnixNano())
	prizeDoor := rand.Intn(3) + 1

	availableDoors := map[int]bool{1: true, 2: true, 3: true}

	var welcomeMsg string
	if welcome {
		welcomeMsg = "Welcome! "
	} else {
		welcomeMsg = ""
	}

	templ := template.Must(template.New("GameIntro").Parse(msgs.GameIntro))
	err = templ.Execute(os.Stdout, map[string]interface{}{
		"welcome": welcomeMsg,
	})
	if err != nil {
		err = fmt.Errorf("formatting game intro: %w", err)
		return
	}

	firstChoice, err := pickFirstChoice()
	if err != nil {
		err = fmt.Errorf("picking first choice: %w", err)
		return
	}
	delete(availableDoors, firstChoice)

	remainingDoors := reflect.ValueOf(availableDoors).MapKeys()
	doorToRemove := remainingDoors[rand.Intn(len(remainingDoors))].Interface().(int)
	delete(availableDoors, doorToRemove)

	swappableDoor := reflect.ValueOf(availableDoors).MapKeys()[0].Interface().(int)

	templ = template.Must(template.New("DoorRemoval").Parse(msgs.DoorRemoval))
	err = templ.Execute(os.Stdout, map[string]interface{}{
		"door": doorToRemove,
	})
	if err != nil {
		err = fmt.Errorf("formatting door removal: %w", err)
		return
	}

	templ = template.Must(template.New("SecondChoice").Parse(msgs.SecondChoice))
	buf := &bytes.Buffer{}
	err = templ.Execute(buf, map[string]interface{}{
		"chosenDoor":    firstChoice,
		"survivingDoor": swappableDoor,
	})
	if err != nil {
		err = fmt.Errorf("formatting second choice: %w", err)
		return
	}

	stay, err := confirmChoice(buf.String(), "stay", "swap")
	if err != nil {
		err = fmt.Errorf("picking second choice: %w", err)
		return
	}

	var secondChoice int
	if stay {
		secondChoice = firstChoice
	} else {
		secondChoice = swappableDoor
	}

	var outcome string
	if secondChoice == prizeDoor {
		outcome = "won"
	} else {
		outcome = "lost"
	}

	templ = template.Must(template.New("GameOutcome").Parse(msgs.GameOutcome))
	err = templ.Execute(os.Stdout, map[string]interface{}{
		"outcome": outcome,
		"door":    prizeDoor,
	})
	if err != nil {
		err = fmt.Errorf("formatting game outcome: %w", err)
		return
	}

	return
}

func pickFirstChoice() (choice int, err error) {
	fmt.Printf(msgs.FirstChoice)

	var retrieved bool
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)

		choice, err = strconv.Atoi(input)

		if err == nil && choice >= 1 && choice <= 3 {
			retrieved = true
			break
		}

		fmt.Printf(msgs.FirstChoice)
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("reading standard input: %w", err)
		return
	}

	if !retrieved {
		err = fmt.Errorf("no input received")
		return
	}

	return
}

func confirmChoice(msg string, yes string, no string) (confirmation bool, err error) {
	fmt.Print(msg)

	var confirmed bool

scan:
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)

		switch input {
		case yes:
			confirmation, confirmed = true, true
			break scan
		case no:
			confirmation, confirmed = false, true
			break scan
		default:
			fmt.Print(msg)
		}
	}

	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("reading standard input: %w", err)
		return
	}

	if !confirmed {
		err = fmt.Errorf("no input received")
		return
	}

	return
}

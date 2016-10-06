package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Frage den Spieler nach dem Namen
// Frage nach der Anzahl der Aufgaben
// Frage nach den Rechenarten
// Frage nach der Höhe der Zahlen
type userProfile struct {
	name     string
	number   int
	calculus string
	max      int
}

func (u *userProfile) askUserName() error {

	fmt.Printf("Hallo!\nWie heißt du?\n>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	u.name = scanner.Text()
	fmt.Printf("Hallo %s!\n", u.name)

	return nil
}

func (u *userProfile) askNumber() error {
	for {
		fmt.Printf("Wieviele Aufgaben darf ich dir stellen?\n>")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		val := scanner.Text()
		myint, err := strconv.Atoi(val)
		if err == nil {
			u.number = myint
			return nil
		}
		fmt.Println("Das war keine Zahl!")
	}
}

func (u *userProfile) askCalculus() error {
	// Hier wird noch nicht gefragt, ist nur vorbereitung
	return nil
}

func (u *userProfile) askMaxHeightNumber() error {
	return nil
}

// Errechnet eine Zufallszahl anhand von Min- und Max-Wert
func calcRandVal(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	myval := min + rand.Intn((max+1)-min)
	return myval
}

func calcPercentage(u *userProfile) {
	// ersten Prozentwert errechnen
	a := calcRandVal(10, 90)
	anew := float64(a)
	max := float64(u.max)
	firstval2 := (max / 100.) * anew
	val1 := int(firstval2)

	// Zweiten Prozentwert errechnen
	secondMax := 100 - a
	b := calcRandVal(10, secondMax)
	bnew := float64(b)
	secval := (max / 100.) * bnew
	val2 := int(secval)

	fmt.Printf("%d + %d = ?\n", val1, val2)

}

func createTasks(u *userProfile) {

	for i := 0; i < u.number; i++ {
		calcPercentage(u)
	}
}

//--------------------------------------
func main() {
	fmt.Printf("\n~~~~~~ Willkommen bei Mathematica ~~~~~\n\n")

	// Namen abfragen
	u := userProfile{}
	/*err := u.askUserName()
	if err != nil {
		panic(err)
	}*/

	// Anzahl abfragen
	err := u.askNumber()
	if err != nil {
		panic(err)
	}

	// Rechenarten abfragen
	u.calculus = "+-"

	// Höhe der Zahlen abfragen
	u.max = 50

	createTasks(&u)
	//

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().Unix())
	dayOfWeek := rand.Intn(8) + 1
	fmt.Println("Day of Week", dayOfWeek)

	var result string
	switch dayOfWeek {
	case 1:
		result = "It's Sunday. Going riding!"
	case 2:
		result = "It's Monday. Got to setup the tone for the week!"
	case 3:
		result = "It's Tuesday. Did I accomplish my Monday's objectives?"
	case 4:
		result = "It's Wednesday."
	case 5:
		result = "It's Thursday. Don't forget to update your Lattice"
	case 6:
		result = "It's Friday. Winding up and preparing for church"
	case 7:
		result = "It's Saturday. Going to church!"
	default:
		result = "Oh! Oh! Is that a day?"
	}

	fmt.Println(result)
}

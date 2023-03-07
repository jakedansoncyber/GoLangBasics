package main

import "fmt"

func main() {

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Zombieland"] = 9

	recentHead2Wins := map[string]int{
		"Sunderland": 6,
		"Zombieland": 9,
	}

	for _, i := range recentHead2Wins {
		fmt.Println(i)
	}

	// delete an entry
	delete(leagueTitles, "Sunderland")

	leagueTitles["Happy Gilmore"] = 8

}

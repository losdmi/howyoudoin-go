package main

import (
	"fmt"
	"os"

	"github.com/losdmi/howyoudoin-go/internal"
)

func main() {
	nextEpisode, err := internal.GetNextEpisode()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	fmt.Println(nextEpisode)
	fmt.Println()

	fmt.Println("Press Enter to exit...")
	var input string
	_, _ = fmt.Scanln(&input)
}

//nolint:unused
func generateEpisodesList() {
	seasonsWithEpisodes := map[int]int{
		1:  24,
		2:  24,
		3:  25,
		4:  24,
		5:  24,
		6:  25,
		7:  24,
		8:  24,
		9:  23,
		10: 17,
	}

	for season := 1; season <= 10; season++ {
		for episode := 1; episode <= seasonsWithEpisodes[season]; episode++ {
			fmt.Printf("\"s%02de%02d\",\n", season, episode)
		}
	}
}

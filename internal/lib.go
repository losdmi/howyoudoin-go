package internal

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

const DB_FILENAME = "seen_episodes.txt"

func GetNextEpisode() (string, error) {
	seenEpisodes, err := readDBFromFile(DB_FILENAME)
	if err != nil {
		return "", err
	}

	selectedEpisode, err := selectNextEpisode(seenEpisodes)
	if err != nil {
		return "", err
	}

	seenEpisodes = append(seenEpisodes, selectedEpisode)
	err = saveDBToFile(seenEpisodes, DB_FILENAME)
	if err != nil {
		return "", err
	}

	return selectedEpisode, nil
}

func readDBFromFile(dbFilename string) ([]string, error) {
	bytes, err := os.ReadFile(dbFilename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	asString := strings.TrimSpace(string(bytes))
	if asString == "" {
		return nil, nil
	}

	parts := strings.Split(asString, "\n")

	result := make([]string, 0, len(parts))
	for i := len(parts) - 1; 0 <= i; i-- {
		result = append(result, parts[i])
	}

	return result, nil
}

func selectNextEpisode(seenEpisodes []string) (string, error) {
	seenMap := make(map[string]struct{})
	for _, seenEpisode := range seenEpisodes {
		seenMap[seenEpisode] = struct{}{}
	}

	episodesCopy := EPISODES
	rand.Shuffle(len(episodesCopy), func(i, j int) {
		episodesCopy[i], episodesCopy[j] = episodesCopy[j], episodesCopy[i]
	})

	for _, episode := range episodesCopy {
		if _, exists := seenMap[episode]; !exists {
			return episode, nil
		}
	}

	return "", fmt.Errorf("no unseen episodes")
}

func saveDBToFile(seenEpisodes []string, dbFilename string) error {
	if len(seenEpisodes) == 0 {
		err := os.WriteFile(dbFilename, nil, 0666)
		if err != nil {
			return err
		}

		return nil
	}

	seenEpisodesReversed := make([]string, 0, len(seenEpisodes))

	for i := len(seenEpisodes) - 1; 0 <= i; i-- {
		seenEpisodesReversed = append(seenEpisodesReversed, seenEpisodes[i])
	}

	err := os.WriteFile(dbFilename, []byte(strings.Join(seenEpisodesReversed, "\n")+"\n"), 0666)
	if err != nil {
		return err
	}

	return nil
}

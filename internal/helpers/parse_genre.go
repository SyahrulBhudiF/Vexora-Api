package helpers

import "strings"

func ParseGenres(genres string) []string {
	if genres == "" {
		return []string{}
	}
	genreSlice := strings.Split(genres, ",")

	for i, genre := range genreSlice {
		genreSlice[i] = strings.TrimSpace(genre)
	}

	return genreSlice
}

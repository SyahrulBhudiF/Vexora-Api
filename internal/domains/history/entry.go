package history

import (
	entity3 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	entity2 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
	"github.com/zmb3/spotify"
)

func CreateHistoryEntry(user *entity2.User, mood string, recommendations *entity3.MoodResponse) *entity3.History {
	history := &entity3.History{
		UserUUID: user.UUID,
		Mood:     mood,
		Music:    make([]entity.Music, len(recommendations.Music)),
	}

	for i, track := range recommendations.Music {
		history.Music[i] = entity.Music{
			ID:        track.ID,
			MusicName: track.MusicName,
			Path:      track.Path,
			Artist:    track.Artist,
			Thumbnail: track.Thumbnail,
		}
	}

	return history
}

var availableGenres = []string{
	"genre:pop",
	"genre:rock",
	"genre:hip hop",
	"genre:electronic",
	"genre:indie",
	"genre:jazz",
	"genre:classical",
	"genre:country",
	"genre:r&b",
	"genre:alternative",
	"genre:blues",
	"genre:reggae",
	"genre:folk",
	"genre:latin",
	"genre:metal",
	"genre:punk",
	"genre:ambient",
}

var GenreMoodTrackAttributes = map[string][]string{
	"sad":     {"Blues", "Classical", "Ballads", "Indie Folk", "Sad Indie"},
	"happy":   {"Pop", "Dance", "Funk", "Reggae", "Upbeat Indie"},
	"angry":   {"Rock", "Metal", "Punk", "Hardcore", "Hip-Hop"},
	"neutral": {"Ambient", "Chillout", "Lo-fi", "Indie", "Acoustic"},
}

var MoodTrackAttributes = map[string]spotify.TrackAttributes{
	"sad": *spotify.NewTrackAttributes().
		MaxDanceability(0.65). // Cluster 3
		MaxEnergy(0.42).       // Cluster 3
		MaxValence(0.30).      // Cluster 3
		MaxAcousticness(0.32), // Cluster 3

	"happy": *spotify.NewTrackAttributes().
		MinDanceability(0.70). // Cluster 1
		MaxEnergy(0.65).       // Cluster 1
		MinValence(0.59).      // Cluster 1
		MaxSpeechiness(0.50),  // Cluster 1

	"angry": *spotify.NewTrackAttributes().
		MaxDanceability(0.44). // Cluster 2
		MinEnergy(0.84).       // Cluster 2
		MaxValence(0.28).      // Cluster 2
		MaxLiveness(0.32),     // Cluster 2

	"neutral": *spotify.NewTrackAttributes().
		MaxDanceability(0.56).     // Cluster 0
		MaxEnergy(0.80).           // Cluster 0
		MaxValence(0.27).          // Cluster 0
		MaxInstrumentalness(0.81), // Cluster 0
}

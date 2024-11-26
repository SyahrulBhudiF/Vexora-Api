package history

import (
	entity3 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/history/entity"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/domains/music/entity"
	entity2 "github.com/SyahrulBhudiF/Vexora-Api/internal/domains/user/entity"
)

func CreateHistoryEntry(user *entity2.User, mood string, recommendations *entity3.PlaylistResponse) *entity3.History {
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

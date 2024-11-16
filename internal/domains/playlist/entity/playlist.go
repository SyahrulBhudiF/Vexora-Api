package entity

type RandomPlaylist struct {
	PlaylistName string `json:"playlist_name"`
	Artist       string `json:"artist"`
	Path         string `json:"path"`
	Thumbnail    string `json:"thumbnail"`
}

type PlaylistResponse struct {
	Playlists []RandomPlaylist `json:"playlists"`
}

func NewRandomPlaylist(name, artist, path, thumbnail string) *RandomPlaylist {
	return &RandomPlaylist{
		PlaylistName: name,
		Artist:       artist,
		Path:         path,
		Thumbnail:    thumbnail,
	}
}

func NewPlaylistResponse(playlists []RandomPlaylist) *PlaylistResponse {
	return &PlaylistResponse{
		Playlists: playlists,
	}
}

type Music struct {
	Name      string `json:"music_name"`
	Artist    string `json:"artist"`
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
}

type MusicResponse struct {
	Musics []Music `json:"musics"`
}

func NewMusic(name, artist, path, thumbnail string) *Music {
	return &Music{
		Name:      name,
		Artist:    artist,
		Path:      path,
		Thumbnail: thumbnail,
	}
}

func NewMusicResponse(musics []Music) *MusicResponse {
	return &MusicResponse{
		Musics: musics,
	}
}

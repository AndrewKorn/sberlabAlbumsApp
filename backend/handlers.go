package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (dbe *DBEngine) getGroupsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups := dbe.getGroups()
	err := json.NewEncoder(w).Encode(groups)
	if err != nil {
		return
	}
}

func (dbe *DBEngine) getGroupAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	albums := dbe.getGroupsAlbums(params["group_name"])
	_ = json.NewEncoder(w).Encode(albums)
}

func (dbe *DBEngine) getAlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	songs := dbe.getAlbumSongs(params["group_name"], params["album_name"])

	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		return
	}
}

func (dbe *DBEngine) createAlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	album := &Album{}
	_ = json.NewDecoder(r.Body).Decode(&album)
	album = dbe.createAlbum(params["group_name"], album.AlbumName)

	err := json.NewEncoder(w).Encode(album)
	if err != nil {
		return
	}
}

func (dbe *DBEngine) createSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	song := &Song{}
	_ = json.NewDecoder(r.Body).Decode(&song)
	song = dbe.createSong(params["group_name"], params["album_name"], song)

	err := json.NewEncoder(w).Encode(&song)
	if err != nil {
		return
	}
}

func (dbe *DBEngine) deleteAlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	album := dbe.deleteAlbum(params["group_name"], params["album_name"])
	err := json.NewEncoder(w).Encode(album)
	if err != nil {
		return
	}
}

func (dbe *DBEngine) deleteSongHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["song_id"])
	song := dbe.deleteSong(params["group_name"], params["album_name"], uint(id))
	err := json.NewEncoder(w).Encode(song)
	if err != nil {
		return
	}
}

func (dbe *DBEngine) createGroupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	group := &Group{}
	_ = json.NewDecoder(r.Body).Decode(&group)
	group.Albums = make([]Album, 0)
	group = dbe.createGroup(group)

	err := json.NewEncoder(w).Encode(group)
	if err != nil {
		return
	}
}

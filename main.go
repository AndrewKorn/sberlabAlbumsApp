package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Song struct {
	ID       string `json:"id"`
	Artist   string `json:"artist"`
	SongName string `json:"song_name"`
}

type Album struct {
	ID        string `json:"id"`
	AlbumName string `json:"album_name"`
	Songs     []Song `json:"songs"`
}

var albums []Album

func main() {
	r := mux.NewRouter()
	var songs []Song
	songs = append(songs, Song{"1", "kizaru", "train"}, Song{"2", "yanix", "moscow"})
	albums = append(albums, Album{"1", "FDO", songs})
	r.HandleFunc("/albums", getAlbums).Methods("GET")
	r.HandleFunc("/albums/{album_id}", getAlbum).Methods("GET")
	r.HandleFunc("/albums/{album_id}/{song_id}", getSong).Methods("GET")
	r.HandleFunc("/albums", createAlbum).Methods("POST")
	r.HandleFunc("/albums/{album_id}", createSong).Methods("POST")
	r.HandleFunc("/albums/{album_id}", updateAlbum).Methods("PUT")
	r.HandleFunc("/albums/{album_id}", deleteAlbum).Methods("DELETE")
	r.HandleFunc("/albums/{album_id}/{song_id}", deleteSong).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":80", r))
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.ID == params["album_id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Album{})
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.ID == params["album_id"] {
			for _, songItem := range item.Songs {
				if songItem.ID == params["song_id"] {
					json.NewEncoder(w).Encode(songItem)
					return
				}
			}
		}
	}
	json.NewEncoder(w).Encode(&Song{})
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var album Album
	_ = json.NewDecoder(r.Body).Decode(&album)
	album.ID = strconv.Itoa(rand.Intn(1000000))
	albums = append(albums, album)
	json.NewEncoder(w).Encode(album)
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.ID == params["album_id"] {
			var song Song
			_ = json.NewDecoder(r.Body).Decode(&song)
			song.ID = strconv.Itoa(rand.Intn(1000000))
			i, _ := strconv.Atoi(item.ID)
			albums[i].Songs = append(albums[i].Songs, song)
			json.NewEncoder(w).Encode(song)
			return
		}
	}
	json.NewEncoder(w).Encode(&Song{})
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range albums {
		if item.ID == params["album_id"] {
			albums = append(albums[:index], albums[index+1:]...)
			var album Album
			_ = json.NewDecoder(r.Body).Decode(&album)
			album.ID = params["album_id"]
			albums = append(albums, album)
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	json.NewEncoder(w).Encode(albums)
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range albums {
		if item.ID == params["album_id"] {
			albums = append(albums[:index], albums[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(albums)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.ID == params["album_id"] {
			for index, songItem := range item.Songs {
				if songItem.ID == params["song_id"] {
					item.Songs = append(item.Songs[:index], item.Songs[index+1])
					break
				}
			}
		}
	}
	json.NewEncoder(w).Encode(albums)
}

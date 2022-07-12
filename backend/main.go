package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
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

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(albums)
	if err != nil {
		return
	}
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.AlbumName == params["album_name"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Album{})
	if err != nil {
		return
	}
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.AlbumName == params["album_name"] {
			for _, songItem := range item.Songs {
				if songItem.ID == params["song_id"] {
					err := json.NewEncoder(w).Encode(songItem)
					if err != nil {
						return
					}
					return
				}
			}
		}
	}
	err := json.NewEncoder(w).Encode(&Song{})
	if err != nil {
		return
	}
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var album Album
	_ = json.NewDecoder(r.Body).Decode(&album)
	album.ID = strconv.Itoa(rand.Intn(1000000))
	album.Songs = []Song{}
	albums = append(albums, album)
	err := json.NewEncoder(w).Encode(album)
	if err != nil {
		return
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.AlbumName == params["album_name"] {
			var song Song
			_ = json.NewDecoder(r.Body).Decode(&song)
			song.ID = strconv.Itoa(rand.Intn(1000000))
			item.Songs = append(item.Songs, song)
			println(item.Songs[0].SongName, item.Songs[1].SongName, item.Songs[2].SongName, item.Songs[3].SongName)
			err := json.NewEncoder(w).Encode(song)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Song{})
	if err != nil {
		return
	}
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range albums {
		if item.AlbumName == params["album_name"] {
			albums = append(albums[:index], albums[index+1:]...)
			var album Album
			_ = json.NewDecoder(r.Body).Decode(&album)
			albums = append(albums, album)
			err := json.NewEncoder(w).Encode(album)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(albums)
	if err != nil {
		return
	}
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range albums {
		if item.AlbumName == params["album_name"] {
			albums = append(albums[:index], albums[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(albums)
	if err != nil {
		return
	}
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.AlbumName == params["album_name"] {
			for index, songItem := range item.Songs {
				if songItem.ID == params["song_id"] {
					item.Songs = append(item.Songs[:index], item.Songs[index+1:]...)
					err := json.NewEncoder(w).Encode(item.Songs)
					if err != nil {
						return
					}
					break
				}
			}
		}
	}
}

var albums []Album

func main() {
	r := mux.NewRouter()
	var songs1 []Song
	songs1 = append(songs1, Song{"1", "kizaru", "train wreck"}, Song{"2", "yanix", "moscow"})
	albums = append(albums, Album{"1", "First Day Out", songs1})

	var songs2 []Song
	songs2 = append(songs2, Song{"1", "OGBuda", "Сайфер"}, Song{"2", "163ONMYNECK", "выключатель"}, Song{"3", "Mayot", "море"})
	albums = append(albums, Album{"2", "Meloners", songs2})

	r.HandleFunc("/albums", getAlbums).Methods("GET")
	r.HandleFunc("/albums/{album_name}", getAlbum).Methods("GET")
	r.HandleFunc("/albums/{album_name}/{song_id}", getSong).Methods("GET")
	r.HandleFunc("/albums", createAlbum).Methods("POST")
	r.HandleFunc("/albums/{album_name}", createSong).Methods("POST")
	r.HandleFunc("/albums/{album_name}", updateAlbum).Methods("PUT")
	r.HandleFunc("/albums/{album_name}", deleteAlbum).Methods("DELETE")
	r.HandleFunc("/albums/{album_name}/{song_id}", deleteSong).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

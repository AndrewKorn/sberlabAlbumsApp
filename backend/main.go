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
	ID       string `json:"song_id"`
	Artist   string `json:"artist"`
	SongName string `json:"song_name"`
	Duration string `json:"duration"`
}

type Album struct {
	ID        string `json:"id"`
	AlbumName string `json:"album_name"`
	Songs     []Song `json:"songs"`
}

type Group struct {
	ID        string  `json:"group_id"`
	GroupName string  `json:"group_name"`
	Albums    []Album `json:"albums"`
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(groups)
	if err != nil {
		return
	}
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range groups {
		if item.GroupName == params["group_name"] {
			_ = json.NewEncoder(w).Encode(item.Albums)
			return
		}
	}
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, itemGroup := range groups {
		if itemGroup.GroupName == params["group_name"] {
			for _, itemAlbum := range itemGroup.Albums {
				if itemAlbum.AlbumName == params["album_name"] {
					err := json.NewEncoder(w).Encode(itemAlbum)
					if err != nil {
						return
					}
					return
				}
			}
		}
	}

	err := json.NewEncoder(w).Encode(&Album{})
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

	params := mux.Vars(r)
	for _, item := range groups {
		if item.GroupName == params["group_name"] {
			item.Albums = append(item.Albums, album)

			for index, it := range groups {
				if it.GroupName == params["group_name"] {
					groups = append(groups[:index], groups[index+1:]...)
					break
				}
			}

			groups = append(groups, item)
		}
	}
	err := json.NewEncoder(w).Encode(album)
	if err != nil {
		return
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, itemGroup := range groups {
		if itemGroup.GroupName == params["group_name"] {
			for _, itemAlbum := range itemGroup.Albums {
				if itemAlbum.AlbumName == params["album_name"] {
					var song Song
					_ = json.NewDecoder(r.Body).Decode(&song)
					itemAlbum.Songs = append(itemAlbum.Songs, song)

					for index, it := range itemGroup.Albums {
						if it.AlbumName == params["album_name"] {
							itemGroup.Albums = append(itemGroup.Albums[:index], itemGroup.Albums[index+1:]...)
							break
						}
					}

					itemGroup.Albums = append(itemGroup.Albums, itemAlbum)

					for index, it := range groups {
						if it.GroupName == params["group_name"] {
							groups = append(groups[:index], groups[index+1:]...)
							break
						}
					}

					groups = append(groups, itemGroup)
					err := json.NewEncoder(w).Encode(song)
					if err != nil {
						println("err")
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

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, itemGroup := range groups {
		if itemGroup.GroupName == params["group_name"] {
			for index, itemAlbum := range itemGroup.Albums {
				if itemAlbum.AlbumName == params["album_name"] {
					itemGroup.Albums = append(itemGroup.Albums[:index], itemGroup.Albums[index+1:]...)

					for i, it := range groups {
						if it.GroupName == params["group_name"] {
							groups = append(groups[:i], groups[i+1:]...)
							break
						}
					}

					groups = append(groups, itemGroup)
					err := json.NewEncoder(w).Encode(itemGroup.Albums)
					if err != nil {
						return
					}
					break
				}
			}
		}
	}
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, itemGroup := range groups {
		if itemGroup.GroupName == params["group_name"] {
			for _, itemAlbum := range itemGroup.Albums {
				if itemAlbum.AlbumName == params["album_name"] {
					for index, itemSong := range itemAlbum.Songs {
						if itemSong.ID == params["song_id"] {
							itemAlbum.Songs = append(itemAlbum.Songs[:index], itemAlbum.Songs[index+1:]...)

							for index, it := range itemGroup.Albums {
								if it.AlbumName == params["album_name"] {
									itemGroup.Albums = append(itemGroup.Albums[:index], itemGroup.Albums[index+1:]...)
									break
								}
							}

							itemGroup.Albums = append(itemGroup.Albums, itemAlbum)

							for index, it := range groups {
								if it.GroupName == params["group_name"] {
									groups = append(groups[:index], groups[index+1:]...)
									break
								}
							}

							groups = append(groups, itemGroup)

							err := json.NewEncoder(w).Encode(itemAlbum.Songs)
							if err != nil {
								return
							}
							break
						}
					}
				}
			}
		}
	}
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group Group
	_ = json.NewDecoder(r.Body).Decode(&group)
	group.Albums = []Album{}

	groups = append(groups, group)
	err := json.NewEncoder(w).Encode(group)
	if err != nil {
		return
	}
}

var groups []Group

func main() {
	r := mux.NewRouter()
	var songs1 []Song
	var albums []Album
	songs1 = append(songs1, Song{"1298045950", "kizaru", "train wreck", "1:59"}, Song{"1235590951", "yanix", "Москва", "2:06"})
	albums = append(albums, Album{"1", "First Day Out", songs1})

	var songs2 []Song
	songs2 = append(songs2, Song{"1045470445", "OGBuda", "Сайфер", "2:56"}, Song{"1217814655", "163ONMYNECK", "выключатель", "1:36"}, Song{"863371033", "Mayot", "море", "2:24"})
	albums = append(albums, Album{"2", "Meloners", songs2})

	groups = append(groups, Group{"123", "PSIHI", albums})

	r.HandleFunc("/", getGroups).Methods("GET")
	r.HandleFunc("/{group_name}/albums", getAlbums).Methods("GET")
	r.HandleFunc("/{group_name}/albums/{album_name}", getAlbum).Methods("GET")
	r.HandleFunc("/{group_name}/albums", createAlbum).Methods("POST")
	r.HandleFunc("/{group_name}/albums/{album_name}", createSong).Methods("POST")
	r.HandleFunc("/{group_name}/albums/{album_name}", deleteAlbum).Methods("DELETE")
	r.HandleFunc("/{group_name}/albums/{album_name}/{song_id}", deleteSong).Methods("DELETE")
	r.HandleFunc("/", createGroup).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":1337", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

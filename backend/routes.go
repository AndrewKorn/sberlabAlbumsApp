package main

import "github.com/gorilla/mux"

var RegisterRoutes = func(router *mux.Router, dbe *DBEngine) {
	router.HandleFunc("/", dbe.getGroupsHandler).Methods("GET")
	router.HandleFunc("/{group_name}/albums", dbe.getGroupAlbumsHandler).Methods("GET")
	router.HandleFunc("/{group_name}/albums", dbe.createAlbumHandler).Methods("POST")
	router.HandleFunc("/{group_name}/albums/{album_name}", dbe.deleteAlbumHandler).Methods("DELETE")
	router.HandleFunc("/{group_name}/albums/{album_name}", dbe.getAlbumHandler).Methods("GET")
	router.HandleFunc("/{group_name}/albums/{album_name}", dbe.createSongHandler).Methods("POST")
	router.HandleFunc("/{group_name}/albums/{album_name}/{song_id}", dbe.deleteSongHandler).Methods("DELETE")
	router.HandleFunc("/", dbe.createGroupHandler).Methods("POST")
}

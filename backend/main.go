package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//var groups []Group

func main() {
	dbConfig := DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "password",
		Name:     "postgres",
		Port:     "5432",
		SSLMode:  "disable",
		Tz:       "Asia/Novosibirsk",
	}

	dbe, _ := NewDBEngine(dbConfig)

	/*var songs1 []Song
	var albums []Album
	songs1 = append(songs1, Song{"1298045950", "kizaru", "train wreck", "1:59"}, Song{"1235590951", "yanix", "Москва", "2:06"})
	albums = append(albums, Album{"1", "First Day Out", songs1})
	var songs2 []Song
	songs2 = append(songs2, Song{"1045470445", "OGBuda", "Сайфер", "2:56"}, Song{"1217814655", "163ONMYNECK", "выключатель", "1:36"}, Song{"863371033", "Mayot", "море", "2:24"})
	albums = append(albums, Album{"2", "Meloners", songs2})
	groups = append(groups, Group{"123", "PSIHI", albums})*/

	router := mux.NewRouter()
	RegisterRoutes(router, dbe)
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Fatal(http.ListenAndServe(":1337", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

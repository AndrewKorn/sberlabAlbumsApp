package main

import (
	"encoding/json"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/gorilla/mux"
	"net/http"
)

func parseIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://soundcloud.com/search?q="+params["artist"]+"%20"+params["song_name"]+"", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			href, _ := r.HTMLDoc.Find("a.sound__coverArt").Attr("href")
			id, _ := GetSongID("https://soundcloud.com" + href + "")
			err := json.NewEncoder(w).Encode(id)
			if err != nil {
				return
			}
			println(id)
		},
	}).Start()

}

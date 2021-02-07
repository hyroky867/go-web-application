package main

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type poll struct {
	ID      bson.ObjectId  `bson:"_id" json:"id"`
	Title   string         `json:"title"`
	Options []string       `json:"options"`
	Results map[string]int `json:"results,omitempty"`
}

func handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlepollsGet(w, r)
		return
	case "POST":
		handlepollsPost(w, r)
		return
	case "DELETE":
		handlepollsDelete(w, r)
		return
	}
	// 未対応のHTTP利用メソッド
	respondHTTPErr(w, r, http.StatusNotFound)
}

func handlepollsGet(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}
func handlepollsPost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}
func handlepollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}

package main

import (
	"errors"
	"net/http"

	"gopkg.in/mgo.v2"
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
	db := GetVars(r, "db").(*mgo.Database)
	c := db.C("polls")
	var q *mgo.Query
	p := NewPath(r.URL.Path)
	if p.HasID() {
		// 特定の調査項目の詳細
		q = c.FindId(bson.ObjectIdHex(p.ID))
	} else {
		// すべての調査項目のリスト
		q = c.Find(nil)
	}

	var result []*poll
	if err := q.All(&result); err != nil {
		respondErr(w, r, http.StatusInternalServerError, err)
		return
	}
	respondErr(w, r, http.StatusInternalServerError, errors.New("未実装です"))
}

func handlepollsPost(w http.ResponseWriter, r *http.Request) {
	db := GetVars(r, "db").(*mgo.Database)
	c := db.C("polls")
	var p poll
	if err := decodeBody(r, &p); err != nil {
		respondErr(w, r, http.StatusBadRequest, "リクエストから調査項目を読み込めません")
		return
	}
	p.ID = bson.NewObjectId()
	if err := c.Insert(p); err != nil {
		respondErr(w, r, http.StatusInternalServerError, "調査項目の格納に失敗しました", err)
		return
	}

	w.Header().Set("Location", "polls/"+p.ID.Hex())
	respond(w, r, http.StatusCreated, nil)
}

func handlepollsDelete(w http.ResponseWriter, r *http.Request) {
	db := GetVars(r, "db").(*mgo.Database)
	c := db.C("polls")
	p := NewPath(r.URL.Path)
	if !p.HasID() {
		respondErr(w, r, http.StatusMethodNotAllowed, "すべての調査項目を削除することはできません")
		return
	}
	if err := c.RemoveId(bson.ObjectIdHex(p.ID)); err != nil {
		respondErr(w, r, http.StatusInternalServerError, "すべての調査項目を削除することはできません", err)
		return
	}

	respond(w, r, http.StatusOK, nil) // 成功
}

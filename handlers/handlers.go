package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"noteservice/db"
	"noteservice/models"
	"time"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var req models.Note
	json.NewDecoder(r.Body).Decode(&req)
	req.Created = time.Now().Local().String()
	db.InsertNote(req)
}

func ReadNotes(w http.ResponseWriter, r *http.Request) {
	notes := db.ReadNotes()

	fmt.Println(">>>>>>>>>", notes)

	json.NewEncoder(w).Encode(notes)
}

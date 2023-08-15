package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Result struct {
	JariJari      string `json:"jariJari"`
	Tinggi        string `json:"tinggi"`
	Volume        string `json:"volume"`
	LuasAlas      string `json:"luasAlas"`
	KelilingAlas  string `json:"kelilingAlas"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	jariJari := 7.0
	tinggi := 10.0

	luasAlas := LuasLingkaran(jariJari)
	kelilingAlas := KelilingLingkaran(jariJari)
	volume := VolumeTabung(jariJari, tinggi)

	result := Result{
		JariJari:      strconv.FormatFloat(jariJari, 'f', -1, 64),
		Tinggi:        strconv.FormatFloat(tinggi, 'f', -1, 64),
		Volume:        strconv.FormatFloat(volume, 'f', -1, 64),
		LuasAlas:      strconv.FormatFloat(luasAlas, 'f', -1, 64),
		KelilingAlas:  strconv.FormatFloat(kelilingAlas, 'f', -1, 64),
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
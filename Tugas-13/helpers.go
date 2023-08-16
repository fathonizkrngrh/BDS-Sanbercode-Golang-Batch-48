package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    any 		`json:"data,omitempty"`
}

func WriteResponse(w http.ResponseWriter, code int, status string, message string, data any) {
	response := jsonResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func GetIndeksNilai(nilai uint) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	}
	return "E"
}

func ParseUint(s string) uint {
	val, _ := strconv.ParseUint(s, 10, 32)
	return uint(val)
}


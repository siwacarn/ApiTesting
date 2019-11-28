package handler

import (
	"api/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllIndicators(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	indicators := []model.Indicator{}
	db.Find(&indicators)
	respondJSON(w, http.StatusOK, indicators)
}

func GetTemperature(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	temp := vars["temperature"]
	respondJSON(w, http.StatusOK, temp)
}
func UpdateTemperature(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	temp := vars["temperature"]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&temp); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, temp)
}

func GetLight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	light := vars["lightintensity"]
	respondJSON(w, http.StatusOK, light)
}

func UpdateLight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	light := vars["lightintensity"]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&light); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, light)
}

func GetHumidity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	humid := vars["humidity"]
	respondJSON(w, http.StatusOK, humid)
}

func UpdateHumidity(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	humid := vars["humidity"]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&humid); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, humid)
}

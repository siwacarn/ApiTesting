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

func CreateIndicators(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondIndicator := model.Indicator{}

	jsonDecoder := json.NewDecoder(r.Body)
	if err := jsonDecoder.Decode(&respondIndicator); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&respondIndicator).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// sucessful return output
	respondJSON(w, http.StatusCreated, respondIndicator)
}

func GetLightValueByDate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Light := vars["light"]
	jsonDecoder := json.NewDecoder(r.Body)
	if err := jsonDecoder.Decode(&Light); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// sucessful return output
	respondJSON(w, http.StatusOK, Light)
}

package handler

import (
	"api/model"
	"encoding/json"
	"log"
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

	Light := vars["Lightintensity"]
	jsonDecoder := json.NewDecoder(r.Body)
	if err := jsonDecoder.Decode(&Light); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// sucessful return output
	respondJSON(w, http.StatusOK, Light)
}

func GetTempValueByDate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Temp := vars["Temperature"]
	log.Println("Temperature by date is:" + Temp)
	jsonDecoder := json.NewDecoder(r.Body)
	if err := jsonDecoder.Decode(&Temp); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// sucessful return output
	respondJSON(w, http.StatusOK, Temp)
}

func GetTempByDateInformation(db *gorm.DB, Temp string, w http.ResponseWriter, r *http.Request) *model.Indicator {
	tempbydate := model.Indicator{}
	err := json.NewDecoder(r.Body).Decode(&tempbydate)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	if err := db.Model(&tempbydate).Find(tempbydate).Error; err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return nil
	}
	return &tempbydate
}

func GetHumidValueByDate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	Humid := vars["Humidity"]
	jsonDecoder := json.NewDecoder(r.Body)
	if err := jsonDecoder.Decode(&Humid); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// sucessful return output
	respondJSON(w, http.StatusOK, Humid)
}

func GetLightValueNow(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	lightnow := model.Indicator{}
	db.Last(&lightnow)
	respondJSON(w, http.StatusOK, lightnow)
}

// func LightValueNowInformation(db *gorm.DB, lightnow int, w http.ResponseWriter, r *http.Request) *model.Indicator {
// 	lightvaluenow := model.Indicator{}
// 	err := json.NewDecoder(r.Body).Decode(&lightvaluenow)
// 	if err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 	}
// 	if err := db.Model(&lightvaluenow).Last(lightvaluenow).Error; err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return nil
// 	}
// 	return &lightvaluenow
// }

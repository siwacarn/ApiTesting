package handler

import (
	"api/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := []model.User{}
	db.Find(&user)
	respondJSON(w, http.StatusOK, user)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]
	user := getUserOr404(db, username, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]
	// update method
	log.Println("username to update: " + username)
	updater := updateUserInformation(db, username, w, r)
	respondJSON(w, http.StatusOK, updater)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]
	// delete method
	log.Println("username to delete:" + username)
	deleter := DeleteUserInfomation(db, username, w, r)
	respondJSON(w, http.StatusNoContent, deleter)
}

func getUserOr404(db *gorm.DB, username string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, model.User{Username: username}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

func updateUserInformation(db *gorm.DB, username string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	if err := db.Model(&user).Where("username = ?", username).Update(user).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

func DeleteUserInfomation(db *gorm.DB, username string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	if err := db.Model(&user).Where("username = ?", username).Delete(user).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

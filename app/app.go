package app

import (
	"api/config"
	"api/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Hostname,
		config.DB.Name,
		config.DB.Charset)

	log.Println(dbURI)

	db, err := gorm.Open(config.DB.Dialect, dbURI)

	if err != nil {
		log.Fatal("Cannot connect to database")
		defer db.Close()
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

//set all required routers
func (a *App) setRouters() {
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

//Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

//Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

//Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

//Wrap the router gor DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
//to be continue

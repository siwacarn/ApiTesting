package app

import (
	"api/config"
	"api/handler"
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
	//user
	a.Get("/alluser", a.GetAllUser)
	a.Post("/alluser", a.CreateUser)
	a.Get("/user/{name}", a.GetUser)
	a.Put("/user/{name}", a.UpdateUser)
	a.Delete("/user/{name}", a.DeleteUser)
	//indicators
	a.Get("/user/{name}/indicator", a.GetAllIndicators)
	a.Get("/user/{name}/indicator/temp", a.GetTemperature)
	a.Put("/user/{name}/indicator/temp", a.UpdateTemperature)
	a.Get("/user/{name}/indicator/light", a.GetLight)
	a.Put("/user/{name}/indicator/light", a.UpdateLight)
	a.Get("/user/{name}/indicator/humid", a.GetHumidity)
	a.Put("/user/{name}/indicator/humid", a.UpdateHumidity)
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

// Handlers to manage User Data
func (a *App) GetAllUser(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUser(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

// Handlers to manage Indicators Data
func (a *App) GetAllIndicators(w http.ResponseWriter, r *http.Request) {
	handler.GetAllIndicators(a.DB, w, r)
}

func (a *App) GetTemperature(w http.ResponseWriter, r *http.Request) {
	handler.GetTemperature(a.DB, w, r)
}

func (a *App) UpdateTemperature(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTemperature(a.DB, w, r)
}

func (a *App) GetLight(w http.ResponseWriter, r *http.Request) {
	handler.GetLight(a.DB, w, r)
}

func (a *App) UpdateLight(w http.ResponseWriter, r *http.Request) {
	handler.UpdateLight(a.DB, w, r)
}

func (a *App) GetHumidity(w http.ResponseWriter, r *http.Request) {
	handler.GetHumidity(a.DB, w, r)
}

func (a *App) UpdateHumidity(w http.ResponseWriter, r *http.Request) {
	handler.UpdateHumidity(a.DB, w, r)
}

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
	a.DB = model.IndicatorDBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

//set all required routers
func (a *App) setRouters() {
	//user
	a.Get("/users", a.GetAllUser)
	a.Post("/user", a.CreateUser)
	a.Get("/user/{username}", a.GetUser)
	a.Put("/user/{username}", a.UpdateUser)
	a.Delete("/user/{username}", a.DeleteUser)
	//indicators
	a.Get("/user/indicator", a.GetAllIndicators)
	a.Post("/user/indicator", a.CreateIndicators)

	// a.Get("/user/indicator/lightnow/latest", a.GetLightValueNow)
	// a.Get("/user/indicator/{parameter}/latest", a.GetTempValueNow)
	// a.Get("/user/indicator/{parameter}/latest", a.GetHumidValueNow)

	a.Get("/user/indicator/light/{date}", a.GetLightValueByDate)
	a.Get("/user/indicator/temp/{date}", a.GetTempValueByDate)
	a.Get("/user/indicator/humid/{date}", a.GetHumidValueByDate)
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

func (a *App) CreateIndicators(w http.ResponseWriter, r *http.Request) {
	handler.CreateIndicators(a.DB, w, r)
}

func (a *App) GetLightValueByDate(w http.ResponseWriter, r *http.Request) {
	handler.GetLightValueByDate(a.DB, w, r)
}

func (a *App) GetTempValueByDate(w http.ResponseWriter, r *http.Request) {
	handler.GetTempValueByDate(a.DB, w, r)
}

func (a *App) GetHumidValueByDate(w http.ResponseWriter, r *http.Request) {
	handler.GetHumidValueByDate(a.DB, w, r)
}

// func (a *App) GetTemperature(w http.ResponseWriter, r *http.Request) {
// 	handler.GetTemperature(a.DB, w, r)
// }

// func (a *App) UpdateTemperature(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateTemperature(a.DB, w, r)
// }

// func (a *App) GetLight(w http.ResponseWriter, r *http.Request) {
// 	handler.GetLight(a.DB, w, r)
// }

// func (a *App) UpdateLight(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateLight(a.DB, w, r)
// }

// func (a *App) GetHumidity(w http.ResponseWriter, r *http.Request) {
// 	handler.GetHumidity(a.DB, w, r)
// }

// func (a *App) UpdateHumidity(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateHumidity(a.DB, w, r)
// }

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

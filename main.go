package FoundeeAPI

import (
	"FoundeeAPI/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func main() {
	log.Print("FoundeeAPI //figlet")

	router := mux.NewRouter()
	RouteUtil(router)

	InitDB()

	http.Handle("/", router)
}

var DBMap *gorm.DB

func InitDB() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := gorm.Open("mysql", "username:password@host/db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Answer{})
	db.AutoMigrate(&models.Ban{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Note{})
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.Vote{})

	DBMap = db
}

package config

// import (
// 	"log"
// 	"os"

// 	"github.com/go-pg/pg/v9"

// 	controllers "github.com/WebDev-Naveen/go-backend/controllers"
// )

// // Connecting to db
// func Connect() *pg.DB {
// 	opts := &pg.Options{
// 		User: "postgres",
// 		Password: "123456",
// 		Addr: "localhost:5432",
// 		Database: "gotest",
// 	}

// 	var db *pg.DB = pg.Connect(opts)
// 	if db == nil {
// 		log.Printf("Failed to connect")
// 		os.Exit(100)
// 	}
// 	log.Printf("Connected to db")
// 	controllers.CreateUserTable(db)
// 	controllers.CreateCoinTable(db)
// 	controllers.InitiateDB(db)
// 	return db
// }

import (
	"log"
	"os"

	"github.com/WebDev-Naveen/go-backend/model"

	"github.com/jinzhu/gorm"
	//dialect for mysql databae
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB -> connection to Database
func DB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error connecting to Database")
		return nil
	}

	db.AutoMigrate(&model.User{}, &model.Coin{})
	return db

}

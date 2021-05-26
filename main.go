package main

import (
	"fmt"

	"github.com/Spores-Labs/spores-nft-backend/app"
	"github.com/Spores-Labs/spores-nft-backend/app/models"
	"github.com/Spores-Labs/spores-nft-backend/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//APIv1 ...
func APIv1(db *gorm.DB) app.IHandlerBuildable {
	hPool := app.HandlerPool{}
	hPool.Push(&app.UserHandler{DB: db})
	return &hPool
}
func main() {
	confOption := conf.GetConfigOption()
	dbConfig := confOption.DBConfig

	// Try to connect to database
	DB_URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName, dbConfig.Password)

	db, err := gorm.Open("postgres", DB_URL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", dbConfig.DBName)
	} else {
		fmt.Printf("Successfully connected to %s database  \n ", dbConfig.DBName)
	}
	defer db.Close()

	//Database migration
	db.AutoMigrate(&models.User{}, &models.Token{})
	hbuilder := APIv1(db)
	s := app.NewServiceServer(confOption.Addr, hbuilder)
	s.Start()

}

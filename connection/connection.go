package connection

import (
	"fmt"
	"log"

	"github.com/prayagsingh/jitsi-event-sync-listener/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	SQL *gorm.DB
}

var dbConn = &DB{}

func SetupDB(db_host, db_user, db_name, db_password, db_port string) *gorm.DB {

	conn_name := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		db_host,
		db_port,
		db_user,
		db_name,
		db_password,
	)

	fmt.Println("Connecting to DB: ", conn_name)
	db, err := gorm.Open(postgres.Open(conn_name), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	db.AutoMigrate(&models.RoomEvents{})
	db.AutoMigrate(&models.RoomDestroyedEvents{}, &models.RoomDestroyedUserDetails{})

	// Initialize the DB instance
	db.Create(&models.RoomEvents{})
	db.Create(&models.RoomDestroyedEvents{})

	dbConn.SQL = db
	return db
}

// GetPostgresDB returns the instance of the DB
func GetPostgresDB() *gorm.DB {
	return dbConn.SQL
}

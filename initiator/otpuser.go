package initiator

import (
	"Auth/internal/glue/routing"
	"Auth/internal/handler/rest"
	"Auth/internal/module/user"
	"Auth/internal/storage/persistence"
	"Auth/platform/redis"
	"Auth/platform/routers/httprouter"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
)

const (
	postgresURL = "host=localhost user=postgres password=dbpass dbname=Usermanagment port=5432 sslmode=disable"

	domain = "user"
)

// User initializes the domain user
func User(testInit bool) {
	err := godotenv.Load("C:\\Users\\Tilefamily\\Desktop\\stygis\\.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")

	redisconnection := "localhost:6379"
	fmt.Printf(redisconnection)
	redispass := os.Getenv("password")
	//redisd := os.Getenv("DB")
	redis := redis.Initialize(redisconnection, redispass, 0)
	//fmt.Println("init", gorm)
	redisDB := redis.Open()
	//gorm.Migrate()
	databaseUser := persistence.UserInit(redisDB)
	//databaseProfile := persistence.ProfileInit(postgresDB)

	//encryptKey := os.Getenv("ACCESS_SECRET")
	//repo := repository.UserInit(encryptKey)

	usecase := user.Initialize(databaseUser)

	handler := rest.UserInit(usecase)
	router := routing.UserRouting(handler)

	//port := os.Getenv("HOST_PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	server := httprouter.Initialize(dbHost, allowedOrigins, domain, router, 8080)

	if testInit {
		logrus.Info("Initialize test mode Finished!")
		os.Exit(0)
	}
	//fmt.Printf("hi")
	server.Serve()
}

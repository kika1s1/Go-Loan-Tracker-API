package main

import (


	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/config"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/routes"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/mongo"
)

func main() {
	// os.Clearenv()
	confs, err := config.Load()
	if err != nil {
		panic(err)
	}
	mongo.ConnectDB(confs.MONGO_URI)
	defer mongo.DisconnectDB()
	mongo.InitializeCollections()

	router := gin.Default()
	// Serve static files from the "upload" directory
	routes.SetUpRoute(router)
	router.Run(":" + confs.GO_PORT)
}

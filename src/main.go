package main

import (
	"time"

	"github.com/dimorinny/queuer/src/auth"
	"github.com/dimorinny/queuer/src/database"
	"github.com/dimorinny/queuer/src/handlers"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

const (
	version = "v1"
	prefix  = "/api/" + version
)

func init() {
	auth.Init()
	database.Init()
	//	database.Migrate()
}

func main() {
	runServer()
	//	createQueue()
}

func createQueue() {
	user1 := database.User{}
	database.Db.First(&user1, 1)

	user2 := database.User{}
	database.Db.First(&user2, 2)

	member1 := database.Member{
		SubscriptionTime: time.Now(),
		User:             user1,
	}

	member2 := database.Member{
		SubscriptionTime: time.Now(),
		User:             user2,
	}

	queue := database.Queue{
		Title:         "Title2",
		Description:   "Description2",
		MaxPeoples:    123,
		Creator:       user1,
		CurrentMember: member1,
		Members:       []database.Member{member1, member2},
		Created:       time.Now(),
		IsActive:      false,
		IsDeleted:     true,
	}

	database.Db.Create(&queue)
}

func runServer() {
	e := echo.New()
	e.Use(mw.Logger())

	api := e.Group(prefix)

	api.Post("/login", auth.Jwt.LoginHandler())
	api.Post("/refresh", auth.Jwt.RefreshTokenHandler())
	api.Post("/register", auth.Register)

	authRequired := api.Group("")
	authRequired.Use(auth.Jwt.AuthRequired())

	authRequired.Get("/queue", handlers.Queues)
	authRequired.Post("/queue", handlers.CreateQueue)
	authRequired.Get("/queue/my", handlers.MyQueues)
	authRequired.Get("/queue/:id", handlers.Queue)
	authRequired.Delete("/queue/:id", handlers.DeleteQueue)
	authRequired.Put("/queue/:id/members", handlers.JoinQueue)
	authRequired.Delete("/queue/:id/members", handlers.LeaveQueue)

	adminRequired := authRequired.Group("")
	adminRequired.Use(auth.AdminRequired)

	adminRequired.Post("/queue/:id/next", handlers.NextMember)
	adminRequired.Delete("/queue/:queueID/member/:memberID", handlers.DeleteMember)
	adminRequired.Post("/queue/:id/active", handlers.ActiveQueue)

	// Start server
	e.Run(":8080")
}

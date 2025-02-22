package main

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite" 
    "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"yourquote/internal/models"
	"yourquote/internal/controller"
	"log"
	"time"
	"os"
	"github.com/patrickmn/go-cache"
	"github.com/gin-contrib/cors"

)
func setupRouter() *gin.Engine{
    router := gin.Default()
	return router
}

func main (){
	err := godotenv.Load()
	if err != nil{
		panic("faild to load env var")
	}
	port := os.Getenv("WEBPORT")
	log.Println(os.Getenv("WEBPORT"))

	databasepath := os.Getenv("DATABASE")
    db,err := sql.Open("sqlite",databasepath)
    if err !=nil{
		panic(err)
    }

	r := gin.Default()
	
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour
    r.Use(cors.New(config))
	
	r.LoadHTMLGlob("internal/views/**/*")
	r.Static("/assets", "./internal/assets")
    c := cache.New(5*time.Second, 10*time.Second)
   


	quoteshandler := controller.QuoteHandler{
		Store : models.QuoteStore{
			Quotes : models.QuoteModel{
				Conn : db,
			},
            Getquote :models.QuoteModel{
                Conn : db,
            }, 
            GetSingle :models.QuoteModel{
                Conn : db,
            }, 
            Exect :models.QuoteModel{
                Conn : db,
            }, 
		},
		Page : models.PaginationStore{
			GetSingle : 	db,
		},
		Cache : c,
	
	}
	r.GET("/", controller.Basepage)
	r.GET("/quote", quoteshandler.GetAllQuote)
	//r.GET("/:id", quoteshandler.GetQuote)
	r.POST("/quote", quoteshandler.CreateQuote)
	//r.PUT("/:id", quoteshandler.UpdateQuote)
	//r.DELETE("/:id", quoteshandler.DeleteQuote)
	
    r.Run(port)
}

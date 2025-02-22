package controller

import (
    "database/sql"
	_ "github.com/glebarez/go-sqlite" 
    "fmt"
    "time"
    "yourquote/internal/models"
    "strings"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/gin-gonic/gin"

    "github.com/patrickmn/go-cache"

)

var hellotestmessage = `{"message":"hello"}`
var succsesmessage = `{"message":"Succsess"}`
var deletemessage =    `{"message":"quote delete"}`
var idnotfoundmessage = `{"message":"id Not found"}`
var wrongformatmessage = `{"message":"json: cannot unmarshal number into Go struct field Quote.message of type string"}`
var wrongformatupdatemessage = `{"message":"json: cannot unmarshal number into Go struct field Quote.message of type string"}{"id":0,"message":"","author":""}`


var fakequotes = []models.Quote{
    {
        Id:"0",
        Message:"Test",
        Author:"Test",
    },
}


var createfakequote = QuoteDTO{
    Message:"Test",
    Author:"Test",
}

var fakedata = Data{
    Quotes : fakequotes,
    Pages : &models.Pagination{
        TotalRecord : 1,
        CurrentPage : 1,
        TotalPage : 1,
        NextPage : 1,
        PrevPage : 0,
    },

}

type QuoteHandlerPositiveTestSuite struct {
	suite.Suite
	testquoteshandler QuoteHandler
	db *sql.DB
	router *gin.Engine
    cache *cache.Cache
}

func (suite *QuoteHandlerPositiveTestSuite) SetupTest(){
    database,dberr := sql.Open("sqlite", ":memory:")
    if dberr != nil{
        fmt.Println("Connection Db Error")
    }
    querycreatetesttable := `CREATE TABLE IF NOT EXISTS QUOTES (
        Id TEXT PRIMARY KEY,
        Message     TEXT NOT NULL,
        Author TEXT,
        Created_At TEXT
        );`
    _,err := database.Exec(querycreatetesttable)
    if err != nil{
        fmt.Println("Created Tabel Error")
    }
	fmt.Println(">>> From SetupSuite")
	suite.db = database
    c := cache.New(5*time.Second, 10*time.Second)
    
	suite.testquoteshandler = QuoteHandler{
       store : models.QuoteStore{
				Quotes : models.QuoteModel{
					Conn : suite.db,
				},
                Getquote :models.QuoteModel{
                    Conn : suite.db,
                }, 
                GetSingle :models.QuoteModel{
                    Conn : suite.db,
                }, 
                Exect :models.QuoteModel{
                    Conn : suite.db,
                }, 
			},
		page : models.PaginationStore{
            GetSingle : suite.db,
		},
		cache : c,
    }
    suite.router = setupRouter()
    
}

func (suite *QuoteHandlerPositiveTestSuite) TestCreateQuoteSucsses(){
    suite.router.POST("/v1/quotes",suite.testquoteshandler.CreateQuote)
    w := httptest.NewRecorder()
    
    thridQuoteJSON , _ := json.Marshal(createfakequote)
    req := httptest.NewRequest("POST","/v1/quotes",strings.NewReader(string(thridQuoteJSON)))
    req.Header.Add("idempotency-key", "cd27e182-a82a-4f2b-b76d-3ce2707bead8")
    suite.router.ServeHTTP(w, req)

    suite.Equal(201, w.Code)
	suite.Equal(succsesmessage, w.Body.String())
	
}

func (suite *QuoteHandlerPositiveTestSuite)TestGetAllQuoteSucsses(){
   
    suite.router.POST("/v1/quotes",suite.testquoteshandler.CreateQuote)
    thridQuoteJSON , _ := json.Marshal(createfakequote)
    req := httptest.NewRequest("POST","/v1/quotes",strings.NewReader(string(thridQuoteJSON)))
    w := httptest.NewRecorder()
    req.Header.Add("idempotency-key", "cd27e182-a82a-4f2b-b76d-3ce2707bead8")
    suite.router.ServeHTTP(w, req)
	 
    suite.router.GET("/v1/quotes/",suite.testquoteshandler.GetAllQuote)
    req = httptest.NewRequest("GET","/v1/quotes/?page=1",nil)
    
    w = httptest.NewRecorder()
    suite.router.ServeHTTP(w, req)
    suite.Equal( 200, w.Code)
    
    currentTime := time.Now()
        
        formatTime := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
            currentTime.Year(),
            currentTime.Month(),
            currentTime.Day(),
            currentTime.Hour(),
            currentTime.Hour(),
            currentTime.Second())
            
    
    
    
    var dat Data
    json.Unmarshal(w.Body.Bytes(),&dat)
    fmt.Println(dat.Quotes[0].Id)
    
    
    var data Data = Data{
        Quotes :[]models.Quote {
            {
                Id:dat.Quotes[0].Id,
                Message:"Test",
                Author:"Test",
                Created_At : formatTime,
            },
        },
        Pages : &models.Pagination {
            TotalRecord : 1,
            CurrentPage : 1,
            TotalPage : 1,
            NextPage : 1,
            PrevPage : 0,
        },
    }
    d,_ := json.Marshal(data)
	suite.Equal(string(d), w.Body.String())
	
}


func TestQuoteHandlerPositiveTestSuite(t *testing.T) {
	suite.Run(t, new(QuoteHandlerPositiveTestSuite))
}




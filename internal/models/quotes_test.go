package models



import (
    "database/sql"
	_ "github.com/glebarez/go-sqlite"
	"github.com/stretchr/testify/suite"

	"fmt"
	"time"
	"testing"
)


type ModelsQuotesTestSuite struct{
	suite.Suite
	quotes QuoteStore
	con *sql.DB
}
var currentTime time.Time = time.Now()

var timenow string = fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())


func (suite *ModelsQuotesTestSuite) SetupTest(){
	db,dberr := sql.Open("sqlite", ":memory:")
    if dberr != nil{
        fmt.Println("Connection Db Error")
    }
    querycreatetesttable := `CREATE TABLE IF NOT EXISTS QUOTES (
        Id TEXT PRIMARY KEY,
        Message     TEXT NOT NULL,
        Author TEXT,
        Created_At TEXT
        );`
    suite.con = db
    _,err := suite.con.Exec(querycreatetesttable)
    if err != nil{
		fmt.Println("Create Table Error")
    }
    suite.quotes = QuoteStore{
		Quotes : QuoteModel{
			Conn : suite.con,
		},
		Getquote :QuoteModel{
			Conn : suite.con,
		}, 
		GetSingle :QuoteModel{
			Conn : suite.con,
		}, 
		Exect :QuoteModel{
			Conn : suite.con,
		}, 
		
    }
}
func (suite *ModelsQuotesTestSuite) TestCreateSuccess(){
	
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
}

func (suite *ModelsQuotesTestSuite) TestAllSuccess(){
		

	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	allquote := suite.quotes.All(1)
	suite.Equal([]Quote{
		Quote{
			Id : "51a9c21f-889a-45af-8373-e7d3736374f4",
			Message : "Test",
			Author : "Test",
			Created_At :timenow,
		},
	},allquote)
	
}
func (suite *ModelsQuotesTestSuite) TestGetByIdSuccess(){

	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	fristquote,err := suite.quotes.Get("51a9c21f-889a-45af-8373-e7d3736374f4")
	suite.Equal(nil,err)
	suite.Equal(
		Quote{
			Id : "51a9c21f-889a-45af-8373-e7d3736374f4",
			Message : "Test",
			Author : "Test",
			Created_At :timenow,
		},fristquote)
}
func (suite *ModelsQuotesTestSuite) TestUpdateSuccess(){
	
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	var uq  = Quote{
		Id : "51a9c21f-889a-45af-8373-e7d3736374f4",
		Message : "updatedTest",
		Author : "updatedTest",
	}
	updatedquote := suite.quotes.Update("51a9c21f-889a-45af-8373-e7d3736374f4",uq)
	suite.Equal(nil,updatedquote)
}
func (suite *ModelsQuotesTestSuite) TestDeleteSuccess(){
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	deletedquote := suite.quotes.Delete("51a9c21f-889a-45af-8373-e7d3736374f4")
	suite.Equal(nil,deletedquote)
}
func (suite *ModelsQuotesTestSuite) TestCheckisExisSuccess(){
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	checkid := suite.quotes.CheckisExis("51a9c21f-889a-45af-8373-e7d3736374f4")
	suite.Equal(nil,checkid)
}

func (suite *ModelsQuotesTestSuite) TestGetByIdFail(){
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	fristquote,err := suite.quotes.Get("100")
	suite.Equal(Quote{Id:"", Message:"", Author:""},fristquote)
	suite.NotEqual(nil,err)
}
func (suite *ModelsQuotesTestSuite) TestUpdatedFail(){
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	var uq  = Quote{
		Id : "51a9c21f-889a-45af-8373-e7d3736374f4",
		Message : "updatedTest",
		Author : "updatedTest",
	}
	updatedquote := suite.quotes.Update("9999",uq)
	suite.NotEqual(nil,updatedquote)
}
func (suite *ModelsQuotesTestSuite) TestDeleteFail(){
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	deletedquote := suite.quotes.Delete("999")
	suite.NotEqual(nil,deletedquote)
}
func (suite *ModelsQuotesTestSuite) TestCheckisExisFail(){
	Id := "51a9c21f-889a-45af-8373-e7d3736374f4"
	Message := "Test"
	Author := "Test"
	
	err := suite.quotes.Create(Id,Message,Author)
	suite.Equal(nil,err)
	checkid := suite.quotes.CheckisExis("000")
	suite.NotEqual(nil,checkid)
}

func (suite *ModelsQuotesTestSuite) TearDownTest(){
	suite.con.Close()
}

func TestModelsTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsQuotesTestSuite))
}

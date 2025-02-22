package models

import (
    "database/sql"
	_ "github.com/glebarez/go-sqlite"
	"github.com/stretchr/testify/suite"
	"log"
	"fmt"
	"time"
	"testing"
)

type ModelsPaginationTestSuite struct{
	suite.Suite
	pagination PaginationStore
	con *sql.DB
}


func (suite *ModelsPaginationTestSuite)SetupTest(){
	db,dberr := sql.Open("sqlite", ":memory:")
    if dberr != nil{
        fmt.Println("Connection Db Error")
    }
    querycreatetesttable := `CREATE TABLE IF NOT EXISTS QUOTES (
        Id INTEGER ,
        Message     TEXT NOT NULL,
        Author TEXT,
        Created_At TEXT
        );`
    suite.con = db
    _,err := suite.con.Exec(querycreatetesttable)
    if err != nil{
		fmt.Println("Create Table Error")
    }
    suite.pagination = PaginationStore{
		GetSingle :QuoteModel{
			Conn : suite.con,
		}, 
    }
	
	for  i := range 100 {
	

		message := "FORR TESTINGG !!!!!"
		author := "TESTINGGGG!!!"

	
		currentTime := time.Now()
		formatTime := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			currentTime.Hour(),
			currentTime.Hour(),
			currentTime.Second())
		_,err := suite.con.Exec("INSERT INTO QUOTES (Id,Message,Author,Created_At) VALUES(?,?,?,?)",i,message,author,formatTime)
		if err != nil{
			log.Println(err)
		}

	}
}

func (suite *ModelsPaginationTestSuite)TestGetCountAllRecord(){
	totalrecord := suite.pagination.GetCountAllRecord()
	suite.Equal(100,totalrecord)
}
func (suite *ModelsPaginationTestSuite)TestGetPagination() {
	
	for i := 1; i <= 10; i++ {
		testpagination := suite.pagination.GetPagination(i)
		var prevpage int = 0
		if i == 1 {
			prevpage =  0 
		}else{
			prevpage = i -1
		}
		page := &Pagination{
			TotalRecord : 100,
			CurrentPage : i,
			TotalPage : 10,
			NextPage : i+1,
			PrevPage : prevpage,
			
		}
		suite.Equal(page,testpagination)
	}
	
}
func (suite *ModelsPaginationTestSuite)TestGetPaginationFail() {
		testpagination := suite.pagination.GetPagination(1000)
		var page *Pagination = nil
		suite.Equal(page,testpagination)
	
}
func (suite *ModelsPaginationTestSuite) TearDownTest(){
	suite.con.Close()
}

func TestPaginationTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsPaginationTestSuite))
}


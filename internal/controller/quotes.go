package controller


import (
    "log"
    "strconv"
    
    "yourquote/internal/models"
    "yourquote/internal/dto"
    
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/patrickmn/go-cache"

    "github.com/google/uuid"
)


type QuoteHandler struct{
    Store models.QuoteStoreer
    Page models.PaginationStore
    Cache *cache.Cache
}

/*
     c.JSON(200, gin.H{
		"message": "hello",
	})
*/

func Basepage(c *gin.Context){
    c.HTML(http.StatusOK,"quotes/index.html",gin.H{
    })
}
func(q QuoteHandler) GetAllQuote(c *gin.Context){
    
    pagenumber := c.DefaultQuery("page", "1")
    number,err := strconv.Atoi(pagenumber)
    
    
    if err != nil{
        log.Println(err)
    }
    pages := q.Page.GetPagination(number)
    
    skippage := (pages.CurrentPage - 1) * 10
    
    leftpage := 0
    if (skippage <= 0){
        leftpage = 10
    }else{
    
        leftpage = skippage - pages.TotalRecord
    }
    log.Println(skippage)
    log.Println(leftpage)
    
    quotes := q.Store.All(skippage,leftpage)
    listdto := dto.QuotetoListDTO(quotes)
    //data := NewData(quotes,pages)
    c.HTML(http.StatusOK,"quotes/quotes.html",gin.H{
        "listdto":listdto,
        "pages":pages,
        },
    )
}
func(q QuoteHandler) CreateQuote(c *gin.Context){
    var newquotes dto.QuoteDTO
    
    if  err := c.Bind(&newquotes);err != nil{
        //log.Println(err)
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "something error",
        })
        return 
    }
   
    id  := uuid.New()
    
    if (newquotes.Author == ""){
        newquotes.Author = "anonymous"
    }
    if err := q.Store.Create(id.String(),newquotes.Message,newquotes.Author) ; err != nil{
        log.Println(err)
    }
    log.Println(id.String(),newquotes.Message,newquotes.Author)
    c.Redirect(http.StatusFound , "/")
    
}


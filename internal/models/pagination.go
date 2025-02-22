package models


import (
    "math"
    "log"
)
type Pagination struct{
    TotalRecord int `json"total_record"`
    CurrentPage int `json"curent_page"`
    TotalPage int `json"total_page"`
    NextPage int `json"next_page"`
    PrevPage int `json"prev_page"`
}


type PaginationStore struct{
    GetSingle GetSingleOperator
}
func (p PaginationStore)GetCountAllRecord()int{
    var countrecord int
    err := p.GetSingle.QueryRow("SELECT COUNT(Id) FROM QUOTES").Scan(&countrecord)
    if err !=nil{
        log.Println(err)
    }
    return countrecord    
}

// TODO need error when page number exeed total page
func (p PaginationStore) GetPagination(pagenumber int) Pagination{
    totalrecord := p.GetCountAllRecord()
    
    totalpage :=  (int) (math.Ceil((float64)(totalrecord) / 10.0))
    //kontol := math.Ceil((float64)totalrecord / 10.0)
    
    
    if totalrecord <= 10{
      totalpage = 1  
    }
    if pagenumber >= totalpage{
        pagenumber = totalpage
    }
    nextpage := pagenumber + 1 
    if nextpage >= totalpage {
        nextpage = totalpage
    }
    
    prevpage := pagenumber - 1
    if  pagenumber <= 1 {
		prevpage =  1 
	}
    
    pages := Pagination{
        TotalRecord : totalrecord,
        CurrentPage : pagenumber,
        TotalPage : totalpage,
        NextPage : nextpage,
        PrevPage : prevpage,
    }
    return pages
}

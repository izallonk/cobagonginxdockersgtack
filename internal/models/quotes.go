package models

import (
    "time"
    //"fmt"
    "log"
    //"strconv"
    _ "github.com/glebarez/go-sqlite"
    
)

type Quote struct{
    Id string 
    Message string 
    Author string 
    Created_At int 
}


type QuoteStoreer interface{
    All(int,int) []Quote
    Get(string) (Quote,error)
    Create(string,string,string)error
    Update(string,Quote) error
    Delete(string) error
    CheckisExis(string) error
}

type QuoteStore struct{
    Quotes Operatorer
    Getquote GetOperator
    GetSingle GetSingleOperator
    Exect ExectOperator
}

// need limit 

/*
func (qs QuoteStore) All() []Quote{
    return qs.quotes.Query("SELECT Id,Message,Author FROM QUOTES ORDER BY Created_At LIMIT 10")
}*/

func (qs QuoteStore) All(skippage, leftpage int) []Quote{ 
    return qs.Getquote.Query("SELECT Id,Message,Author,Created_At FROM QUOTES WHERE Id NOT IN (SELECT Id FROM QUOTES ORDER BY Created_At DESC  LIMIT ?) ORDER BY Created_At DESC  LIMIT ?",skippage,leftpage)
}

func (qs QuoteStore) Get(id string ) (Quote,error){
    var q Quote
    err := qs.GetSingle.QueryRow("SELECT Id,Message,Author,Created_At FROM QUOTES WHERE Id = ?",id).Scan(&q.Id,&q.Message,&q.Author,&q.Created_At)
    if err != nil{
        log.Println(err)
        return Quote{},err
    }
    return q,nil
}
func (qs QuoteStore) Create(id,m,a string) error {
    currentTime := time.Now()
	formatTimeINT := currentTime.Unix()
	
    err := qs.Exect.Execute("INSERT INTO QUOTES (Id,Message,Author,Created_At) VALUES(?,?,?,?)",id,m,a,formatTimeINT)
    if err != nil{
        log.Println("KONTOl")
        log.Println(err)
        return nil
    }
    return err
}
func (qs QuoteStore) Update(id string, q Quote) error {
    if err :=qs.CheckisExis(id);err != nil{
        return err
    }
    err := qs.Exect.Execute("UPDATE QUOTES SET Message=?,Author=? WHERE Id = ?",q.Message,q.Author,q.Id)
    return err
}
func (qs QuoteStore) Delete(id string) error{
    if err :=qs.CheckisExis(id);err != nil{
        return err
    }
    return qs.Exect.Execute("DELETE FROM QUOTES WHERE Id = ?",id)
    
}
func (qs QuoteStore) CheckisExis(id string) error{
   return qs.Quotes.CheckisExis(id)
}

func DtoToQuote(m ,a string)Quote{
    return Quote{
        Message: m,
        Author : a,
    }
}



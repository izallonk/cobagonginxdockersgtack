package models

import (
    "errors"
    "log"
    "database/sql"
    
    _ "github.com/glebarez/go-sqlite" 
)

type GetOperator interface{
    Query(string, ...any) []Quote
}
type GetSingleOperator interface{
    QueryRow(string, ...any) *sql.Row
}
type ExectOperator interface{
    Execute(string, ...any)error
}


type Operatorer interface{    
    CheckisExis(string) error
}


type QuoteModel struct{
    Conn *sql.DB
}

func (d QuoteModel)Query(query string, args ...any) []Quote{
    rows,err := d.Conn.Query(query,args...)
    if err != nil{
        log.Println(err)
    }
    var quotes []Quote
    for rows.Next(){
        var quote Quote
        err := rows.Scan(&quote.Id,&quote.Message,&quote.Author,&quote.Created_At)
        if err != nil{
            log.Println(err)
        }
        quotes = append(quotes,quote)    
    }
    return quotes

}
func (d QuoteModel)QueryRow(query string, args ...any) *sql.Row{
    err := d.Conn.QueryRow(query,args...)
	return err
}
func (d QuoteModel)Execute(query string, args ...any)error{
    tx,err := d.Conn.Begin()
    if err != nil{
        log.Println(err)
        return err
    }
    _,exeerr := tx.Exec(query,args...)
    if exeerr != nil{
        tx.Rollback()
        log.Println(exeerr)
        return  exeerr
    }
    if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}
    return nil
}
func (d QuoteModel)CheckisExis(id string)error{
    err := d.Conn.QueryRow("SELECT Id FROM QUOTES WHERE Id = ?",id).Scan(&id)
    switch {
        case err != nil:
            return errors.New("Id Not Found")
        default:
            log.Println("check Id : &d",id)
    }
    return nil
}

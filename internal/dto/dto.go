package dto


import (
    "yourquote/internal/models"
    //"yourquote/utils"
)

type Data struct{
    Quotes []models.Quote `json"data"`
    Pages models.Pagination `json"pagenation"`
}

func NewData(q []models.Quote,p models.Pagination) Data{
    return Data{
        Quotes : q,
        Pages : p,
    }
}

type QuoteDTO struct{
    Message string `form:"message" binding:"required"`
    Author string `form:"author"`
}


type QuoteDTOdate struct{
    Message string `form:"message"`
    Author string   `form:"author"`

}


func QuotetoDTO(quote models.Quote) QuoteDTO{
    return QuoteDTO{
        Message : quote.Message,
        Author : quote.Author,
    
    }
}



func QuotetoListDTO(quote []models.Quote) []QuoteDTOdate{
    var listDTO []QuoteDTOdate
    for i := range(len(quote)){
       
        var dto QuoteDTOdate
        dto.Message = quote[i].Message
        dto.Author = quote[i].Author
        listDTO = append(listDTO,dto)
    }
    return listDTO
}

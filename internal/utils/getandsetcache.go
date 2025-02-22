package utils

import (
    "fmt"
    "log"
    "github.com/patrickmn/go-cache"
	"time"
	"strings"
	"strconv"
)

func Setidomkey(c *cache.Cache,key string){
    c.Set(key,1, 2 * time.Second)
    log.Println(fmt.Printf("key : %s is save\n" ,key))
} 


func Getidomkey(c *cache.Cache,key string) error{
    _,d := c.Get(key)
    log.Println(fmt.Printf("Getting key : %s  \n" ,key))
    
    if !d {
        return fmt.Errorf("Data Expire or Not Found")
    }
    log.Println(fmt.Printf("Key Is exsis %v \n",d))
    return nil
}


func DateintToFormatString(time int)string{
    b := strconv.Itoa(time)
	res := strings.Split(b, "")

	tahun := strings.Join(res[0:4], "")
	bulan := strings.Join(res[4:5], "")
	tanggal := strings.Join(res[5:7], "")
	jam := strings.Join(res[7:9], "")
	menit := strings.Join(res[9:11], "")
	detik := strings.Join(res[11:13], "")

	formatTime := fmt.Sprintf("%s-%s-%s %s:%s:%s",
		tahun,
		bulan,
		tanggal,
		jam,
		menit,
		detik,
	)

	return formatTime

}

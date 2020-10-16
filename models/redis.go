package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
	})

func CountSent(number int) {
	val,err := rdb.Get(ctx,strconv.Itoa(number)).Result()
	i,_ := strconv.Atoi(val)
	if err != nil {
		rdb.Set(ctx,strconv.Itoa(number),1,0).Err()
	}
	err = rdb.Set(ctx,strconv.Itoa(number),i + 1,0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func TopNum(x int) (count []string,err error){
	//m := []interface{}{}
	val,_ := rdb.Keys(ctx,"*").Result()
	var n = len(val)
	if n > 10 {
		n = 10
	}
	var item1int,item2int int
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			item1,_ := rdb.Get(ctx,val[j-1]).Result()
			item1int,_ = strconv.Atoi(item1)
			item2,_ := rdb.Get(ctx,val[j]).Result()
			item2int,_ = strconv.Atoi(item2)
			if item1int > item2int {
				val[j-1], val[j] = val[j], val[j-1]
			}
			j = j - 1
		}
	}

	return val,nil
}



package main 


import (
	"fmt"
	"context"
	redis "github.com/go-redis/redis/v8"
)



var ctx = context.Background()

func RedisSetter(redcl redis.Client,key string, value interface{})error{

	err := redcl.Set(ctx, key , value, 0).Err()
	if err!=nil{
		fmt.Println("Error on setting chache ", err)
		return err 
	}
	return nil
	

}



func RedisGetter(redcl redis.Client, key string,)(interface{}, error){
	val, err := redcl.Get(ctx,key).Result()
	if err!=nil{
		if err==redis.Nil{
			fmt.Println("this key does not excists")
			return nil,err
		}
		return nil,err
	}

	return val, nil
}


func main(){
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	

	pong ,err := client.Ping(context.Background()).Result()
	if err!=nil{
		fmt.Println("There are some error like ",err)
		return 
	}

	fmt.Println("Redis server worked all is ok ",pong)

	serr :=RedisSetter(*client ,"name","Kesha") 

	if serr!=nil{
		panic(serr)
		return 
	}

	val,ver:= RedisGetter(*client, "name");
	if ver!=nil{
		fmt.Println(ver)
		return 
	}
	fmt.Println("the value is ", val)





}
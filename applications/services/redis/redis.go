package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	eh "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/redis/errorhandler"
	pb "github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/redis/rpc"

	"github.com/go-redis/redis"
)

var (
	redis_key = "attendance_code"
)

func GetRedisClient(config Configuration) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{config.Redis.Broker},
		Password: "",
	})
}

func RandomCode() int64 {
	low := 1000000
	hi := 9999999
	rand.Seed(time.Now().UnixNano())
	return int64(low + rand.Intn(hi-low))
}

func GetUniqueCode(config Configuration) int64 {
	rdb := GetRedisClient(config)
	code := RandomCode()
	for {
		exists := rdb.HExists(redis_key, strconv.FormatInt(code, 10)).Val()
		if !exists {
			break
		}
		code = RandomCode()
	}
	return code
}

func CreateAttendanceCodeInRedis(in *pb.AttendanceCodeCreate, config Configuration) (int64, int64, float64, float64, error) {
	rdb := GetRedisClient(config)
	code := GetUniqueCode(config)
	unix := time.Now().UnixNano()/1000000 + (in.MinutesToLive * 60 * 1000)
	dataAsJson := fmt.Sprintf(`{"unix": %d, "lat": %f, "long": %f}`, unix, in.Lat, in.Long)

	result := rdb.HSet(redis_key, strconv.FormatInt(code, 10), dataAsJson).Val()
	if !result {
		log.Printf("error when adding code to redis")
		return 0, 0, 0, 0, errors.New("error when adding code to redis")
	}

	return code, unix, in.Lat, in.Long, nil
}

type jsonData struct {
	Unix int64   `json:"unix,omitempty"`
	Lat  float64 `json:"lat,omitempty"`
	Long float64 `json:"long,omitempty"`
}

func GetAttendanceCodeFromRedis(code int64, config Configuration) (int64, int64, float64, float64, error) {
	rdb := GetRedisClient(config)
	exists := rdb.HExists(redis_key, strconv.FormatInt(code, 10)).Val()
	if !exists {
		log.Printf("code not found in redis")
		return 0, 0, 0, 0, errors.New("code not found in redis")
	}

	result := rdb.HGet(redis_key, strconv.FormatInt(code, 10)).Val()

	s, _ := strconv.Unquote(string(result))
	var data jsonData
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		eh.PanicOnError(err, "Couldn't convert json result to JSON data")
	}

	// unix, err := strconv.ParseInt(result, 10, 64)
	// eh.PanicOnError(err, "Error converting unix to int64")

	return code, data.Unix, data.Lat, data.Long, nil
}

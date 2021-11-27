package main

import (
	"net/http"
	"strconv"

	pb "api_service/attendancecode"
	eh "api_service/errorhandler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// getAlbums responds with the list of all albums as JSON.
func CreateAttendanceCode(c *gin.Context) {
	minutesToLiveStr := c.Param("minutesToLive")
	minutesToLive, err := strconv.ParseInt(minutesToLiveStr, 10, 64)
	eh.PanicOnError(err, "Minutes to live is not an int")

	conn, err := grpc.Dial("redis-service.redis:50051", grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")
	defer conn.Close()

	client := pb.NewAttendanceCodeProtoClient(conn)

	newAttendanceCode := &pb.AttendanceCodeCreate{MinutesToLive: minutesToLive}
	attendancecode, err := client.CreateAttendanceCode(c, newAttendanceCode)
	eh.PanicOnError(err, "Failed to create attendance code")

	c.IndentedJSON(http.StatusOK, attendancecode)
}

func GetAttendanceCode(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.ParseInt(codeStr, 10, 64)
	eh.PanicOnError(err, "Minutes to live is not an int")

	conn, err := grpc.Dial("redis-service.redis:50051", grpc.WithInsecure())
	eh.PanicOnError(err, "fail to dial")
	defer conn.Close()

	client := pb.NewAttendanceCodeProtoClient(conn)

	attendancecode, err := client.GetAttendanceCode(c, &wrapperspb.Int64Value{Value: code})
	eh.PanicOnError(err, "Failed to create attendance code")

	c.IndentedJSON(http.StatusOK, attendancecode)
}

func main() {
	router := gin.Default()
	router.GET("/api/attendance_code/:code", GetAttendanceCode)
	router.POST("/api/attendance_code/:minutesToLive", CreateAttendanceCode)

	router.Run("0.0.0.0:8080")
}

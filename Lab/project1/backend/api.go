package main

import (
  //  "encoding/json"
    "context"
    "net/http"
    "fmt"
  //  "bytes"
  //  "log"
    "os"
   // "io"

    "github.com/gin-gonic/gin"
)

type courserecord struct {
  Curso    string `json:"curso"`
  Facultad string `json:"facultad"`
  Carrera  string `json:"carrera"`
  Region   string `json:"region"`
}

var (
    grpcClientUrl = fmt.Sprintf("%s:%s", os.Getenv("GRPC_CLIENT_HOST"), os.Getenv("GRPC_CLIENT_PORT"))
//    grpcServerUrl = fmt.Sprintf("%s:%s", os.Getenv("GRPC_SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT"))
    ctx = context.Background()
//    rustServerUrl = fmt.Sprintf("http://%s:%s", os.Getenv("RUST_SERVER_HOST"), os.Getenv("RUST_SERVER_PORT"))
)

func allGood(c *gin.Context) {
    c.String(http.StatusOK, "Course REST API Server Ready")
}

func postCourse(c *gin.Context) {
    var courseRecord courserecord

    
    if courseDataError := c.BindJSON(&courseRecord); courseDataError != nil {
        fmt.Println(courseDataError)
        c.String(http.StatusBadRequest, courseDataError.Error())
        return
    }

    
    c.String(http.StatusOK, "ok")

}

func main() {
    router := gin.Default()
    router.GET("/", allGood)
    router.POST("/course", postCourse)

    router.Run(grpcClientUrl)
}

// export GRPC_CLIENT_PORT=8000 \
// export GRPC_CLIENT_HOST=localhost

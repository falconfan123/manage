package main

import (
	srv "falconfan123/manage/internal/common/run"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	srv.Run(r, "webcenter", ":80")
}

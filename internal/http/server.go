package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port    string
	Handler *Handler
}

func (s *Server) Run() {
	r := gin.Default()
	r.GET("/excel-type-1", s.Handler.getExcelType1)
	r.GET("/excel-type-2", s.Handler.getExcelType2)
	r.GET("/excel-type-3", s.Handler.getExcelType3)
	err := r.Run(fmt.Sprintf(":%v", s.Port))
	if err != nil {
		return
	}
}

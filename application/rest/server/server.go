package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {

	r := gin.Default()

	r.Use(otelgin.Middleware("my-server"))

	r.GET("/", s.GetUser)

	_ = r.Run(":3000")
}

func (s *Server) GetUser(c *gin.Context) {

	var id string
	id = "123"
	var tracer = otel.Tracer("gin-server")
	_, span := tracer.Start(c.Request.Context(), "getUser",
		oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	if id == "123" {
		c.JSON(http.StatusOK, gin.H{
			"message": id,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unknown"})
	}
}

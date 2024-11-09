package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tittuvarghese/core/logger"
	"github.com/tittuvarghese/customer-service/proto"
	"github.com/tittuvarghese/gateway/constants"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"time"
)

type Server struct {
	Router          *gin.Engine
	CustomerService proto.AuthServiceClient
}

var log = logger.NewLogger("httpserver")

func NewHttpSerer() *Server {
	return &Server{Router: gin.New()}
}

func (s *Server) EnableLogger() {
	s.Router.Use(gin.Logger())
}

func (s *Server) EnableRecovery() {
	s.Router.Use(gin.Recovery())
}

func (s *Server) EnableRateLimiter() {
	s.Router.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 10), time.Hour // limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 1 hour
	}, func(c *gin.Context) {
		c.AbortWithStatus(429) // handle exceed rate limit request
	}))
}

func (s *Server) AddHandler(method, service, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		s.Router.GET(constants.ApiBasePath+service+path, handler)
	case "POST":
		s.Router.POST(constants.ApiBasePath+service+path, handler)
	}
}

func (s *Server) Run(port string) {
	if err := s.Router.Run(":" + port); err != nil {
		log.Error("Unable to start the server", err)
	}
}

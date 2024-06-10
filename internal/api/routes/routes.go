package routes

import (
	"github.com/sanniv2002/nether/internal/api/handlers"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RegisterRoutes(r *gin.Engine, grpcConn *grpc.ClientConn) {
	containerHandler := handlers.NewContainerHandler(grpcConn)

	r.POST("/spawn", containerHandler.StartContainer)
	r.DELETE("/kill", containerHandler.KillContainer)
}

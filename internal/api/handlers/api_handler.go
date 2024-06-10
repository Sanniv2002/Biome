package handlers

import (
	"context"
	"net/http"

	"github.com/sanniv2002/nether/proto/containerinit"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ContainerHandler struct {
	grpcClient containerinit.ContainerInitClient
}

func NewContainerHandler(grpcConn *grpc.ClientConn) *ContainerHandler {
	return &ContainerHandler{
		grpcClient: containerinit.NewContainerInitClient(grpcConn),
	}
}

func (h *ContainerHandler) StartConfig(c *gin.Context) {
	var req containerinit.StartContainerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.grpcClient.StartContainer(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ContainerHandler) KillConfig(c *gin.Context) {
	var req containerinit.KillConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.grpcClient.KillConfig(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

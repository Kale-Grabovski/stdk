package api

import (
	"net/http"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetModule struct {
	logger      *zap.Logger
	versionRepo repository.VersionRepositoryInterface
}

func NewGetModule(
	logger *zap.Logger,
	versionRepo repository.VersionRepositoryInterface,
) domain.Action {
	return &GetModule{logger, versionRepo}
}

func (c *GetModule) Handle(ctx *gin.Context) {
	v, err := c.versionRepo.GetActiveVersionByModule(ctx.Param("id"))
	if err != nil {
		c.logger.Error("Cannot get active version", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	filePath := "files/" + v.Filename
	ctx.Header("X-HASH", v.Hash) // SDK should validate this header
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+filePath)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(filePath)
}

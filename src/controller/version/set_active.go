package version

import (
	"net/http"
	"strconv"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SetActive struct {
	logger      *zap.Logger
	versionRepo repository.VersionRepositoryInterface
}

func NewSetActive(
	logger *zap.Logger,
	versionRepo repository.VersionRepositoryInterface,
) domain.Action {
	return &SetActive{logger, versionRepo}
}

func (c *SetActive) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.logger.Error("Cannot get version ID from url")
		c.back(ctx)
		return
	}

	err = c.versionRepo.SetActive(id)
	if err != nil {
		c.logger.Error("Cannot set version active", zap.Error(err))
		c.back(ctx)
		return
	}

	c.back(ctx)
}

func (c *SetActive) back(ctx *gin.Context) {
	ref := ctx.GetHeader("Referer")
	if ref == "" {
		ref = "/"
	}

	ctx.Redirect(http.StatusMovedPermanently, ref)
}

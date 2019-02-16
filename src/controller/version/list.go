package version

import (
	"net/http"
	"strconv"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VersionsList struct {
	logger      *zap.Logger
	versionRepo repository.VersionRepositoryInterface
}

func NewVersionsList(
	logger *zap.Logger,
	versionRepo repository.VersionRepositoryInterface,
) domain.Action {
	return &VersionsList{logger, versionRepo}
}

func (c *VersionsList) Handle(ctx *gin.Context) {
	moduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.logger.Error("Cannot get module_id from url")
		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	versions, err := c.versionRepo.Get(moduleId)
	if err != nil {
		c.logger.Error("Cannot get versions", zap.Error(err), zap.Int("module_id", moduleId))
		ctx.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	ctx.HTML(http.StatusOK, "versions.tpl", gin.H{
		"versions": versions,
		"moduleId": moduleId,
	})
}

package api

import (
	"net/http"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InstalledModules struct {
	logger      *zap.Logger
	versionRepo repository.VersionRepositoryInterface
}

func NewInstalledModules(
	logger *zap.Logger,
	versionRepo repository.VersionRepositoryInterface,
) domain.Action {
	return &InstalledModules{logger, versionRepo}
}

func (c *InstalledModules) Handle(ctx *gin.Context) {
	var installedModules domain.InstalledModules

	if err := ctx.ShouldBind(&installedModules); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format of body data"})
		return
	}

	activeVersions, err := c.versionRepo.GetActiveVersions()
	if err != nil {
		c.logger.Error("Cannot get active versions", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	resp := domain.RevalidateModules{DeviceId: installedModules.DeviceId}
	for _, v := range installedModules.Modules {
		if v.Version != activeVersions[v.Id].Id {
			resp.Modules = append(resp.Modules, domain.RevalidateModule{
				Id:       v.Id,
				Version:  activeVersions[v.Id].Id,
				Settings: activeVersions[v.Id].Settings,
			})
		}
	}

	ctx.JSON(http.StatusOK, resp)
}

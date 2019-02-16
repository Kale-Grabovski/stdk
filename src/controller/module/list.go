package module

import (
	"net/http"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ModulesList struct {
	logger     *zap.Logger
	moduleRepo repository.ModuleRepositoryInterface
}

func NewModulesList(
	logger *zap.Logger,
	moduleRepo repository.ModuleRepositoryInterface,
) domain.Action {
	return &ModulesList{logger, moduleRepo}
}

func (c *ModulesList) Handle(ctx *gin.Context) {
	modules, err := c.moduleRepo.Get()
	if err != nil {
		c.logger.Error("Cannot get modules", zap.Error(err))
		return
	}

	ctx.HTML(http.StatusOK, "index.tpl", gin.H{
		"modules": modules,
	})
}

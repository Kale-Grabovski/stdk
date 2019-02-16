package module

import (
	"net/http"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ModuleCreate struct {
	logger     *zap.Logger
	moduleRepo repository.ModuleRepositoryInterface
}

func NewModuleCreate(
	logger *zap.Logger,
	moduleRepo repository.ModuleRepositoryInterface,
) domain.Action {
	return &ModuleCreate{logger, moduleRepo}
}

func (c *ModuleCreate) Handle(ctx *gin.Context) {
	var m domain.Module

	// Ignoring any errors
	if err := ctx.ShouldBind(&m); err == nil {
		err := c.moduleRepo.Create(m.Name)
		if err != nil {
			c.logger.Error("Cannot save module", zap.Error(err))
			return
		}
	}

	ctx.Redirect(http.StatusMovedPermanently, "/")
}

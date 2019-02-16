package version

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/Kale-Grabovski/stdk/src/service"

	"github.com/Kale-Grabovski/stdk/src/domain"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VersionCreate struct {
	logger       *zap.Logger
	versionRepo  repository.VersionRepositoryInterface
	cryptService service.CryptServiceInterface
}

func NewVersionCreate(
	logger *zap.Logger,
	versionRepo repository.VersionRepositoryInterface,
	cryptService service.CryptServiceInterface,
) domain.Action {
	return &VersionCreate{logger, versionRepo, cryptService}
}

func (c *VersionCreate) Handle(ctx *gin.Context) {
	moduleId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.logger.Error("Cannot get module_id from url")
		c.back(ctx)
		return
	}

	v := domain.Version{ModuleId: moduleId}

	// Ignoring any errors
	if err := ctx.ShouldBind(&v); err != nil {
		c.logger.Error("Validation failed", zap.Error(err))
		c.back(ctx)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		c.logger.Error("Cannot get file from POST request", zap.Error(err))
		c.back(ctx)
		return
	}

	v.Filename = c.genFilename(file.Filename)
	if err != nil {
		c.logger.Error("Cannot get file hash", zap.Error(err))
		c.back(ctx)
		return
	}

	err = ctx.SaveUploadedFile(file, "files/"+v.Filename)
	if err != nil {
		c.logger.Error("Cannot move uploaded file", zap.Error(err))
		c.back(ctx)
		return
	}

	v.Hash, err = c.cryptService.HashFile("files/" + v.Filename)
	err = c.versionRepo.Create(&v)
	if err != nil {
		c.logger.Error("Cannot save version", zap.Error(err))
		c.back(ctx)
		return
	}

	c.back(ctx)
}

func (c *VersionCreate) genFilename(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))[:16] + "_" + strconv.Itoa(int(time.Now().Unix()))
}

func (c *VersionCreate) back(ctx *gin.Context) {
	ref := ctx.GetHeader("Referer")
	if ref == "" {
		ref = "/"
	}

	ctx.Redirect(http.StatusMovedPermanently, ref)
}

package handlers

import (
	"github.com/denizgursoy/go-touch-projects/internal/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type (
	ResourceHandler struct {
		s ResourceService
		l *zap.Logger
	}
	ResourceService interface {
		GetResourceByID(id int64) (*services.Resource, error)
	}
)

func NewResourceHandler(e *echo.Echo, resourceService ResourceService) {
	handler := ResourceHandler{
		s: resourceService,
	}
	e.GET("/resource/:id", handler.GetResourceByID)
}

func (r *ResourceHandler) GetResourceByID(ctx echo.Context) error {
	resourceIDParam := ctx.Param("id")
	resourceID, err := strconv.ParseInt(resourceIDParam, 10, 64)

	if err != nil {
		return ctx.String(http.StatusBadRequest, "id must be integer")
	}

	resource, err := r.s.GetResourceByID(resourceID)
	if err != nil {
		r.l.Error(err.Error())
		return err
	}

	if resource == nil {
		return ctx.JSON(http.StatusNoContent, resource)
	} else {
		return ctx.JSON(http.StatusOK, resource)
	}
}

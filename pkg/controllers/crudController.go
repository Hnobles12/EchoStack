package controllers

import (
	"EchoStack/EchoStack/repositories"
	"net/http"

	"github.com/labstack/echo"
)

// Controller handles CRUD operations for a specific model type.
type Controller[T any] struct {
	Repo *repositories.Repository[T]
}

// NewController creates a new Controller instance.
func NewController[T any](repo *repositories.Repository[T]) *Controller[T] {
	return &Controller[T]{Repo: repo}
}

// CreateHandler handles creating a new entity.
func (c *Controller[T]) CreateHandler(ctx echo.Context) error {
	var entity T
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.Repo.Create(&entity); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, entity)
}

// FindAllHandler handles retrieving all entities.
func (c *Controller[T]) FindAllHandler(ctx echo.Context) error {
	entities, err := c.Repo.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, entities)
}

// FindByIDHandler handles retrieving a single entity by ID.
func (c *Controller[T]) FindByIDHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	entity, err := c.Repo.FindByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, entity)
}

// UpdateHandler handles updating an entity.
func (c *Controller[T]) UpdateHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	var entity T
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// Ensure the ID is correctly set on the entity.
	if err := c.Repo.DB.Model(&entity).Where("id = ?", id).Updates(entity).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, entity)
}

// DeleteHandler handles deleting an entity.
func (c *Controller[T]) DeleteHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := c.Repo.DeleteByID(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

// RegisterRoutes adds the standard CRUD routes to an Echo group.
func (c *Controller[T]) RegisterRoutes(group *echo.Group) {
	group.POST("", c.CreateHandler)
	group.GET("", c.FindAllHandler)
	group.GET("/:id", c.FindByIDHandler)
	group.PUT("/:id", c.UpdateHandler)
	group.DELETE("/:id", c.DeleteHandler)
}

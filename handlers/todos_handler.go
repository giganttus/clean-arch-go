package handlers

import (
	"net/http"
	"strconv"
	"todoapp/helpers"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateTodo(ctx *gin.Context) {
	var newTodo models.Todos

	if err := ctx.ShouldBind(&newTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.CreateTodo(ctx, newTodo)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"response": "success"})
}

func (h *Handlers) GetTodos(ctx *gin.Context) {
	res, err := h.Services.GetTodos(ctx)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": res})
}

func (h *Handlers) UpdateTodo(ctx *gin.Context) {
	var upTodo models.Todos

	if err := ctx.ShouldBind(&upTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err := h.Services.UpdateTodo(ctx, upTodo)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}

func (h *Handlers) DeleteTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"response": "Bad Request"})
		return
	}

	err = h.Services.DeleteTodo(ctx, todoID)
	if err != nil {
		if err.Error() == helpers.ErrForbidden.Error() {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"response": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "success"})
}

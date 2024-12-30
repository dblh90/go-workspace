package internals

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"go-fullstack/views"
)

const appTimout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *Config) indexPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimout)
		defer cancel()

		todos, err := app.getAllTodosService()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		var viewsTodos []*views.Todo
		for _, t := range todos {
			viewsTodo := &views.Todo{
				Id:          t.Id,
				Description: t.Description,
			}
			viewsTodos = append(viewsTodos, viewsTodo)
		}
		render(ctx, http.StatusOK, views.Index(viewsTodos))
	}
}

func (app *Config) createTodoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimout)
		description := ctx.PostForm("description")
		defer cancel()

		newTodo := TodoRequest{
			Description: description,
		}

		data, err := app.createTodoService(&newTodo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, data)
	}
}

func (app *Config) deleteTodoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimout)
		id := ctx.Param("id")

		defer cancel()

		data, err := app.deleteTodoService(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusAccepted, fmt.Sprintf("Todo with ID: %s deleted successfully", data))
	}
}

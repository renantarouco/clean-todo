package api

import (
	"github.com/gin-gonic/gin"
	"github.com/renantarouco/clean-todo/core"
)

type TasksAPI struct {
	router  *gin.Engine
	service *core.TasksService
}

func NewTasksAPI(service *core.TasksService) *TasksAPI {
	api := &TasksAPI{
		router:  gin.New(),
		service: service,
	}
	api.SetupRoutes()
	return api
}

func (api *TasksAPI) Run() error {
	return api.router.Run()
}

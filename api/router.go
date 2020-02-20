package api

func (api *TasksAPI) SetupRoutes() {
	api.router.GET("tasks", api.listHandler)
	api.router.POST("/tasks", api.addHandler)
}

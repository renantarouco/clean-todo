package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renantarouco/clean-todo/model"
)

func (api *TasksAPI) listHandler(c *gin.Context) {
	tasks, err := api.service.List()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (api *TasksAPI) addHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	task := model.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := api.service.Add(task); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusCreated, "task added")
}

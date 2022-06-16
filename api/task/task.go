package task

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"pandora/db"
	"pandora/ent/task"
	"time"
)

type Task struct {
	name     string
	taskType task.Type
	describe string
	f        func()
}

func NewTask(name string, taskType task.Type, describe string, f func()) Task {
	return Task{
		name,
		taskType,
		describe,
		f,
	}
}

func (t *Task) Start(c *gin.Context) {
	ctx := context.Background()
	startTime := time.Now()
	userId, ok := c.Get("userId")
	if !ok {
		panic("no current user")
	}
	tk, err := db.Client.Task.Create().
		SetName(t.name).
		SetType(t.taskType).
		SetDescribe(t.describe).
		SetStartDate(startTime).
		SetStatus(1).
		SetCreatedAt(startTime).
		SetCreatedBy(userId.(int)).
		Save(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("in defer")
		endTime := time.Now()
		cost := endTime.Sub(startTime).Seconds()
		tku := tk.Update()
		tku.SetEndDate(endTime).SetCostTime(int(cost))
		if p := recover(); p != nil {
			tku.SetStatus(3)
		} else {
			tku.SetStatus(2)
		}
		_, err := tku.Save(ctx)
		if err != nil {
			panic(err)
		}
	}()
	t.f()
	fmt.Println("down...")
}

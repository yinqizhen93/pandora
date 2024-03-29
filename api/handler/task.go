package handler

import (
	"context"
	"fmt"
	"pandora/ent"
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

func (t *Task) Start(db *ent.Client, createdBy int) {
	ctx := context.Background()
	startTime := time.Now()
	tk, err := db.Task.Create().
		SetName(t.name).
		SetType(t.taskType).
		SetDescribe(t.describe).
		SetStartDate(startTime).
		SetStatus(1).
		SetCreatedAt(startTime).
		SetCreatedBy(createdBy).
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

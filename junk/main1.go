package asyncq

import (
  "log"
)

var TaskQueue chan Task
var TaskWorkerQueue chan chan Task

type TaskWorker struct {
  ID int
  TaskChannel chan Task
  TaskWorkerQueue chan chan Task
}

func NewTaskWorker(id int, taskWorkerQueue chan chan Task) {
  task := TaskWorker{
    ID: id,
    TaskChannel: make(chan Task),
    TaskWorkerQueue: taskWorkerQueue
  }
  return taskWorkerQueue
}

func (t *TaskWorker) Start() {
  go func() {
    for {
      t.TaskWorkQueue <- t.TaskChannel
      select {
      case task := <-t.TaskChannel:
        log.Printf("TaskWorker [%d] Started.\n", t.ID)
        task.Perform()
      }
    }
  }
}


// Accepts task from the TaskQueue and dispatches to next TaskWorker
func StartTaskDispatcher(taskWorkerSize int) {
  TaskWorkerQueue = make(chan chan Task, taskWorkerSize)

  for i := 0; i < taskWorkerSize; i++ {
    log.Print("Starting TaskWorker #%d\n", i+1)
    taskWorker := NewTaskWorker(i+1, TaskWorkerQueue)
  }
}

func init() {
  TaskQueue = make(chan Task, 108)
}

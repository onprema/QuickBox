package asyncq

type Task interface {

  // Where the task gets carried out
  Perform()
}

package fastHands

import (
  "fmt"
  "time"
  "math/rand"
)

var stations int = 5
var repetitions int = 2
// var difficulty []int = []int{1,2,3,4,5}

func build_training_center() []int {
  board := make([]int, stations)
  for i := 0; i < stations; i ++ {
    board[i] = i + 1
  }
  return board
}

func getExercise(level int){
  exercise := make([]int, level)
  for i := 0; i < level; i ++ {
    rand.Seed(time.Now().UnixNano())
    randomNum := int(rand.Intn(stations - 1) + 1)
    if i > 0 {
      if randomNum == 0 || exercise[i -1] == randomNum {
        exercise[i] = randomNum + 1
      } else {
        exercise[i] = randomNum
      }
    } else {
      exercise[i] = randomNum
    }
  }
  fmt.Println(exercise)
}

func Run() {
  for station := 0; station < stations; station ++ {
    for rep := 1; rep <= repetitions; rep ++ {
      getExercise(station + 1)
      time.Sleep(5 * time.Second)
    }
  }
}

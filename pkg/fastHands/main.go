package fastHands

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Data struct {
	FastHands map[string]map[string]interface{} `yaml:"FastHands"`
	MetaData  map[string]string                 `yaml:"MetaData"`
	Weights   map[string]map[string]interface{} `yaml:"Weights"`
}

func Run() Data {
	var data Data
	yamlFile, err := ioutil.ReadFile("drills.yaml")
	if err != nil {
		log.Printf("Error reading yaml file %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Cant unmarshal yaml %v", err)
	}

	return data

}

// var stations int = 5
// var repetitions int = 2
// // var difficulty []int = []int{1,2,3,4,5}
//
// func build_training_center() []int {
//   board := make([]int, stations)
//   for i := 0; i < stations; i ++ {
//     board[i] = i + 1
//   }
//   return board
// }
//
// func getExercise(level int){
//   exercise := make([]int, level)
//   for i := 0; i < level; i ++ {
//     rand.Seed(time.Now().UnixNano())
//     randomNum := int(rand.Intn(stations - 1) + 1)
//     if i > 0 {
//       if randomNum == 0 || exercise[i -1] == randomNum {
//         exercise[i] = randomNum + 1
//       } else {
//         exercise[i] = randomNum
//       }
//     } else {
//       exercise[i] = randomNum
//     }
//   }
//   fmt.Println(exercise)
// }
//
// func Run() {
//   for station := 0; station < stations; station ++ {
//     for rep := 1; rep <= repetitions; rep ++ {
//       getExercise(station + 1)
//       time.Sleep(5 * time.Second)
//     }
//   }
// }
//

package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main(){
//   fmt.Println(os.Args)
//   args := os.Args[1:]
//   var iArgs = []int{}
//   for _, value := range args {
//     val, err := strconv.Atoi(value)
//     if err != nil {
//       panic(err)
//     }
//     iArgs = append(iArgs, val)
//   }
//   max := 0
//   for _, val := range iArgs {
//     if val > max {
//       max = val
//     }
//   }
//   fmt.Println("Max value:", max)
// }

package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)


func main(){
  fmt.Println("Hello", stuff.Name)
  intArr := []int{2, 3, 5, 7}
  strArr := stuff.IntArrToStrArr(intArr)
  fmt.Println(strArr)
  fmt.Println(reflect.TypeOf(strArr))

  date := stuff.Date{}
  err := date.SetMonth(12)
  if err != nil {
    log.Fatal(err)
  }

  dErr := date.SetDay(21)
  if dErr != nil {
    log.Fatal(err)
  }
  
  mErr := date.SetYear(1974)
  if mErr != nil {
    log.Fatal(err)
  }
  fmt.Printf("1st Day: %d/%d/%d\n", date.Day(), date.Month(), date.Year())
}

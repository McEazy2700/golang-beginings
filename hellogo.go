package main

import (
	// "bufio"
	"bufio"
	"errors"
	"fmt"

	// "log"

	// "log"
	// "math/rand"
	// "time"

	// "unicode/utf8"
	// "strings"
	// "log"
	// "strconv"
	// "reflect"
	// "log"
	"os"
)

// Alias for print with new line
var printLn = fmt.Println

// An alias for printing formats
var printF = fmt.Printf

// An alial for normal printing
var prnt = fmt.Print

// An alias to bufio.NewReader
var reader = bufio.NewReader(os.Stdin)

// Returns the sum and reduced result of two numbers
func getSumReduce(x int, y int) (int, int) {
  return x + y, x - y
}

/*********
Concurrency
*********/

func printTo15(){
  for i := 0; i <= 15; i++ {
    printLn("Func1:", i)
  }
}

func printTo10(){
  for i := 0; i <= 10; i++ {
    printLn("Func2:", i)
  }
}

/*********
Channels
*********/
func nums1(channel chan int) {
  channel <- 1
  channel <- 2
  channel <- 3
}

func nums2(channel chan int) {
  channel <- 4
  channel <- 5
  channel <- 6
}
/*********
Interfaces
*********/

type Animal interface {
  AngrySound()
  HappySound()
}

type Cat string

func (c Cat) Attack(){
  printLn("Cat attacks it's prey")
}

func (c Cat) Name() string {
  return string(c)
}

func (c Cat) AngrySound() {
  printLn("Cat says Hisssss")
}


func (c Cat) HappySound() {
  printLn("Cat says Purrr")
}

/**********
Types
**********/
type MyConstraint interface {
  int | float64
}

type Tsp float64
type TBs float64
type ML float64

func (tsp Tsp) ToML() ML {
  return ML(tsp * 4.92)
}

func (tbs TBs) ToML() ML {
  return ML(tbs * 14.79)
}

func tsToML(tsp Tsp) ML {
  return ML(tsp * 4.92)
}

func tBsToML(tBs TBs) ML {
  return ML(tBs * 14.79)
}

/*********
Structs
*********/

// type customer struct {
//   name string
//   address string
//   bal float64
// }

// func getCustomerInfo(c customer){
//   printF("%s owes us %.2f\n", c.name, c.bal)
// }

// func newCustomerAdd(c *customer, address string){
//   c.address = address
// }

// type rectangle struct {
//   length, height float64
// }

// func (r rectangle) Area() float64 {
//   return r.length * r.height
// }

// type contact struct {
//   firstName, lastName, phone string
// }

// type buisiness struct {
//   name, address string
//   contact 
// }

// func (b buisiness) info() {
//   printF("Contact at %s is %s %s\n", b.name, b.contact.firstName, b.contact.lastName)
// }

func getSumGen[T MyConstraint](x T, y T) T {
  return x + y
}

// Returns the division of two values
func getQuotient(x float64, y float64) (answ float64, err error) {
  if y == 0{
    return 0, errors.New("You can't divide by 0")
  }else {
    return x /y, nil
  }
}

// Returns the sum of an array of values
func getSum(nums ...int) int {
  sum := 0
  for _, num := range nums {
    sum += num
  }
  return sum
}

func changeVal(ptr *int, val int){
  *ptr = val
}

func main(){
  // printLn("What is you name")
  // reader := bufio.NewReader(os.Stdin)
  // name, err := reader.ReadString('\n')
  // if err == nil {
  //   printLn("Hello ", name)
  // }else {
  //   log.Fatal(err)
  // }

  // // Type Declared and initialized variable
  // var meName string = "Eazy"
  // // Simultaneously type declared and initialized variables
  // var meBestNo, meAge int = 7, 20
  // // Type inferred variable
  // var v3 = "Hello"
  // // Shorthand initialized and inferred variable
  // v4 := 24

  /*********
  Data Types
  *********/
  // int, float64, bool, string, rune
  // Defaults: 0, 0.0, false, ""
  // printLn(reflect.TypeOf(25))
  // printLn(reflect.TypeOf(3.15))
  // printLn(reflect.TypeOf(true))
  // printLn(reflect.TypeOf("Hello"))
  // printLn(reflect.TypeOf("🦍"))

  /********
  Castring
  *******/
  // cV1 := 1.5
  // cV2 := int(cV1)
  // printLn(cV2)
  // // String to integer
  // cV3 := "50000000"
  // cV4, err := strconv.Atoi(cV3)
  // if err != nil {
  //   log.Fatal(err)
  // }
  // printLn(cV4)
  // // Int to string
  // cV5 := 5000000
  // cV6 := strconv.Itoa(cV5)
  // printLn(cV6)

  /*********
  Conditionals && Control Flow
  *********/
  // iAge := 8
  // if (iAge >= 1) && (iAge <= 18) {
  //   printLn("Important Bithday")
  // } else if (iAge == 21) || (iAge == 50) {
  //   printLn("Important Bithday")
  // } else if (iAge >= 65) {
  //   printLn("Important Bithday")
  // } else {
  //   printLn("Not an Important Bithday")
  // }
  // printLn("!true =", !true)
  /*********
  Strings: Arrays of bytes []byte
  *********/
  // sV1 := "A Word"
  // replacer := strings.NewReplacer("A", "Another")
  // sV2 := replacer.Replace(sV1)
  // printLn(sV2)
  // printLn("Length:", len(sV2))
  // printLn("Contains Another:", strings.Contains(sV2, "Another"))
  // printLn("o's index in string:", strings.Index(sV2, "o"))
  // printLn("Replace :", strings.Replace(sV2, "o", "0", -1))
  // sV3 := "\nSome word\n"
  // sV3 = strings.TrimSpace(sV3)
  // printLn(sV3)
  // printLn("Split :", strings.Split("a-b-c-d", "-"))
  // printLn("Lower :", strings.ToLower(sV2))
  // printLn("Lower :", strings.ToUpper(sV2))
  // printLn("Prefix :", strings.HasPrefix("tacocat", "taco"))
  // printLn("Prefix :", strings.HasSuffix("tacocat", "cat"))
  /**********
  Runes: In Go characters are called runes
  rundes are unicodes that represent characters
  **********/
  // rStr := "abcefg"
  // printLn("Rune Count:", utf8.RuneCountInString(rStr))
  // for i, runVal := range rStr {
  //   printF("{index: %d, rune: %#U, char: %c}\n", i, runVal, runVal)
  // }

  /********
  Time
  ********/
  // now := time.Now()
  // printLn(now.Year(), now.Month(), now.Day())
  // printF("%d:%d:%d\n", now.Hour(), now.Minute(), now.Day())

  /*******
  Math
  *******/
  // printLn("5 + 5 =", 5 + 4)
  // printLn("5 - 5 =", 5 - 4)
  // printLn("5 * 5 =", 5 * 4)
  // printLn("5 / 5 =", 5 / 4)
  // printLn("5 % 5 =", 5 % 4)
  // mInt := 1
  // mInt ++

  /*******
  Random values
  needs seed value for this action
  *******/
  // seedSecs := time.Now().Unix()
  // rand.Seed(seedSecs)
  // randNum := rand.Intn(50) + 1
  // printLn("Random:", randNum)
  // // Math function are many
  
  /*********
  Formatted print
  *********/
  /*
  %d: Integer
  %c: Character
  %f: Float
  %t: Boolean
  %s: String
  %o: Base 8
  %x: Base 16
  %v: Guesses bassed on data type. I'll take this as any variable
  %T: Type of supplied value
  */

  /*********
  For loops or Loops in general
  *********/
  // // for initialization (pre-run); conditional (to be met); postStatement (after run) {body}
  // // very similar to JavaScript in many ways
  // for x := 0; x <= 5; x++ {
  //   printLn(x)
  // }
  // // x only exists within the scope of the for loop
  // fX := 0
  // // fX exists outside for loop
  // for fX <=5 {
  //   printLn(fX)
  //   fX++
  // }
  /**********
  For for while loops
  Gueesing game
  **********/
  // seedSecs := time.Now().Unix()
  // rand.Seed(seedSecs)
  // tries := 0
  // randNum := rand.Intn(10) + 1

  // for tries <= 5 {
  //   prnt("Guess a number between 0 and 10: ")

  //   guess, err := reader.ReadString('\n')
  //   if err != nil {
  //     log.Fatal(err)
  //   }

  //   guess = strings.TrimSpace(guess)
  //   userGuess, err := strconv.Atoi(guess)
  //   if err != nil {
  //     log.Fatal(err)
  //   }

  //   if userGuess == randNum {
  //     printLn("\n")
  //     printLn("Your guess", userGuess, "was correct.")
  //     printLn("You win!")
  //     break
  //   } else if (userGuess < randNum) && (tries <= 4){
  //     printLn("\n")
  //     printLn("Awww.. the number you guessed is to low.")
  //     printLn("Guess higher")
  //   } else if (userGuess > randNum) && (tries <= 4){
  //     printLn("\n")
  //     printLn("So close but your guess is to high.")
  //     printLn("Guess lower")
  //   }

  //   tries++

  //   if tries == 5 {
  //     printLn("\n")
  //     printLn("You have failed to guess the number in 4 tries.")
  //     printLn("You have one more try")
  //   }
  //   if tries > 5 {
  //     printLn("\n")
  //     printLn("You have failed to guess the number in 5 tries.")
  //     printLn("You loose")
  //   }
  // }

  /*******
  Cycling through arrays
  ******/
  // aNums := []int{1, 2, 4}
  // for index, value := range aNums {
  //   printF("{index: %v, value: %v}\n", index, value)
  // }

  /*******
  Arrays: Collection of values with the same data type
  * Size of the array cannot change
  *******/

  // // Declare the array
  // var arr1 [5]int
  // // initialize it latter
  // arr1[0] = 23
  // // Declare and initialize together
  // arr2 := [5]int{2, 4, 5, 6, 7}
  // printLn("Index 0:", arr2[0])
  // printLn("Array length:", len(arr2))
  // // iterate over array
  // for i := 0; i < len(arr2); i++ {
  //   printLn(arr2[i])
  // }
  // // iterate with range
  // for i, v := range arr2 {
  //   printLn(i, v)
  // }

  /********
  Slice: Collection that can grow
  ********/
  // sl1 := make([]string, 6)
  // sl1[0] = "Eazy"

  /*******
  Functions
  *******/
  // printLn(getSumReduce(56, 34))
  // printLn(getQuotient(23, 4))
  // printLn(getSum(1, 2, 3, 5))

  /******
  Pointers
  ******/
  // val := 34
  // printLn(val)
  // changeVal(&val, 22)
  // printLn(val)

  /******
  Files
  ******/
  // file, err := os.Create("data.txt")
  // if err != nil {
  //   log.Fatal(err)
  // }
  // defer file.Close()
  // iPrimeArr := []int{2, 3, 5, 7, 11}
  // var sPrimeArr []string
  // for _, val := range iPrimeArr {
  //   sPrimeArr = append(sPrimeArr, strconv.Itoa(val))
  // }
  // for _, str := range sPrimeArr {
  //   _, err = file.WriteString(str + "\n")
  //   if err != nil {
  //     log.Fatal(err)
  //   }
  // }
  // file, err = os.Open("data.txt")
  // if err != nil {
  //   log.Fatal(err)
  // }
  // defer file.Close()
  // scan1 := bufio.NewScanner(file)
  // for scan1.Scan() {
  //   printLn("Prime :", scan1.Text())
  // }
  // if err := scan1.Err(); err != nil {
  //   log.Fatal(err)
  // }
  
  // _, err := os.Stat("data.txt")
  // if errors.Is(err, os.ErrNotExist) {
  //   printLn("File Doesn't exist")
  // }else {
  //   file, err := os.OpenFile("data.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
  //   if err != nil {
  //     log.Fatal(err)
  //   }
  //   defer file.Close()
  //   if _, err := file.WriteString("13\n"); err != nil {
  //     log.Fatal(err)
  //   }
  // }

  /**********
  Maps: collections of key: value pairs
  a key is anything that can be compared using '='
  the key can be a completely diff data type that the value
  **********/
  // var heroes map[string]string
  // heroes = make(map[string]string)
  // villians := make(map[string]string)
  // heroes["Batman"] = "Bruce Wayne"
  // heroes["Sueperman"] = "Clark Kent"
  // heroes["The Flash"] = "Barry Allen"
  // villians["Penguin"] = "Oswald Coblepot"
  // villians["Lex Luther"] = "Lex Luther"

  // superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
  // printF("Batman is %v\n", heroes["Batman"])
  // println("Chip :%v", superPets[3])
  // _, ok := superPets[3]
  // printLn("Is there a third pet?", ok)
  // for key, value := range heroes {
  //   printF("%s is %s\n", key, value)
  // }

  /************
  Generics
  ************/
  // printLn("5 + 4 =", getSumGen(5, 4))
  // printLn("5.6 + 4.7 =", getSumGen(5.6, 4.7))

  /**********
  Structs: allow you to store values with different data types in a structured way
  **********/
  // var cus1 customer
  // cus1.name = "Vice Codes"
  // cus1.address = "The SanFransisco Bay area"
  // cus1.bal = 234.56

  // getCustomerInfo(cus1)
  // newCustomerAdd(&cus1, "Silicon Valley")
  // printLn("Address:", cus1.address)
  // cus2 := customer{"Sally Smith", "Toky Japan", 0.0}
  // printF("%s owes us %.2f\n", cus2.name, cus2.bal)

  // rect1 := rectangle{10.0, 23.5}
  // printLn("Rect area: ", rect1.Area())

  /*************
  Inheritance & Composition
  Go does not support inheritance, but it does support composition
  which can be achieved by embeding a struct inside another struct
  *************/
  // con1 := contact{firstName: "John", lastName: "Smith", phone: "23534565"}
  // bus1 := buisiness{
  //   name: "ABC Plumbing",
  //   address: "234 North Streen",
  //   contact: con1,
  // }
  // bus1.info()

  /***********
  Types: You can also define types that enhance other built in data types
  ***********/
  // ml1 := ML(Tsp(3) * 4.92)
  // printF("3 tsps = %.2f ML\n", ml1)
  // ml2 := ML(TBs(3) * 14.79)
  // printF("3 TBs = %.2f ML\n", ml2)
  // printLn("3tsp + 2tsp =", Tsp(3) + Tsp(2))
  // printLn("3tsp > 2tsp =", Tsp(3) > Tsp(2))
  // printF("3tsp = %.2fml \n", tsToML(3))
  // printF("3TBs = %.2fml \n", tBsToML(3))

  /**********
  Type associate methods
  **********/
  // printF("3tsp = %.2f ml\n", Tsp(3).ToML())
  // printF("3tbs = %.2f ml\n", TBs(3).ToML())

  /**********
  Interfaces: allow you to define contracts
  that any thing composed of that interface must abind to
  **********/

  // var kitty1 Animal
  // kitty1 = Cat("Kitty")
  // kitty1.AngrySound()

  /*********
  Concurrency: allow tasks to run simultaneously on different threads
  concurrent tasks in Go are called Go routines
  *********/
  // to carry run a function as a go routine just add go in front of it
  // With go routines you cannot trust the order in which they will run
  // to fix that we use Channels

  // go printTo15()
  // go printTo10()
  // time.Sleep(2 * time.Second)

  // chan1 := make(chan int)
  // chan2 := make(chan int)
  // go nums1(chan1)
  // go nums2(chan2)
  // printLn(<-chan1)
  // printLn(<-chan1)
  // printLn(<-chan1)
  // printLn(<-chan2)
  // printLn(<-chan2)
  // printLn(<-chan2)
}

package main

import (
	"bufio"
	"fmt"
	"gopher/internal/game"
	"gopher/internal/types"
	"gopher/internal/utils"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// -------------------------------------
// Constants and Enumerations
// -------------------------------------

const (
	First = iota + 1 // using iota for enumeration
	Second
	_ // Skips the third value
	Fourth
	Fifth
)

func main() {
	demonstrateBasics()
	demonstrateConcurrency()
	demonstrateControlStructures()
	demonstratePointers()
	demonstrateErrorHandling()
	demonstrateComplexTypes()
	demonstrateInterfaces()
	demonstrateUserInputAndFileHandling()
	demonstrateAdvancedConcurrency()
	demonstrateInheritanceLikeBehavior()
	game.DemonstrateGuessGame()
	demonstrateUpdateIPointer()
}

// -------------------------------------
// Basic Demonstrations
// -------------------------------------

func demonstrateBasics() {
	// Demonstrating various basic features of Go
	var str = new(string)
	*str = "Denis"
	greeting, age := SayHello(str)
	fmt.Printf("Greeting: %s, Age: %d\n", greeting, age)

	// Constants and variables
	const constTest string = "this is a const"
	var varTest string = "this is a var"
	fmt.Println(constTest, varTest)

	// Short variable declaration
	shortVar := "short variable declaration"
	fmt.Println(shortVar)

	// Multiple variable declaration
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// Type inference
	inferredType := "inferred type"
	fmt.Printf("Type: %T, Value: %v\n", inferredType, inferredType)

	// Demonstrating iota
	fmt.Println("Iota values:", First, Second, Fourth, Fifth)
}

// -------------------------------------
// Additional Supporting Functions
// -------------------------------------

func SayHello(name *string) (greeting string, age int) {
	greeting, _ = utils.InterpolFormat("Hello {1}", *name)
	age = 1
	fmt.Println(greeting)
	return // implicit return of greeting and age
}

// -------------------------------------
// Concurrency Examples
// -------------------------------------

func demonstrateConcurrency() {
	var counter int64
	var wg sync.WaitGroup

	// Using goroutines and wait groups
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Final count:", atomic.LoadInt64(&counter))
}

// -------------------------------------
// Control Structures: If, Switch, Loops
// -------------------------------------

func demonstrateControlStructures() {
	// Using switch with custom type enumeration
	car := types.Volkswagen
	switch car {
	case types.Volkswagen:
		fmt.Println("Volkswagen")
	case types.Ford:
		fmt.Println("Ford")
	case types.Toyota:
		fmt.Println("Toyota")
	case types.Kia:
		fmt.Println("Kia")
	case types.Hyundai:
		fmt.Println("Hyundai")
	}

	// For loop equivalent to while
	num := 0
	for num < 10 {
		num++
	}
	fmt.Println("Number:", num)

	// If-else statement
	if num >= 5 {
		fmt.Println("Number is greater than or equal to 5")
	} else {
		fmt.Println("Number is less than 5")
	}

	// Range-based for loop
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

func demonstratePointers() {
	// Using pointers to modify data
	x := 10
	p := &x
	*p = 20
	fmt.Println("Modified value through pointer:", x)

	// Using new to allocate memory
	ptr := new(int)
	*ptr = 30
	fmt.Println("Value assigned to new memory location:", *ptr)
}

// -------------------------------------
// Error Handling
// -------------------------------------

func demonstrateErrorHandling() {
	div, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", div)
	}
}

func safeDivide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, &MyError{"division by zero"}
	}
	return dividend / divisor, nil
}

// -------------------------------------
// Complex Types: Maps, Slices, Arrays
// -------------------------------------

func demonstrateComplexTypes() {
	// Using slices and maps
	mySlice := []int{1, 2, 3}
	myMap := map[int]string{1: "one", 2: "two"}
	fmt.Println("Slice:", mySlice, "Map:", myMap)

	// Arrays
	myArray := [3]string{"1", "2", "3"}
	fmt.Println("Array:", myArray)

	// Appending to slices
	mySlice = append(mySlice, 4)
	fmt.Println("Appended slice:", mySlice)
}

// -------------------------------------
// Interfaces and Polymorphism
// -------------------------------------

func demonstrateInterfaces() {
	var t Turner = &Truck{}
	t.TurnOn()
	fmt.Println(t)
}

// -------------------------------------
// Interfaces and Methods
// -------------------------------------

type Turner interface {
	TurnOn()
}

// TurnOn sets initial values for a Truck
func (t *Truck) TurnOn() {
	t.time = 1
	t.timezone = "UTC"
}

// String implements the Stringer interface for Truck
func (t *Truck) String() string {
	return fmt.Sprintf("Time: %v, Timezone: %s", t.time, t.timezone)
}

// -------------------------------------
// User Input and File Handling
// -------------------------------------

func demonstrateUserInputAndFileHandling() {
	fmt.Println("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s!\n", name)

	// File handling: Writing to a file
	//file, err := os.Create("example.txt")
	//if err != nil {
	//	fmt.Println("Error creating file:", err)
	//	return
	//}
	//defer file.Close()
	//_, err = file.WriteString("Hello, file handling in Go!\n")
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//}
	//fmt.Println("File written successfully")
}

// -------------------------------------
// Advanced Concurrency with Channels
// -------------------------------------

func demonstrateAdvancedConcurrency() {
	ch := make(chan string)

	// Producer
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "data from goroutine"
		close(ch)
	}()

	// Consumer
	for data := range ch {
		fmt.Println("Received:", data)
	}
	fmt.Println("Channel closed, no more data.")
}

// -------------------------------------
// Inheritance-Like Behavior with Type Embedding
// -------------------------------------

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound.")
}

type Dog struct {
	Animal // Embedding Animal struct
	Breed  string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "barks.")
}

func demonstrateInheritanceLikeBehavior() {
	dog := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Labrador",
	}
	dog.Speak() // Calls Dog's Speak method, demonstrating polymorphism

	// Accessing embedded type method
	dog.Animal.Speak()
}

func foo(t *int) {
	k := 101
	t = &k
}

func bar(t *int) {
	k := 102
	*t = k
}

func car(t int) {
	k := 103
	t = k
}

func demonstrateUpdateIPointer() {
	i := 100
	fmt.Println(i) // 100

	foo(&i)
	fmt.Println(i) // 100 - foo does not change i

	bar(&i)
	fmt.Println(i) // 102 - bar changes i to 102

	car(i)
	fmt.Println(i) // 102 - car does not change i
}

// -------------------------------------
// Custom Types and Aliases
// -------------------------------------

type (
	MyError struct {
		message string
	}

	Truck struct {
		time     float64
		timezone string
	}
)

// Error implements the error interface for MyError
func (e *MyError) Error() string {
	return e.message
}

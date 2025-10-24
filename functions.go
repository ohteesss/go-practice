package main

import (
	"fmt"
	"strings"
)


type User struct {
	Name string
	Age int
	Location string
}

func add(x int, y int) (string, int) {
	return fmt.Sprintln(x + y), x +y
}

func split(sum int) (x, y int){
	x = sum * 4/9
	y = sum - x
	return
}

func sum(nums []int) int{
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

//Variadic function
func sumV(str string, nums ...int)int{
	 total := 0
    for _, num := range nums {
        total += num
    }
	return total
}

func joinStrings(sep string, words ...string)string{
	if len(words) == 0 {
		return ""
	}
	result := words[0]
	for i:=1; i < len(words); i++ {
		result += sep + words[i]
	}
	return result
}

func createUser(name string, attrs ...interface{}) User {
	user := User{
		Name: name,
		Age: 23,
		Location: "Unknown",
	}

	for i := 0; i <  len(attrs); i+= 2{
		key := attrs[i].(string)
		value := attrs[i+1]
		switch strings.ToLower(key){
		case "age":
			user.Age = value.(int)
		case "location":
			user.Location = value.(string)
		}
	}
	return  user
}

// func calculateStats(nums ...int)(sum int, max int, avg float64){
// 	if len(nums) == 0 {
// 		return 0, 0, 0
// 	}

// 	for _,val := range nums{
// 		if val >= max {
// 			max = val
// 		}
// 		sum += val
// 	}
	

// 	return sum, max, float64(sum/ len(nums))

// }
func calculateStats(nums ...int)(sum int, max int, avg float64){
	if len(nums) == 0 {
		return 0, 0, 0
	}

	for _,val := range nums{
		if val >= max {
			max = val
		}
		sum += val
	}
	
	avg = float64(sum/ len(nums))

	return 

}

func divideNumbers(a int, b int)(ans float64, err error){
	if b == 0 {
		err = fmt.Errorf("Second parameter cannot be zero")
		return
	}

	ans = float64(a/b)

	return
}
type ValidationError struct {
    Field string
    Value interface{}
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s = %v", e.Field, e.Value)
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x */
   *x = *y      /* put y into x */
   *y = temp    /* put temp into y */
}
// func swap(x int, y int) {
//    var temp int
//    temp = x    /* save the value at address x */
//    x = y      /* put y into x */
//    y = temp    /* put temp into y */
// }

func main(){
	var anything interface{}

	

	var newError ValidationError

	newError = ValidationError{
		Field: "name",
		Value: "cannot be void",
	}

	newError.Error()

	

	numbers := []int{10, 20, 30, 40, 50}
	anything = "tobi"

	val := fmt.Sprint(anything.(string))

	words := []string{"Software Engineer", "Entrepreneur"}
	fmt.Println( add(23,45))
	fmt.Println(split(45))
	fmt.Println(sum([]int{1,2,3}))
	fmt.Println(sumV("Sum of some random numbers" ,1,5,6,7,34,78))
	fmt.Println(joinStrings(" | ",words... ))
	fmt.Println(val)
	user1 := createUser("Alice")
    user2 := createUser("Bob", "age", 25, "location", "New York")
    fmt.Printf("%v\n", user1) 
    fmt.Printf("%+v\n", user2) 
	fmt.Println(calculateStats(1,2,3))

	fmt.Println(divideNumbers(5,0))
	total, maximum, average := calculateStats(numbers...)
    fmt.Printf("Total: %d, Maximum: %d, Average: %.2f\n", total, maximum, average)


	// Anonymous functions

	( func(){
        fmt.Println("Inside anonymous function")
    })()

	 v := func(){
        fmt.Println("Inside another anonymous function")
    }
	g := func(val string) string{
		fmt.Println(val)
		return val
	}

	func (v string, g func(v string) string)  {

		g(v)
		
	}("hello", g)

	v()


	   /* local variable definition */
   var a int = 100
   var b int = 200

   fmt.Printf("Before swap, value of a : %d\n", a )
   fmt.Printf("Before swap, value of b : %d\n", b )

   /* calling a function to swap the values */
   swap(&a, &b)

   fmt.Printf("After swap, value of a : %d\n", a )
   fmt.Printf("After swap, value of b : %d\n", b )

	// call by value




}
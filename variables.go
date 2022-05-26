package main

import "fmt"

func main() {
	//:= : the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value.)
	name := "Reza"  //inferred type:string
	age := 24       //inferred type:int
	weight := 80.32 //inferred type:float
	fmt.Println(name, "is", age, "years old.")
	fmt.Println("he is", weight, "kg.")

	//declaring variables
	var student string = "asqar"
	var score float64 = 19.5
	var student_id int = 1234
	var ispassed bool = true
	fmt.Println(student, score, student_id, ispassed)

}

package main

import "fmt"


//IF and Else
func makeBabies(x string, y string, message *string) bool{
	if x == "male" && y == "female"{
		msg:="Yes, these two person can make babies."
		*message = msg
		return true
	} else if x == "female" && y == "male"{
		msg :="Yes, thse two person can make babies."
		*message = msg
		return true
	} else{
		msg :="No, these two person cannot reproduce."
		*message = msg
		return false
	}
}


//Functions that returns another functions
type mySumType func(int)

func myPrintSum(yourSum int) mySumType {
	return func(s int) {
		fmt.Println(s + yourSum)
	}
}

func giveMeOurSum(sIn int, totalSum mySumType ){
	totalSum(sIn)
}

//Functions that accepts function as arguments
type convert func(int) string

func value(x int) string {
    return fmt.Sprintf("%v", x)
}

func sayNumber(fn convert) string {
    return `"` + fn(123) + `"`
}


func main(){

	
	//using the If/Else function
	fmt.Println("\n")
	jimmy :="female"
	vic :="male"
	var mg string
	ableTo := makeBabies(jimmy,vic,&mg)
	if ableTo==true{
		fmt.Println(mg)
	}else{
		fmt.Println(mg)
	}
	fmt.Println("\n---------------------------")
	
	//Switch statement
	var fruitBasket = []string{"apple", "pear", "orange", "peach"}
	var i int = 3
	switch i {
		case 0:
			fmt.Println("The first fruit in the fruitBasket is ")
			fmt.Println(fruitBasket[i] + ".")
		case 1:
			fmt.Println("The second fruit in the fruitBasket is ")
			fmt.Println(fruitBasket[i] + ".")
		case 2:
			fmt.Println("The third fruit in the fruitBasket is ")
			fmt.Println(fruitBasket[i] + ".")
		case 3:
			fmt.Println("The forth fruit in the fruitBasket is ")
			fmt.Println(fruitBasket[i] + ".")
		default:
			fmt.Println("There is only four fruits in the basket.\n\n")
			
	}
	fmt.Println("\n---------------------------\n\n")
	
	
	var x int = 10
	giveMeOurSum(x, myPrintSum(6))
	
	fmt.Println("\n---------------------------\n\n")
	var result string
	result = value(123)
	fmt.Println(result)

	result = sayNumber(value)
	fmt.Println(result)
	
	
}
package main

import "fmt"

func main(){
	
	//for loop
	for i:=1; i<=5; i++ {
		fmt.Printf("%d Hello, I love saying 'Great!' \n", i)
	}
	
	//while loop in Go
	i:=0
	for {
		if i >= 5 {break} //break out of the lopp after 5 iteration
		fmt.Println("This is coo1 forever!")
		i++
	}
	
	//Using continue
	//Prints all the odd number from 1-100
    for j := 0; j<100 ; j++ {
           if j%2 == 0 { 
               continue 
           }
           fmt.Println("Odd:", j)  
     }
	 
	 //Using Loops over Ranges with Slices
	 names := []string {"Mike", "John", "Luke", "Chris"}
	 
	 fmt.Println("")
	 for i := range names {
	     fmt.Println(i,names[i])
	 }
	 
	 //Using loops over Ranges with Maps
	 colorLikes := map[string]int{
	     "green": 3711,
	     "blue":   2138,
	     "red": 1908,
	     "orange": 912,
	 }
	
	 fmt.Println("")
	 for i := range colorLikes {
	     fmt.Println(i,colorLikes[i])
	 }
	
}
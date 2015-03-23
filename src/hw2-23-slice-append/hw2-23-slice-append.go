package main

import (
    "fmt"
)

//deleting form a slice
type MyList struct {
    list []int
}

func (l *MyList) Pop() {
    if len(l.list) == 0 { return }
    l.list = l.list[:len(l.list)-1]
}

func (l *MyList) Add(i... int) {
    l.list = append(l.list, i...)
}

func (l *MyList) String() string {
    return fmt.Sprintf("%#v",l.list)
}


//appding to slice part
var a = []int{3,4}

func Testing(s []int) {
    s = append(s, 30)

    fmt.Println(s)
}

//Methods attached and return
type Rectangle struct {
    length, width int
}

func (r Rectangle) Area_by_value() int {
    return r.length * r.width
}

func (r *Rectangle) Area_by_reference() int {
    return r.length * r.width
}



func main() {

	//appending 30 to slice
    for i := 0; i < 1; i++ {
        a[i] = i
    }

    Testing(a)

    fmt.Println(a)
	
	
	//deleting from slices
	mylist := &MyList{[]int{1,2,3}} 
	    fmt.Println(mylist)
	    mylist.Pop()
	    fmt.Println(mylist)
	    mylist.Add([]int{4,5,6}...)
	    fmt.Println(mylist)
		
		fmt.Println("")
		
		r1 := Rectangle{4, 3}
		    fmt.Println("Area: ", r1)
		    fmt.Println("Area: ", r1.Area_by_value())
		    fmt.Println("Area: ", r1.Area_by_reference())
		    fmt.Println("Area: ", (&r1).Area_by_value())
		    fmt.Println("Area: ", (&r1).Area_by_reference())
}
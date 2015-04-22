package main

import (
	"bufio"
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"os"
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
	user := "root"
	pass := ""
	dbname := "Game_Collection"

	db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)

	err := db.Connect()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database Connected!\n")
	}

	s := "run"
	consolereader := bufio.NewReader(os.Stdin)
	for s == "run" {
		fmt.Print("Enter Querry: ")
		userInput, err := consolereader.ReadString('\n') // this will prompt the user for input
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			res, err := db.Start(userInput)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Print fields names
			i := 0
			for _, field := range res.Fields() {
				if i == 0 {
					fmt.Printf("%s", field.Name)
				}
				if i == 1 {
					fmt.Printf("%20s", field.Name)
				}
				if i == 2 {
					fmt.Printf("%18s", field.Name)
				}
				if i == 3 {
					fmt.Printf("%13s", field.Name)
				}
				if i == 4 {
					fmt.Printf("%28s", field.Name)
				}
				i++
			}
			fmt.Println()

			// Print all rows
			for {
				row, err := res.GetRow()
				if err != nil {
					panic(err)
				}
				if row == nil {
					// No more rows
					break
				}

				// Print all cols
				for _, col := range row {
					if col == nil {
						fmt.Print("<NULL>")
					} else {
						os.Stdout.Write(col.([]byte))
					}
					fmt.Print("\t\t")
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}

}

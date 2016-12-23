package main

import (
	"fmt"
	"log"
	"strconv"
	"os"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	database string = "div3element"
	max_element_number int = 103
)
var (
	err error
)

// get element name from element number via MySQL
func getElementName(n int) (name string) {
	// connect to database
	db, err := sql.Open("mysql", "root:" + os.Getenv("MYSQL_PASSWD") + "@/" + database)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	// query
	query := "select name from element where n = " + strconv.Itoa(n) + ";"
	// TODO: SQLi対策が必要？
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	
	// result
	rows.Next() 
	err = rows.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
		
	return name
}

// calculate product of slice n
func Prod(n []int) (p int) {
	p = 1
	for _, v := range n {
		p *= v
	}	
	return
}


func main() {
	// stdin
	args := os.Args
	if len(args) != 4 {
		log.Fatal("Please input 3 numbers")
	}

	cards := make([]int, 3)
	for i := 0; i < len(args) - 1; i++ {
		cards[i], err = strconv.Atoi(args[i+1])
		if err != nil {
			log.Fatal(err)
		}
	}
	
	// calc element number
	element_number := Prod(cards) % max_element_number
	
	name := getElementName(element_number)
	fmt.Println("Answer: " + strconv.Itoa(element_number) + ", " + name)
}
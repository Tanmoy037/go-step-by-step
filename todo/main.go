package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbInstanceEndpoint = "mydb.crxoqxzjvfze.ap-south-1.rds.amazonaws.com:3306"
	username           = "somnath"
	password           = "rds123456"
)

func main() {
	// Create a connection to the RDS MySQL endpoint.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/mydatabase", username, password, dbInstanceEndpoint))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a table to store the tasks.
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks (task TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	// Exec the statement to create the table.
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the tasks from the user.
	tasks := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter your tasks here: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		tasks = append(tasks, input)
	}

	// Store the tasks in the database.
	for _, task := range tasks {
		stmt, err := db.Prepare("INSERT INTO tasks (task) VALUES (?)")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(task)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Print the tasks from the database.
	rows, err := db.Query("SELECT task FROM tasks")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task string
		err := rows.Scan(&task)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(task)
	}
}

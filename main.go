package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Acc1234$$"
	dbname   = "demo"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Erroe in connecting to DB ", err)
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	rt1 := GetNewOrder()
	sqlStatement := `
	INSERT INTO retail_order (customer_name,customer_address,order_detail)
	VALUES ($1,$2,$3) RETURNING id`
	var ID int
	//_, err = conn.Exec(sqlStatement, rt1.customer_name, rt1.customer_address, rt1.order_detail)
	err = conn.QueryRow(sqlStatement, rt1.customer_name, rt1.customer_address, rt1.order_detail).Scan(&ID)

	if err != nil {
		panic(err)
	}
	fmt.Println(ID)
	defer conn.Close()
}

func GetNewOrder() Retail_Order {

	ro := Retail_Order{
		customer_name:    "Donald Modi",
		customer_address: "Delhi",
		order_detail:     "Votes and Votes",
	}
	return ro

}

type Retail_Order struct {
	customer_name    string
	customer_address string
	order_detail     string
}

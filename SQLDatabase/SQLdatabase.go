package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Item  string `json:"item"`
	Shop  string `json:"shop"`
	Price string `json:"price"`
}

func main() {

	//Open SQL conenction
	db, err := sql.Open("mysql", "root:Xuanze94@tcp(127.0.0.1:3006)/mock_itemdb") //username:password@tcp(host:port)/db
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	//Insert
	insert, err := db.Query("INSERT INTO ITEMSDB(ITEMS) VALUES ('123')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	//Query
	results, err := db.Query("SELECT ITEMID,SHOPID,PRICE FROM ITEMS")
	if err != nil {
		panic(err.Error())
	}

	//Loop all rows
	for results.Next() {
		//New variable of struct 'Item'
		var itemSelected Item
		err = results.Scan(&itemSelected.Item, &itemSelected.Shop, &itemSelected.Price)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Item: " + itemSelected.Item + " Shop: " + itemSelected.Shop + " Price: " + itemSelected.Price)
	}

}

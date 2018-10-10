package main

import (
	"fmt"
	"library"
	"stringReverse"
)

var (
	request string
	rame    string
	id      uint64
)

func main() {
	fmt.Println(stringReverse.Reverse("dlrow, olleH"))

	request = "SELECT Rame,id FROM info_bilan WHERE id=5375495"

	fmt.Println(library.Merde())
	server := sqlQuery.ConnectDb()
	defer server.Close()

	result := sqlQuery.QueryRow(request, server)
	errScan := result.Scan(&rame, &id)
	if errScan != nil {
		panic(errScan.Error())
	}
	fmt.Println("Resultat de la fonction queryRow() "+rame, id)

}

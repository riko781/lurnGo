package main

import (
	"fmt"
	"library/merde"
	"library/sql"
	"stringReverse"
)

var (
	request string
	nom     string
	prenom  string
	db      database
)

type database struct {
	user    string
	mdp     string
	name    string
	adresse string
}

func main() {
	fmt.Println(stringReverse.Reverse("dlrow, olleH"))
	fmt.Println(merde.Merde())

	request = "SELECT nom,prenom FROM user"
	db.user = "root"
	db.mdp = "maladede781"
	db.name = "test"
	db.adresse = "localhost"
	fmt.Println(db.user)

	server := sql.ConnectDb(db.user, db.mdp, db.adresse, db.name)
	defer server.Close()

	result := sql.QueryRow(request, server)
	errScan := result.Scan(&nom, &prenom)
	if errScan != nil {
		panic(errScan.Error())
	}
	fmt.Println("Resultat de la fonction queryRow() "+nom, prenom)

}

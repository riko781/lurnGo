package sql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectDb() (sqlConnect *sql.DB) {
	var server *sql.DB
	var err error
	server, err = sql.Open("mysql", "inovatic:sysiphe@tcp(192.168.1.147)/production")
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func queryRow(request string, server *sql.DB) *sql.Row {

	row := server.QueryRow(request)
	return row
}

func query(request string, server *sql.DB) (rows *sql.Rows) {
	rows, err := server.Query(request)

	if err != nil {
		panic(err.Error())
	}
	return rows
}

var (
	request string
	rame    string
	id      uint64
)

/*
func main() {
	request = "SELECT Rame,id FROM info_bilan WHERE id=5375495"

	server := connectDb()
	defer server.Close()

	result := queryRow(request, server)
	errScan := result.Scan(&rame, &id)

	if errScan != nil {
		panic(errScan.Error())
	}

	fmt.Println("Resultat de la fonction queryRow() "+rame, id)


	request = "SELECT id,rame FROM info_bilan"

	results := query(request, server)
	defer results.Close()

	for results.Next() {
		err := results.Scan(&id, &rame)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, rame)
	}

}*/

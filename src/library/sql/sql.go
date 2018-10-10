package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(user, mdp, adresse, name string) *sql.DB {
	var server *sql.DB
	var err error
	server, err = sql.Open("mysql", user+":"+mdp+"@tcp("+adresse+")/"+name)
	if err != nil {
		panic(err.Error())
	}
	return server
}

func QueryRow(request string, server *sql.DB) *sql.Row {
	row := server.QueryRow(request)
	return row
}

func Query(request string, server *sql.DB) *sql.Rows {
	rows, err := server.Query(request)
	if err != nil {
		panic(err.Error())
	}
	return rows
}

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

package Repository

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type RepositoryMysql struct{
	db *sql.DB
}

var (
	username 	string 	="usnbuus7qz634mfn"
	password 	string 	="eGPmxLkhfNLevfwyn6hW"
	host			string	="bsifbrjg2n6llr3biozg-mysql.services.clever-cloud.com"
	port			int			=3306
	database	string	="bsifbrjg2n6llr3biozg"
)

func SetupRepositoryMysql() (*RepositoryMysql, error){
	db,err := sql.Open("mysql",url())
	if err != nil{
		return nil, err
	}
	return &RepositoryMysql{db}, nil
}

func url() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",username,password,host,port,database)
}


//------------------------------------QUERY MYSQL
func(s *RepositoryMysql) Query(query string, args ...interface{}) (*sql.Rows,error){
	rows,err := s.db.Query( query, args...)
	if err != nil{
		log.Println(err)
	}
	return rows, err
}

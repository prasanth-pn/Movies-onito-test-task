package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
)


func ConnectDB(cofig config.Config)*sql.DB{
	db,err:=sql.Open("postgres",cofig.DBSource)
	if err!=nil{
		log.Fatalf("postgres database not conected  %s",err)
	}
	err=db.Ping()
	if err!=nil{
		log.Fatalf("Error in ping")
	}


	return db
}
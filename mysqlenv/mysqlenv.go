// Package mysqlenv gets a MySQL database from environment variables.
// See Get() for more details.
package mysqlenv

import (
	"bytes"
	"database/sql"

	//mysql support
	"strings"

	"github.com/codercom/env"
	_ "github.com/go-sql-driver/mysql"
)

// DSN builds a valid MySQL dsn. password is optional
func DSN(user, password, host, database string) string {
	//fmtStr := "%v:%v@(%v)/%v?parseTime=true"
	buf := new(bytes.Buffer)
	buf.WriteString(user)
	if password != "" {
		buf.WriteByte(':')
		buf.WriteString(password)
	}
	if !strings.Contains(host, ":") {
		host += ":3306"
	}
	buf.WriteString("@(")
	buf.WriteString(host)
	buf.WriteString(")/")
	buf.WriteString(database)
	buf.WriteString("?parseTime=true")
	return buf.String()
}

// Get gets databse credentials from the environment.
// It requires MYSQL_HOST, MYSQL_DB, MYSQL_USER.
// MYSQL_PASS is optional.
// Get pings the DB to verify credentials.
func Get() *sql.DB {
	var user, password, host, database string
	env.Env(env.E{
		"MYSQL_PASS": env.Optional(&password),
		"MYSQL_HOST": env.Required(&host),
		"MYSQL_DB":   env.Required(&database),
		"MYSQL_USER": env.Required(&user),
	})

	dsn := DSN(user, password, host, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

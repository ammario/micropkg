package mysqlenv

import "testing"

func TestBuildMySQLDSN(t *testing.T) {
	t.Run("NoPass", func(t *testing.T) {
		dsn := DSN("root", "", "127.0.0.1:3306", "stratex")

		if dsn != "root@(127.0.0.1:3306)/stratex?parseTime=true" {
			t.Errorf("DSN == %v ", dsn)
		}
	})

	t.Run("Pass", func(t *testing.T) {
		dsn := DSN("root", "dogsownus", "127.0.0.1:3306", "stratex")

		if dsn != "root:dogsownus@(127.0.0.1:3306)/stratex?parseTime=true" {
			t.Errorf("DSN == %v ", dsn)
		}
	})

	t.Run("No Port", func(t *testing.T) {
		dsn := DSN("root", "dogsownus", "127.0.0.1", "stratex")

		if dsn != "root:dogsownus@(127.0.0.1:3306)/stratex?parseTime=true" {
			t.Errorf("DSN == %v ", dsn)
		}
	})
}

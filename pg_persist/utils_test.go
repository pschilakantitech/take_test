package pg_persist

import (
	"fmt"
)

func init() {
	cfg := Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "",
		Database: "mycoinscap",
	}
	if err := ConnectToPGDB(cfg); err != nil {
		fmt.Println("fail to get DB connection")
		fmt.Println(err)
	}
}

//func TestGetUserCoinsAndCounts(t *testing.T) {
//	_, err := GetUserCoinsAndCounts("novalue")
//	if err.Error() != ErrNoRecords.Error() {
//		t.Fatal("execting no records error, got ", err)
//	}
//	arr, err := GetUserCoinsAndCounts("test")
//	if err != nil {
//		t.Fatal("execting no error, got ", err)
//	}
//	if len(arr) == 0 {
//		t.Fatal("execting few records, no records got ", err)
//	}
//}

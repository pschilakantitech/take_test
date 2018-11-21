package pg_persist

import (
	"fmt"
	"testing"
)

func init() {
	cfg := Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "praveensc48",
		Database: "take_test_db",
	}
	if err := ConnectToPGDB(cfg); err != nil {
		fmt.Println("fail to get DB connection")
		fmt.Println(err)
	}
}

func TestGetNextQuestion(t *testing.T) {

	d, err := isScoreCalculated("633c78b2d5dfecf3be118bf111d11930")
	t.Fatal("here we got the error", err, d)

}

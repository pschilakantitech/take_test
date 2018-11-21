package pg_persist

import (
	"fmt"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type TestDetailsResp struct {
	TestQuestions int    `json:"test_questions"`
	TestTime      string `json:"test_time"`
}

type StudDetailsReq struct {
	FistName string `json:"fist_name"`
	LastName string `json:"last_name"`
	MobileNo string `json:"mobile_no"`
	Email    string `json:"email"`
}

type StudResultResp struct {
	FistName  string `json:"fist_name"`
	LastName  string `json:"last_name"`
	MobileNo  string `json:"mobile_no"`
	Email     string `json:"email"`
	Score     string `json:"score"`
	TakenOn   string `json:"taken_on"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
	Duration  string `json:"duration"`
	RefNo     string `json:"ref_no"`
}

type QustionResp struct {
	SerialNo int    `json:"serial_no"`
	QID      int    `json:"question_id"`
	Question string `json:"question"`
	OptionA  string `json:"option_a"`
	OptionB  string `json:"option_b"`
	OptionC  string `json:"option_c"`
	OptionD  string `json:"option_d"`
	RefNo    string `json:"ref_no"`
}

// basic validations only
func (l *StudDetailsReq) Validate() error {
	if strings.TrimSpace(l.FistName) == "" {
		return fmt.Errorf("first name is empty")
	}
	if strings.TrimSpace(l.LastName) == "" {
		return fmt.Errorf("last name is empty")
	}
	if strings.TrimSpace(l.MobileNo) == "" {
		return fmt.Errorf("mobile no is empty")
	}
	if strings.TrimSpace(l.Email) == "" {
		return fmt.Errorf("emil id is empty")
	}
	if !emailRegex.MatchString(l.Email) {
		return fmt.Errorf("invalid email ID")
	}

	return nil
}

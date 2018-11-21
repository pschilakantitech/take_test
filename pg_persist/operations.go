package pg_persist

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrNoRecords      = errors.New("no_records")
	ErrNoRowsAffected = errors.New("no_rows_affected")
	ErrAlreadyUpdated = errors.New("answare_already_updated")
)

const (
	selQuestionsCount = "SELECT count(id) FROM public.questions_ans"
	insStudent        = `INSERT INTO public.test_details(ref_no, stud_first_name, stud_last_name, stud_mobile_no, stud_email, taken_on, start_at) VALUES ($1, $2, $3, $4, $5, current_timestamp, current_timestamp)`
	selQuestion       = `SELECT serial_no ,id, question, option_a, option_b, option_c, option_d FROM public.questions_ans where serial_no=$1`
	insAnswer         = "INSERT INTO public.test_answers(question_id, test_id, option_selected) SELECT $1,id,$2 FROM public.test_details where ref_no=$3"
	sleNextSlNo       = `SELECT serial_no FROM public.questions_ans where id=$1`

	sleGetResult = `select count(1) from public.test_answers as test 
	                      inner join public.test_details as details on details.id = test.test_id
                          inner join public.questions_ans ans on ans.id = test.question_id and lower(ans.answer) = lower(test.option_selected)
                          where details.ref_no=$1`

	updScoreForRefNo = `UPDATE public.test_details SET score=$1, ended_at=current_timestamp WHERE ref_no=$2`

	sleStuWithResult = `SELECT ref_no, stud_first_name, stud_last_name, stud_mobile_no, stud_email,score,
                          TO_CHAR(taken_on, 'dd-Mon-YYYY') taken_on,to_char(start_at, 'dd-Mon-YYYY HH12:MI:SS') as start_at,
                          to_char(ended_at, 'dd-Mon-YYYY HH12:MI:SS') as ended_at, DATE_PART('minute', ended_at - start_at) || ' min'  as test_duration
                          FROM public.test_details where ref_no=$1`

	sleAllStuResults = `SELECT ref_no, stud_first_name, stud_last_name, stud_mobile_no, stud_email,score,
                          TO_CHAR(taken_on, 'dd-Mon-YYYY') taken_on,to_char(start_at, 'dd-Mon-YYYY HH12:MI:SS') as start_at,
                          to_char(ended_at, 'dd-Mon-YYYY HH12:MI:SS') as ended_at, DATE_PART('minute', ended_at - start_at) || ' min'  as test_duration
                          FROM public.test_details where score != ''`

	sleCheckScore = `SELECT score FROM public.test_details where ref_no =$1`
)

func GetQuestionsCount() (id int, err error) {
	queryStmt, _ := Db.Prepare(selQuestionsCount)
	err = queryStmt.QueryRow().Scan(&id)
	if err == sql.ErrNoRows {
		return 0, ErrNoRecords

	} else if err != nil {
		return 0, errors.Wrap(err, "failed to get questions counts")
	}

	return
}

func AddStudentDetails(refNo string, req StudDetailsReq) (err error) {
	res, err := Db.Exec(insStudent, refNo, req.FistName, req.LastName, req.MobileNo, req.Email)
	if err != nil {
		return errors.Wrap(err, "insert student failed")
	}
	if v, _ := res.RowsAffected(); v <= 0 {
		return ErrNoRowsAffected
	}

	return
}

func GetNextQuestion(currentSlNo int) (*QustionResp, error) {
	prepare, err := Db.Prepare(selQuestion)
	rows, err := prepare.Query(currentSlNo + 1)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get next question")
	}
	if rows.Next() {
		r := QustionResp{}
		err = rows.Scan(&r.SerialNo, &r.QID, &r.Question, &r.OptionA, &r.OptionB, &r.OptionC, &r.OptionD)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan next question details")
		}
		return &r, nil
	}

	return nil, ErrNoRecords
}

func UpdateAnswer(refNo, questionID, ans string) error {
	res, err := Db.Exec(insAnswer, questionID, ans, refNo)
	if err != nil { // Note: Not right way; think more
		if strings.Contains(err.Error(), "unique_com_qid_test_id") {
			return ErrAlreadyUpdated
		}
		return err
	}
	if v, _ := res.RowsAffected(); v <= 0 {
		return ErrNoRowsAffected
	}

	return nil
}

func GetSlNo(questionID string) (id int, err error) {
	prepare, err := Db.Prepare(sleNextSlNo)
	rows, err := prepare.Query(questionID)
	defer rows.Close()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get sl no")
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get sl no")
		}
		return
	}

	return 0, ErrNoRecords
}

func GetResultByRefNo(refNo string) (result *StudResultResp, err error) {
	isCalculated, err := isScoreCalculated(refNo)
	if err != nil {
		return nil, err
	}

	if isCalculated {
		return getStudWithScore(refNo)
	}

	return calculateResultByRefNo(refNo)
}

func calculateResultByRefNo(refNo string) (*StudResultResp, error) {
	prepare, err := Db.Prepare(sleGetResult)
	rows, err := prepare.Query(refNo)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get result")
	}
	if rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get result")
		}
		if err = updateScore(refNo, id); err != nil {
			return nil, err
		}
		return getStudWithScore(refNo)
	}

	return nil, ErrNoRecords
}

func updateScore(refNo string, score int) (err error) {
	outOf, err := GetQuestionsCount()
	if err != nil {
		return
	}
	res, err := Db.Exec(updScoreForRefNo, fmt.Sprintf("%d/%d", score, outOf), refNo)
	if err != nil {
		return errors.Wrap(err, "insert answare failed")
	}
	if v, _ := res.RowsAffected(); v <= 0 {
		return ErrNoRowsAffected
	}

	return nil
}

func getStudWithScore(refNo string) (*StudResultResp, error) {
	prepare, err := Db.Prepare(sleStuWithResult)
	rows, err := prepare.Query(refNo)
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get student results")
	}
	if rows.Next() {
		r := StudResultResp{}
		err = rows.Scan(&r.RefNo, &r.FistName, &r.LastName, &r.MobileNo, &r.Email, &r.Score, &r.TakenOn, &r.StartedAt, &r.EndedAt, &r.Duration)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get student results")
		}
		return &r, nil
	}

	return nil, ErrNoRecords
}

func GetAllResult() (results []*StudResultResp, err error) {
	prepare, err := Db.Prepare(sleAllStuResults)
	rows, err := prepare.Query()
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get students results")
	}
	for rows.Next() {
		r := StudResultResp{}
		err = rows.Scan(&r.RefNo, &r.FistName, &r.LastName, &r.MobileNo, &r.Email, &r.Score, &r.TakenOn, &r.StartedAt, &r.EndedAt, &r.Duration)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get students results")
		}
		fmt.Println("upending hbskdcb")
		fmt.Println(r)
		results = append(results, &r)
		fmt.Println(len(results))
	}

	return

}

func isScoreCalculated(refNo string) (check bool, err error) {
	var score sql.NullString
	prepare, err := Db.Prepare(sleCheckScore)
	rows, err := prepare.Query(refNo)
	defer rows.Close()
	if err != nil {
		return false, errors.Wrap(err, "failed to check the score")
	}
	if rows.Next() {
		err = rows.Scan(&score)
		if err != nil {
			return false, errors.Wrap(err, "failed to check the score")
		}
		return score.Valid, nil
	}

	return false, ErrNoRecords
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pschilakantitech/avitar/log"
	"github.com/pschilakantitech/avitar/utils"
	"github.com/pschilakantitech/take_test/env"
	"github.com/pschilakantitech/take_test/pg_persist"
)

func assignHandlers() {
	e.GET("/testdetails", testDetails)
	e.POST("/enroll", enroll)
	e.GET("/starttest", starTest)
	e.GET("/updateans", updateAnsAndGetNextQuestion)
	e.GET("/getresult", getResultForRefNo)
	e.GET("/admin", getAllResults)

	e.Static("static", "web_ui/assets")
	e.File("/", "web_ui/home.html")

	e.GET("/enrollView", enrollPage)
	e.GET("/firstquestion", firstquestion)
	e.GET("/finishtest", finishtest)

	return

}

// UI pages

func enrollPage(c echo.Context) error {
	return c.File("web_ui/enroll.html")
}
func firstquestion(c echo.Context) error {
	return c.File("web_ui/questions.html")
}
func finishtest(c echo.Context) error {
	return c.File("web_ui/finishtest.html")
}

func testDetails(c echo.Context) error {
	count, err := pg_persist.GetQuestionsCount()
	if err != nil {
		if err.Error() == pg_persist.ErrNoRecords.Error() {
			return apiErrNoRecords().render(c)
		}
		log.Error("Failed to get test details", err)
		return apiErrGetTestDetailsFailed().render(c)
	}

	resp := pg_persist.TestDetailsResp{TestQuestions: count, TestTime: env.TestTime.String()}
	return c.JSONPretty(http.StatusOK, &resp, jsonIndent)

}

func enroll(c echo.Context) error {
	defer c.Request().Body.Close()
	reqPara := pg_persist.StudDetailsReq{}
	err := json.NewDecoder(c.Request().Body).Decode(&reqPara)
	if err != nil {
		log.Error("Failed to parse test details ", err)
		return apiErrStartTestFailed().render(c)
	}
	err = reqPara.Validate()
	if err != nil {
		log.Error("student details validation failed ", err)
		return apiErrStudDetailsValidationFail(err).render(c)
	}

	refNo := utils.RandomMD5Hash()
	err = pg_persist.AddStudentDetails(refNo, reqPara)
	if err != nil {
		log.Error("Failed to insert student details", err)
		return apiErrStudDetailsInsertionFail(err).render(c)
	}

	return c.String(http.StatusOK, refNo)

}

func starTest(c echo.Context) error {
	refNo := c.QueryParam("ref_no")
	if refNo == "" {
		log.Error("ref no is empty")
		return apiErrRefNoEmpty().render(c)
	}

	questionResp, err := pg_persist.GetNextQuestion(0)
	if err == pg_persist.ErrNoRecords {
		return apiErrNoMoreQuestions().render(c)
	}
	if err != nil {
		log.Error("Failed to get question", err)
		return apiErrGetQuestionFailed().render(c)
	}

	questionResp.RefNo = refNo
	return c.JSONPretty(http.StatusOK, &questionResp, jsonIndent)

}

func updateAnsAndGetNextQuestion(c echo.Context) error {
	refNo, questionID, ans := c.QueryParam("ref_no"), c.QueryParam("q_id"), c.QueryParam("ans")
	if refNo == "" || questionID == "" || ans == "" {
		log.Error("One of ref no, question id or ans is empty")
		return apiErrEmptyAnsParaValues().render(c)
	}

	err := pg_persist.UpdateAnswer(refNo, questionID, ans)
	if err != nil {
		if err.Error() == pg_persist.ErrAlreadyUpdated.Error() {
			log.Error("answare already updated", err)
			return apiErrAnsAlreadyUpdated().render(c)
		}
		log.Error("failed to update the ans", err)
		return apiErrAnsUpdateFailed().render(c)
	}

	slNo, err := pg_persist.GetSlNo(questionID)
	if err != nil {
		log.Error("failed to get sl no", err)
		if err.Error() == pg_persist.ErrNoRecords.Error() {
			return apiErrNoRecords().render(c)
		}
		return apiErrGetSlNoFailed().render(c)
	}

	questionResp, err := pg_persist.GetNextQuestion(slNo)
	if err != nil {
		log.Error("Failed to get question", err)
		if err.Error() == pg_persist.ErrNoRecords.Error() {
			return apiErrNoRecords().render(c)
		}
		return apiErrGetQuestionFailed().render(c)
	}

	questionResp.RefNo = refNo
	return c.JSONPretty(http.StatusOK, &questionResp, jsonIndent)

}
func getResultForRefNo(c echo.Context) error {
	refNo := c.QueryParam("ref_no")
	if refNo == "" {
		log.Error("ref no is empty")
		return apiErrRefNoEmpty().render(c)
	}

	result, err := pg_persist.GetResultByRefNo(refNo)
	if err != nil {
		log.Info("get result failed", err)
		if err.Error() == pg_persist.ErrNoRecords.Error() {
			return apiErrNoRecords().render(c)
		}
		return apiErrGetResultFailed().render(c)
	}

	return c.JSONPretty(http.StatusOK, &result, jsonIndent)
}
func getAllResults(c echo.Context) error {
	result, err := pg_persist.GetAllResult()
	if err != nil {
		log.Info("failed to get all results", err)
		if err.Error() == pg_persist.ErrNoRecords.Error() {
			return apiErrNoRecords().render(c)
		}
		return apiErrGetAllResultsFailed().render(c)
	}
	return c.JSONPretty(http.StatusOK, &result, jsonIndent)
}

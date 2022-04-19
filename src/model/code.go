package model

import (
	"fmt"
	"github.com/FlyInThesky10/antiNCP-backend/controller/param"
	"time"
)

func AddSubmission(sbm param.ReqUserSubmitCode, id string) error {
	_, err := Db.Exec(
		"INSERT INTO submission(id, vaccine_stage, health, travel_history, health_code, travel_code, accepted, submit_time)values(?, ?, ?, ?, ?, ?, ?, ?)",
		id,
		sbm.VaccineStage,
		sbm.Health,
		sbm.TravelHistory,
		sbm.HealthCode,
		sbm.TravelCode,
		0,
		time.Now().Unix(),
	)
	if err != nil {
		return err
	}
	return nil
}
func GetSubmissionById(id string, startIndex int, length int) ([]param.Submission, error) {
	submissions := make([]param.Submission, 0)
	err := Db.Select(&submissions, fmt.Sprintf("SELECT id,i,vaccine_stage,health,travel_history,health_code,travel_code,submit_time,accepted FROM submission WHERE id=%s LIMIT %d,%d", id, startIndex, length))
	if err != nil {
		return []param.Submission{}, err
	}
	return submissions, nil
}
func GetSubmissions(startIndex int, length int) ([]param.Submission, error) {
	submissions := make([]param.Submission, 0)
	err := Db.Select(&submissions, fmt.Sprintf("SELECT id,i,vaccine_stage,health,travel_history,health_code,travel_code,submit_time,accepted FROM submission LIMIT %d,%d", startIndex, length))
	if err != nil {
		return []param.Submission{}, err
	}
	return submissions, nil
}
func GetStatus(id string) (param.Code, error) {
	code := param.Code{}
	err := Db.Get(&code, fmt.Sprintf("SELECT status,verify_id,verify_time FROM code WHERE id=%s", id))
	if err != nil {
		return param.Code{}, err
	}
	return code, nil
}
func UpdateStatus(index int, verifyId string, status int) error {
	id := param.Id{}
	err := Db.Get(&id, fmt.Sprintf("SELECT id FROM submission WHERE i=%d", index))
	if err != nil {
		return err
	}

	tx := Db.MustBegin()
	tx.MustExec(fmt.Sprintf("UPDATE code SET status=%d, verify_id=%s, verify_time=%d WHERE id=%s", status, verifyId, time.Now().Unix(), id.Id))
	tx.MustExec(fmt.Sprintf("UPDATE submission SET accepted=%d WHERE i=%d", status, index))
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

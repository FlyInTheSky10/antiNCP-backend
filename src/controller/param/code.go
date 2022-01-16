package param

type Submission struct {
	Id            string `db:"id"`
	I             int    `db:"i"`
	VaccineStage  int    `db:"vaccine_stage"`
	Health        int    `db:"health"`
	TravelHistory int    `db:"travel_history"`
	HealthCode    int    `db:"health_code"`
	TravelCode    int    `db:"travel_code"`
	SubmitTime    int    `db:"submit_time"`
	Accepted      int    `db:"accepted"`
}
type Code struct {
	Status     int `db:"status"`
	VerifyId   int `db:"verify_id"`
	VerifyTime int `db:"verify_time"`
}
type Id struct {
	Id string `db:"id"`
}
type ReqUserSubmitCode struct {
	VaccineStage  int `query:"vaccine_stage"`
	Health        int `query:"health"`
	TravelHistory int `query:"travel_history"`
	HealthCode    int `query:"health_code"`
	TravelCode    int `query:"travel_code"`
}
type ReqAdminVerifyCode struct {
	I      int `query:"i"`
	Status int `query:"status"`
}

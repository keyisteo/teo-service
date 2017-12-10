package models

import "time"

type Report struct {
	Id           int     `form:="-"`
	Timestamp    Time    `form:="-"`
	IdReporter   int     `form:="id_reporter"`
	HandleReport boolean `form:="-"`
	Category     int     `form:="category"`
	Detail       string  `form:="detail"`
	LinkPhoto    string  `form:"client,text,client:"`
}

func getReport(id int) Report {
	//Database logic by comparing string

	//return report file
	return
}

func insertReport(r Report) {
	//Manipulate Database
	r.Timestamp = time.Now()
}

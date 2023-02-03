package main

import (
	"github.com/willfaught/gockle"
)

type ReputationScore struct {
	Id             string
	Value1, Value2 uint16
	Timestamp      int64
}

func ExecuteQuery(session gockle.Session, providedScore *ReputationScore) {
	session.Query(`INSERT INTO table1 (id, value1, value2, timestamp) VALUES (?, ?, ?, ?)`,
		providedScore.Id, providedScore.Value1, providedScore.Value2, providedScore.Timestamp).Exec()
}

func main() {

}

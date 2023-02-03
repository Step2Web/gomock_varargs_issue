package main

import (
	"time"

	"github.com/willfaught/gockle"
)

type SomeStruct struct {
	Id             string
	Value1, Value2 uint16
	Timestamp      int64
}

func ExecuteQuery(session gockle.Session, someStruct *SomeStruct) {
	session.Query(`INSERT INTO table1 (id, value1, value2, timestamp) VALUES (?, ?, ?, ?)`,
		someStruct.Id, someStruct.Value1, someStruct.Value2, time.UnixMilli(someStruct.Timestamp)).Exec()
}

func main() {

}

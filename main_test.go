package main

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
)

func TestExecuteQuery(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name             string
		configureExpects func(session *MockSession, query *MockQuery, now time.Time)
		score            *SomeStruct
	}{
		{
			// Succeeds
			name: "Use Any for all parameters",
			configureExpects: func(sessionMock *MockSession, queryMock *MockQuery, now time.Time) {
				sessionMock.EXPECT().Query(`INSERT INTO table1 (id, value1, value2, timestamp) VALUES (?, ?, ?, ?)`, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(queryMock).Times(1)
				queryMock.EXPECT().Exec().Times(1)

			},
			score: &SomeStruct{
				Id:        "t2_123",
				Value1:    555,
				Value2:    0,
				Timestamp: now.UnixMilli(),
			},
		},
		{
			// Succeeds
			name: "Use Any for all parameters, except Id",
			configureExpects: func(sessionMock *MockSession, queryMock *MockQuery, now time.Time) {
				sessionMock.EXPECT().Query(`INSERT INTO table1 (id, value1, value2, timestamp) VALUES (?, ?, ?, ?)`, "t2_123", gomock.Any(), gomock.Any(), gomock.Any()).Return(queryMock).Times(1)
				queryMock.EXPECT().Exec().Times(1)

			},
			score: &SomeStruct{
				Id:        "t2_123",
				Value1:    555,
				Value2:    0,
				Timestamp: now.UnixMilli(),
			},
		},
		{
			// Fails - although expected to succeed
			name: "Exact Match",
			configureExpects: func(sessionMock *MockSession, queryMock *MockQuery, now time.Time) {
				sessionMock.EXPECT().Query(`INSERT INTO table1 (id, value1, value2, timestamp) VALUES (?, ?, ?, ?)`, "t2_123", 555, 0, now).Return(queryMock).Times(1)
				queryMock.EXPECT().Exec().Times(1)

			},
			score: &SomeStruct{
				Id:        "t2_123",
				Value1:    555,
				Value2:    0,
				Timestamp: now.UnixMilli(),
			},
		},
		{
			// Fails - But I would not expect this to succeed
			name: "Varargs as interface{} list",
			configureExpects: func(sessionMock *MockSession, queryMock *MockQuery, now time.Time) {
				sessionMock.EXPECT().Query(`INSERT INTO table1 (id, value1, value2, timestamp) VALUES (?, ?, ?, ?)`, []interface{}{"t2_123", 555, 0, now}).Return(queryMock).Times(1)
				queryMock.EXPECT().Exec().Times(1)

			},
			score: &SomeStruct{
				Id:        "t2_123",
				Value1:    555,
				Value2:    0,
				Timestamp: now.UnixMilli(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			sessionMock := NewMockSession(ctrl)
			queryMock := NewMockQuery(ctrl)

			now := time.Now()

			test.configureExpects(sessionMock, queryMock, now)

			ExecuteQuery(sessionMock, test.score)

			ctrl.Finish()
		})
	}

}

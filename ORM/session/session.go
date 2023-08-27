package geeSession

import (
	"database/sql"
	"strings"
)

type Session struct {
	db *sql.DB
	sql strings.Builder
	sqlValues []interface{}
}

func NewSession(db *sql.DB) *Session {
	return &Session{db : db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlValues = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

/*
 * 处理原始数据，将当前session的sql和values添加到session中
*/
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql + " ")
	s.sqlValues = append(s.sqlValues, values...)
	return s
}

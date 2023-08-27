package geeSession

import (
	"database/sql"
	"gee/ORM/log"
)

func (s *Session) Exec() (result sql.Result, err error){
	defer s.Clear()

	geeLog.Info(s.sql.String(), s.sqlValues)

	if result, err = s.DB().Exec(s.sql.String(), s.sqlValues...); err != nil{
		geeLog.Error(err)
	}

	return
}

/*
 * 单行查询
*/
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	geeLog.Info(s.sql.String(), s.sqlValues)
	return s.DB().QueryRow(s.sql.String(), s.sqlValues...)
}

/*
 * 多行查询
*/
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	geeLog.Info(s.sql.String(), s.sqlValues)
	
	if rows, err = s.DB().Query(s.sql.String(), s.sqlValues...); err != nil {
		geeLog.Error(err)
	}

	return

}


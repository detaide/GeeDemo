package geeorm


import (
	"database/sql"
	"gee/ORM/log"
	"gee/ORM/session"
	"strings"
)


/*
 * 顶层封装，提供给外部使用
 */

type Engine struct {
	db *sql.DB
}



func NewEngine(driver , source string) (engine *Engine, err error) {

	if driver == "" || source == "" {
		geeLog.Error("Function NewEngine parameters are missing, [driver] | [source] is empty ")
	}

	parts := strings.Split(source, "/")
	source = parts[len(parts) - 1]
	db, err := sql.Open(driver, source)

	if err != nil {
		geeLog.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		geeLog.Error(err)
		return
	}

	engine = &Engine{db : db}
	geeLog.Info("Connect SqlDataBase - ", driver, "-", source)

	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		geeLog.Error("Failed to close database")
		return
	}
	geeLog.Info("close database success")

}

func (engine *Engine) NewSession() *geeSession.Session {
	return geeSession.NewSession(engine.db)
}
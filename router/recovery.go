package router

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"net/http"
)

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr  // skip first 3 caller
	n := runtime.Callers(3, pcs[:])
	fmt.Println(n)

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				ctx.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		ctx.Next()
	}
}
package utils

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"net/http"
	"runtime/debug"
)

func HandlePanic(ctx context.Context, w http.ResponseWriter) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			log.Errorf(ctx, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else if errStr, ok := r.(string); ok {
			log.Errorf(ctx, errStr)
			http.Error(w, errStr, http.StatusInternalServerError)
		}
		log.Errorf(ctx, "%s", debug.Stack())
	}
}

func PanicIfError(err error, args ...interface{}) {
	if err != nil {
		additionalErrorMsg := ""
		for _, arg := range args {
			additionalErrorMsg += arg.(string) + "|"
		}
		panic(additionalErrorMsg + err.Error())
	}
}

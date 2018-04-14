// Auxiliary functions for HTTP handling
package main

import (
	"fmt"
//	"github.com/op/go-logging"
	"net/http"
	"path/filepath"
	"runtime"
)

// checkErrHTTP returns an error via HTTP and also logs the error.
func checkErrHTTP(w http.ResponseWriter, httpStatus int, errorMessage string, err error) {
	if err != nil {
		http.Error(w, fmt.Sprintf(errorMessage, err), httpStatus)
		pc, file, line, ok := runtime.Caller(1)
		Log.Error("(", http.StatusText(httpStatus), ") ", filepath.Base(file), ":", line, ":", pc, ok, " - error:", errorMessage, err)	
	}
}

// checkErrPanicHTTP returns an error via HTTP and logs the error with a panic.
func checkErrPanicHTTP(w http.ResponseWriter, httpStatus int, errorMessage string, err error) {
	if err != nil {
		http.Error(w, fmt.Sprintf(errorMessage, err), httpStatus)
		pc, file, line, ok := runtime.Caller(1)
		Log.Panic("(", http.StatusText(httpStatus), ") ", filepath.Base(file), ":", line, ":", pc, ok, " - panic:", errorMessage, err)
	}
}

// logErrHTTP assumes that the error message was already composed and writes it to HTTP and logs it.
//  this is mostly to avoid code duplication and make sure that all entries are written similarly 
func logErrHTTP(w http.ResponseWriter, httpStatus int, errorMessage string) {
	http.Error(w, errorMessage, httpStatus)
	Log.Error("(" + http.StatusText(httpStatus) + ") " + errorMessage)
}

// funcName is @Sonia's solution to get the name of the function that Go is currently running.
//  This will be extensively used to deal with figuring out where in the code the errors are!
//  Source: https://stackoverflow.com/a/10743805/1035977 (20170708)
func funcName() string {
    pc, _, _, _ := runtime.Caller(1)
    return runtime.FuncForPC(pc).Name()
}
package main

/*
#include <logging.h>
*/
import "C"
import (
	"fmt"
	"strings"

	"github.com/imroc/req/v3"
)

//export logToConsole
func logToConsole(message *C.char) {
	logToConsoleGo(C.GoString(message))
}

func logToConsoleGo(message string) {
	message_changed := strings.Replace(message, " ", "_", -1)
	req.MustGet(fmt.Sprintf("%s/%s", DUMMYHTTP_URL, message_changed))
}

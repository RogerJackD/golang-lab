package logs

import "log"

func Info(message string) {
	log.Println("[INFO]", message)
}

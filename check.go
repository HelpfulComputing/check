package check

import (
	"errors"
	"log"
)

func Panic(err error, message string) {
	if err != nil {
		log.Panic(message+": ", err.Error())
	}
}

func Print(err error, message string) {
	if err != nil {
		log.Println(message+": ", err.Error())
	}
}

func Allow(err error, allowed []error) error {
	if err == nil {
		return nil
	}
	for _, a := range allowed {
		if errors.Is(err, a) {
			return nil
		}
	}
	return err
}

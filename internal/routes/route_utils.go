package routes

import "log"

func logErr(n int, err error) bool {
	isErr := err != nil

	if isErr {
		log.Printf("Error: %v", err)
	}

	return isErr
}
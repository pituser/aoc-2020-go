package util

import "log"

// CheckError halts with an error message if the given error is not nil.
func CheckError(err error, msg string) {
	if err != nil {
		if msg == "" {
			msg = "unexpected error"
		}
		log.Fatalf("%v: %v", msg, err)
	}
}

func ReduceInt(list []int, init int, reducer func(int, int) int) int {
	res := init
	for _, i := range list {
		res = reducer(res, i)
	}
	return res
}

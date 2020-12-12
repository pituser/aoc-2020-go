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

func Error(msg string) {
	if msg == "" {
		msg = "unexpected error"
	}
	log.Fatalf("%v", msg)

}

func ReduceInt(list []int, init int, reducer func(int, int) int) int {
	res := init
	for _, i := range list {
		res = reducer(res, i)
	}
	return res
}

// MaxInt returns the maximun in the given slice.
func MaxInt(list []int) (max int) {
	if len(list) == 0 {
		panic("empty list")
	}
	for i, v := range list {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

// MinInt returns the minimun in the given slice.
func MinInt(list []int) (min int) {
	if len(list) == 0 {
		panic("empty list")
	}
	for i, v := range list {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

// AbsInt returns the absolute value of an integer
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

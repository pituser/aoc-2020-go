package util

func ReduceInt(list []int, init int, reducer func(int, int) int) int {
	res := init
	for _, i := range list {
		res = reducer(res, i)
	}
	return res
}

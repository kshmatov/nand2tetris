package counter

import "strconv"

var labelCnt uint64

func Get() string {
	r := strconv.FormatUint(labelCnt, 10)
	labelCnt++
	return r
}

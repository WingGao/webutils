package wtime

import "time"

func NowPtr() *time.Time {
	n := time.Now()
	return &n
}

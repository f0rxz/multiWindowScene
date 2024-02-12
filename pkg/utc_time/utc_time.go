package utctime

import "time"

func Get() time.Time {
	return time.Now().UTC()
}

func Unix() int64 {
	return Get().Unix()
}

func UnixMilli() int64 {
	return Get().UnixMilli()
}

func UnixNano() int64 {
	return Get().UnixNano()
}

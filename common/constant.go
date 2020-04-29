package common

import "time"

const DEFAULT_TIME_LAYOUT = "2006-01-02 15:04:05"
var DEFAULT_TIME, _ = time.ParseInLocation(DEFAULT_TIME_LAYOUT, "1900-01-01 00:00:00", time.Local)
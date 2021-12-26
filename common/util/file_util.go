package util

import "github.com/gabriel-vasile/mimetype"

func GetMimeByStream(bytes []byte) string {
	return mimetype.Detect(bytes).String()
}

package api

import (
	_ "picasso/api/command"
	_ "picasso/api/file"
	_ "picasso/api/op"
	_ "picasso/api/upload"
	_ "picasso/api/user"
) // make sure this is the last line, or code gen will fail !!!

package api

import (
	_ "picasso/api/command"
	_ "picasso/api/media"
	_ "picasso/api/op"
	_ "picasso/api/space"
	_ "picasso/api/user"
) // make sure this is the last line, or code gen will fail !!!
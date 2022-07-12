package global

import (
	"log"
	"os"
)

var (
	Passphrase, Filepath string

	PATH_SEPERATOR = string(os.PathSeparator)
	Errorlog       *log.Logger
	Infolog        *log.Logger
)

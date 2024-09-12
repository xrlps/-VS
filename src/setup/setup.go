package setup

import (
	"log"
	"os"
	"time"
)

type Colour struct {
	Reset  string
	Purple string
	Green  string
	Red    string
}

func CreateColours() Colour {
	c := Colour{"\033[0m", "\033[35m", "\033[32m", "\033[31m"}
	return c
}

func CreateLogo() string {
	c := CreateColours()
	ascii := c.Purple + `

	██▒   █▓ █    ██  ██▓     ███▄    █   ██████  ▄████▄   ▄▄▄       ███▄    █ 
	▓██░   █▒ ██  ▓██▒▓██▒     ██ ▀█   █ ▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █ 
	 ▓██  █▒░▓██  ▒██░▒██░    ▓██  ▀█ ██▒░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒
	  ▒██ █░░▓▓█  ░██░▒██░    ▓██▒  ▐▌██▒  ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒
	   ▒▀█░  ▒▒█████▓ ░██████▒▒██░   ▓██░▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░
	   ░ ▐░  ░▒▓▒ ▒ ▒ ░ ▒░▓  ░░ ▒░   ▒ ▒ ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
	   ░ ░░  ░░▒░ ░ ░ ░ ░ ▒  ░░ ░░   ░ ▒░░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░░ ░░   ░ ▒░
		 ░░   ░░░ ░ ░   ░ ░      ░   ░ ░ ░  ░  ░  ░          ░   ▒      ░   ░ ░ 
		  ░     ░         ░  ░         ░       ░  ░ ░            ░  ░         ░ 
		 ░                                        ░                             	
	` + c.Reset

	return ascii
}

/*
* Creates a logger to print formatted success/error messages
 */
func CreateStdoutLog() (*log.Logger, *log.Logger, *log.Logger) {
	c := CreateColours()

	successLog := log.New(os.Stdout, time.Now().Format("15:04:04")+c.Green+" [success] "+c.Reset, 0)
	errorLog := log.New(os.Stdout, time.Now().Format("15:04:04")+c.Red+" [error] "+c.Reset, 0)
	logger := log.New(os.Stdout, "", log.Ltime)
	return successLog, logger, errorLog
}

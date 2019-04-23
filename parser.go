

import (
	"log"
	"os/exec"
)

func StartItunes() bool {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return true
}

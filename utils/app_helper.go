//Package utils contains polka specific utility functions.
package utils

import (
	//"log"
	//"encoding/json"
	"fmt"
	"os"
)

func AppErrorMessage() error {
	currentDir, _ := os.Getwd()
	return fmt.Errorf(`
Hi there!  You likely wanted to execute this command in a polka project directory.
For example, if you ran:

$ cd ~/code
$ polka new todo

then ~/code/todo is your polka project dir.  It'll have a polka/app polka/config,
and the rest of the polka generated files.

This time you ran polka in %v
`, currentDir)
}

func AreWeInProjectDir() (string, error) {
	//TODO push through this refactor
	return AreWeInAppRootDir()
}

//AreWeInAppRootDir returns nil if we're in the app root directory; otherwise it retuns the error.
func AreWeInAppRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	//we are in the app root dir if we have both a ./app and
	//a ./config in the current working dir
	checkDirs := [...]string{
		fmt.Sprintf("%v/app", currentDir),
		fmt.Sprintf("%v/config", currentDir), //expand if needed
	}
	for _, checkDir := range checkDirs {
		if _, err := os.Stat(checkDir); err != nil {
			return "", AppErrorMessage()
		}
	}
	//if we made it here, all the checkDirs exist, which means we should be good
	return currentDir, nil
}

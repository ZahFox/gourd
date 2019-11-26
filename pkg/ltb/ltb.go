package ltb

import (
	"fmt"
	"os"

	"github.com/zahfox/gourd/pkg/utils"
)

// EnsureInstalled makes sure that the linx-toolbox is properly installed
func EnsureInstalled() error {
	var err error = nil
	// TODO: Check for different platforms and adjust this accordingly
	path := "/opt/linux-toolbox"
	_, ferr := utils.Exists(path)

	switch ferr {
	case utils.FEINSFPERM:
		{
			err = fmt.Errorf("needs permission to access the linux-toolbox at %s", path)
			break
		}
	case utils.FENOEXIST:
		{
			err = install(path)
			break
		}
	}

	if err != nil {
		return err
	}

	gitPath := path + "/.git"
	_, ferr = utils.Exists(gitPath)

	switch ferr {
	case utils.FEINSFPERM:
		{
			err = fmt.Errorf("needs permission to access the linux-toolbox git folder at %s", gitPath)
			break
		}
	case utils.FENOEXIST:
		{
			err = install(path)
			break
		}
	}

	return err
}

func install(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	utils.MkdirIfNotExist(path)

	id, _ := utils.GetGourdID()
	err := os.Chown(path, id.UID, id.GID)

	// TODO: Git clone the linux toolbox

	return err
}

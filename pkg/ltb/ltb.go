package ltb

import (
	"fmt"
	"os"
	"os/user"

	"github.com/zahfox/gourd/pkg/utils"
	"github.com/zahfox/gourd/pkg/utils/git"
)

const (
	ltburl   string = "https://github.com/ZahFox/linux-toolbox"
	ltbpath  string = "/opt/gourd/linux-toolbox"
	userpath string = ".gourd/linux-toolbox"
)

// EnsureInstalled makes sure that the linx-toolbox is properly installed
func EnsureInstalled() error {
	var err error = nil
	_, ferr := utils.Exists(ltbpath)

	switch ferr {
	case utils.FEINSFPERM:
		{
			err = fmt.Errorf("needs permission to access the linux-toolbox at %s", ltbpath)
			break
		}
	case utils.FENOEXIST:
		{
			err = install(ltbpath)
			break
		}
	}

	if err != nil {
		return err
	}

	gitPath := ltbpath + "/.git"
	_, ferr = utils.Exists(gitPath)

	switch ferr {
	case utils.FEINSFPERM:
		{
			err = fmt.Errorf("needs permission to access the linux-toolbox git folder at %s", gitPath)
			break
		}
	case utils.FENOEXIST:
		{
			err = install(ltbpath)
			break
		}
	}

	return err
}

// Install will install or reinstall linux-toolbox
func Install() error {
	return install(ltbpath)
}

// InstallForUser will install or reinstall linux-toolbox and then link it for a user
func InstallForUser(username string) error {
	if err := install(ltbpath); err != nil {
		return err
	}

	usr, err := user.Lookup(username)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", usr.HomeDir, userpath)
	return os.Symlink(ltbpath, path)
}

func install(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	err := git.Clone(ltburl, ltbpath)
	if err != nil {
		return err
	}

	id, _ := utils.GetGourdID()
	err = os.Chown(path, id.UID, id.GID)

	return err
}

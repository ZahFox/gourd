package git

import (
	git "gopkg.in/src-d/go-git.v4"
	"os"
)

// Clone will download a remote git repository
func Clone(url string, dest string) error {
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}

package vagrantutil

import (
	"io"
	"os"
)

// writeFile is a fork of ioutil.WriteFile that additionally calls fsync on
// the written file.
func writeFile(name string, p []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(p)
	if err == nil && n < len(p) {
		err = io.ErrShortWrite
	}
	return nonil(err, f.Sync(), f.Close())
}

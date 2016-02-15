package vagrantutil

import (
	"errors"
	"strings"
)

// The following errors are returned by the Wait function:
var (
	ErrBoxAlreadyExists  = errors.New("box already exists")
	ErrBoxInvalidVersion = errors.New("invalid box version")
	ErrBoxNotAvailable   = errors.New("box is not available")
	ErrBoxDownload       = errors.New("unable to download the box")
	ErrNoMachine         = errors.New("no Vagrantfile found")
)

// Wait is a convenience function that consumes Vagrant output looking
// for well-known errors. It returns an error variable for any known
// failure, which makes it easier to recover from it.
func Wait(out <-chan *CommandOutput, err error) error {
	var e error
	for o := range out {
		for s, err := range errMapping {
			if strings.Contains(o.Line, s) {
				e = nonil(e, err)
			}
		}
		e = nonil(e, o.Error)
	}
	return nonil(e, err)
}

var errMapping = map[string]error{
	"The box you're attempting to add already exists.":                ErrBoxAlreadyExists,
	"Gem::Requirement::BadRequirementError":                           ErrBoxInvalidVersion,
	"could not be accessed in the remote catalog.":                    ErrBoxNotAvailable,
	"An error occurred while downloading the remote file.":            ErrBoxDownload,
	"A Vagrant environment or target machine is required to run this": ErrNoMachine,
}

func nonil(err ...error) error {
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return nil
}

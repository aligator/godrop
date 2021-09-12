package filesystem

import (
	"errors"
	"github.com/aligator/checkpoint"
	"github.com/aligator/godrop/server/repository"
	"github.com/spf13/afero"
	"os"
	"strings"
)

const (
	stateSuffix  = "#"
	uploadSuffix = ".uploading" + stateSuffix
	deleteSuffix = ".deleted" + stateSuffix
)

var states = []string{"", uploadSuffix, deleteSuffix}

func trimStateSuffix(path string) string {
	for _, state := range states {
		path = strings.TrimSuffix(path, state)
	}

	return path
}

func getStateFilename(path string, stateSuffix string) string {
	trimStateSuffix(path)
	return path + stateSuffix
}

func getStateOf(path string) string {
	// Go through all states, ignoring the first one "" which would stop
	// checking all other states as it fits always.
	for _, state := range states[1:] {
		if strings.HasSuffix(path, state) {
			return state
		}
	}

	return ""
}

// openFileById is the equivalent to fs.OpenFile and opens the file
// no matter what state-suffix it currently has.
func openFileById(fs afero.Fs, id string, flag int, perm os.FileMode) (afero.File, error) {
	// Try all possible paths.
	for _, state := range states {
		file, err := fs.OpenFile(getStateFilename(id, state), flag, perm)
		if errors.Is(err, afero.ErrFileNotFound) {
			continue
		}

		return file, checkpoint.From(err)
	}

	return nil, checkpoint.From(repository.ErrFileNotFound)
}

// openById is the equivalent to fs.Open and opens the file
// no matter what state-suffix it currently has.
// It opens the file readonly.
func openById(fs afero.Fs, id string) (afero.File, error) {
	return openFileById(fs, id, os.O_RDONLY, 0)
}

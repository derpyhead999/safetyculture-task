package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := f.folders

	// Move folder to itself
	if name == dst {
		return []Folder{}, errors.New("Error: Cannot move a folder to itself")
	}

	var source_id uuid.UUID
	var dest_id uuid.UUID
	for _, f := range folders {
		if strings.HasSuffix(f.Paths, name) { // Find source folder
			source_id = f.OrgId
		}
		if strings.HasSuffix(f.Paths, dst) { // Find dest folder
			dest_id = f.OrgId
		}
	}
	// Source folder doesn't exist
	if source_id.IsNil() {
		return []Folder{}, errors.New("Error: Source folder does not exist")
	}
	// Destination folder doesn't exist
	if dest_id.IsNil() {
		return []Folder{}, errors.New("Error: Destination folder does not exist")
	}

	// Move folder to different organisation
	if source_id != dest_id {
		return []Folder{}, errors.New("Error: Cannot move a folder to a different organization")
	}

	// get all child folders

	// find path of 'dst' from root

	// append 'name'.{child_path} to 'path of dst from root'

	return []Folder{}, nil
}

package folder

import (
	"errors"
	"regexp"
	"strings"
)

// O(n^2) complexity, a nested for loop used
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := f.folders

	// Move folder to itself
	if name == dst {
		return []Folder{}, errors.New("Error: Cannot move a folder to itself")
	}

	source_folder := Folder{}
	dest_folder := Folder{}
	for _, f := range folders {
		if strings.HasSuffix(f.Paths, name) { // Find source folder
			source_folder = f
		}
		if strings.HasSuffix(f.Paths, dst) { // Find dest folder
			dest_folder = f
		}
	}
	// Source folder doesn't exist
	if source_folder == (Folder{}) {
		return []Folder{}, errors.New("Error: Source folder does not exist")
	}
	// Destination folder doesn't exist
	if dest_folder == (Folder{}) {
		return []Folder{}, errors.New("Error: Destination folder does not exist")
	}

	// Cannot move folder to child of itself
	re := regexp.MustCompile(`(^|\.)` + regexp.QuoteMeta(name) + `(\..+|$)`)
	if re.MatchString(dest_folder.Paths) {
		return []Folder{}, errors.New("Error: Cannot move a folder to a child of itself")
	}

	// Cannot move folder to different organisation
	if source_folder.OrgId != dest_folder.OrgId {
		return []Folder{}, errors.New("Error: Cannot move a folder to a different organization")
	}

	// get all child folders
	res := f.GetAllChildFolders(source_folder.OrgId, name)
	res = append(res, source_folder) // Include source folder to move

	// find path of 'dst' from root
	dest_path := dest_folder.Paths

	// create slice copy of 'folders'
	mod_folders := make([]Folder, len(folders))
	copy(mod_folders, folders)

	// append 'name'.{child_path_if_exist} to 'path of dst from root'
	for _, f := range res {
		for i, g := range folders {
			if f.Paths == g.Paths { // If folder found, get the 'name'.{child_path_if_exist}
				re := regexp.MustCompile(`(^|\.)` + regexp.QuoteMeta(name) + `(\..+|$)`)
				suffix := re.FindString(g.Paths)
				res_string := ""

				if suffix[0] == '.' { // No need to add '.' if alr present
					res_string = dest_path + suffix
				} else {
					res_string = dest_path + "." + suffix
				}
				mod_folders[i].Paths = res_string
			}
		}
	}
	return mod_folders, nil
}

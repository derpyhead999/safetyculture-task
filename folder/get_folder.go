package folder

import (
	"regexp"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

// Get all directories with 'name', but not at the end (can be at the start)
func IsChildFolder(folder Folder, name string) bool {
	re := regexp.MustCompile(`(^|\.)` + regexp.QuoteMeta(name) + `\..+`)
	return re.MatchString(folder.Paths)
}

func IsValidPath(folder Folder) bool {
	re := regexp.MustCompile(`^[^.]+(\.[^.]+)*$`) // Checks that path name is valid
	return re.MatchString(folder.Paths)
}

// Note: Keep in mind that GetAllChildFolders does not have an error output type, so it was assumed that any errors will only be shown through the Folder array output
// O(n) complexity, 1 for loop
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Only get folders with valid orgID
	org_folders := f.GetFoldersByOrgID(orgID)

	res := []Folder{}
	for _, f := range org_folders {
		if IsChildFolder(f, name) && IsValidPath(f) {
			res = append(res, f)
		}
	}

	return res
}

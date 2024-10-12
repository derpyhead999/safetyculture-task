package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:  "Test case 1",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
			want: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, get, tt.want)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name         string
		orgID        uuid.UUID
		folders      []folder.Folder
		parentFolder string // The folder name which we want child folders of
		want         []folder.Folder
	}{
		{
			name:  "Basic case",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "c",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
			parentFolder: "a",
			want: []folder.Folder{
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
		},
		{
			name:  "Basic case multiple folder same name",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
			parentFolder: "a",
			want: []folder.Folder{
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
		},
		{
			name:  "Advanced base case",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo.charlie",
				},
				{
					Name:  "folder 4",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta",
				},
				{
					Name:  "folder 5",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "echo",
				},
				{
					Name:  "folder 6",
					OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
					Paths: "foxtrot",
				},
			},
			parentFolder: "alpha",
			want: []folder.Folder{
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo.charlie",
				},
				{
					Name:  "folder 4",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta",
				},
			},
		},
		{
			name:  "Error case invalid folder",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
			parentFolder: "invalid_folder",
			want:         []folder.Folder{},
		},
		{
			name:  "Error case no files in organisation",
			orgID: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
			parentFolder: "a",
			want:         []folder.Folder{},
		},
		{
			name:  "Multiple subdirectory different org",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo.charlie",
				},
				{
					Name:  "folder 5",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo.omega",
				},
				{
					Name:  "folder 6",
					OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
					Paths: "alpha.bravo.phi",
				},
				{
					Name:  "folder 4",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta",
				},
			},
			parentFolder: "bravo",
			want: []folder.Folder{
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo.charlie",
				},
				{
					Name:  "folder 5",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.bravo.omega",
				},
			},
		},
		{
			name:  "Regex edge case 1",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "c",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "da.b",
				},
			},
			parentFolder: "a",
			want:         []folder.Folder{},
		},
		{
			name:  "Regex edge case 2",
			orgID: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "c",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.eb.d",
				},
			},
			parentFolder: "b",
			want:         []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.parentFolder)
			assert.Equal(t, get, tt.want)
		})
	}
}

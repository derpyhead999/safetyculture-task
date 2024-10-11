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
		// TODO: your tests here
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

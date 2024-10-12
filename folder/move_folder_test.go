package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

var folder_schema []folder.Folder = []folder.Folder{
	{
		Name:  "alpha",
		OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
		Paths: "alpha",
	},
	{
		Name:  "bravo",
		OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
		Paths: "alpha.bravo",
	},
	{
		Name:  "charlie",
		OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
		Paths: "alpha.bravo.charlie",
	},
	{
		Name:  "delta",
		OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
		Paths: "alpha.delta",
	},
	{
		Name:  "echo",
		OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
		Paths: "alpha.delta.echo",
	},
	{
		Name:  "foxtrot",
		OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
		Paths: "foxtrot",
	},
	{
		Name:  "golf",
		OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
		Paths: "golf",
	},
}

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name         string
		errorCase    bool
		errorMessage string
		sourceFolder string
		destFolder   string
		folders      []folder.Folder
		want         []folder.Folder
	}{
		{
			name:         "Basic case, root directory",
			errorCase:    false,
			errorMessage: "",
			sourceFolder: "c",
			destFolder:   "b",
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
			want: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.c",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
			},
		},
		{
			name:         "Basic case, nested directory",
			errorCase:    false,
			errorMessage: "",
			sourceFolder: "c",
			destFolder:   "b",
			folders: []folder.Folder{
				{
					Name:  "folder 1",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a",
				},
				{
					Name:  "folder 2",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.c",
				},
				{
					Name:  "folder 4",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.c.f",
				},
				{
					Name:  "folder 5",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.d",
				},
				{
					Name:  "folder 6",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.e",
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
					Paths: "a.b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.c",
				},
				{
					Name:  "folder 4",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.c.f",
				},
				{
					Name:  "folder 5",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.d",
				},
				{
					Name:  "folder 6",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "a.b.e",
				},
			},
		},
		{
			name:         "Advanced case",
			errorCase:    false,
			errorMessage: "",
			sourceFolder: "bravo",
			destFolder:   "delta",
			folders:      folder_schema,
			want: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta.bravo",
				},
				{
					Name:  "charlie",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta.bravo.charlie",
				},
				{
					Name:  "delta",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta",
				},
				{
					Name:  "echo",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta.echo",
				},
				{
					Name:  "foxtrot",
					OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
					Paths: "foxtrot",
				},
				{
					Name:  "golf",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "golf",
				},
			},
		},
		{
			name:         "Advanced case 2",
			errorCase:    false,
			errorMessage: "",
			sourceFolder: "bravo",
			destFolder:   "golf",
			folders:      folder_schema,
			want: []folder.Folder{
				{
					Name:  "alpha",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha",
				},
				{
					Name:  "bravo",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "golf.bravo",
				},
				{
					Name:  "charlie",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "golf.bravo.charlie",
				},
				{
					Name:  "delta",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta",
				},
				{
					Name:  "echo",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "alpha.delta.echo",
				},
				{
					Name:  "foxtrot",
					OrgId: uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222")),
					Paths: "foxtrot",
				},
				{
					Name:  "golf",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "golf",
				},
			},
		},
		{
			name:         "Error case folder to child",
			errorCase:    true,
			errorMessage: "Error: Cannot move a folder to a child of itself",
			sourceFolder: "bravo",
			destFolder:   "charlie",
			folders:      folder_schema,
			want:         []folder.Folder{},
		},
		{
			name:         "Error case folder to itself",
			errorCase:    true,
			errorMessage: "Error: Cannot move a folder to itself",
			sourceFolder: "bravo",
			destFolder:   "bravo",
			folders:      folder_schema,
			want:         []folder.Folder{},
		},
		{
			name:         "Error case folder to different org",
			errorCase:    true,
			errorMessage: "Error: Cannot move a folder to a different organization",
			sourceFolder: "bravo",
			destFolder:   "foxtrot",
			folders:      folder_schema,
			want:         []folder.Folder{},
		},
		{
			name:         "Error case source folder not exist",
			errorCase:    true,
			errorMessage: "Error: Source folder does not exist",
			sourceFolder: "invalid_folder",
			destFolder:   "delta",
			folders:      folder_schema,
			want:         []folder.Folder{},
		},
		{
			name:         "Error case destination folder not exist",
			errorCase:    true,
			errorMessage: "Error: Destination folder does not exist",
			sourceFolder: "bravo",
			destFolder:   "invalid_folder",
			folders:      folder_schema,
			want:         []folder.Folder{},
		},
		{
			name:         "Base case no subdirectories",
			errorCase:    false,
			errorMessage: "",
			sourceFolder: "b",
			destFolder:   "c",
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
					Paths: "c",
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
					Paths: "c.b",
				},
				{
					Name:  "folder 3",
					OrgId: uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111")),
					Paths: "c",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.sourceFolder, tt.destFolder)
			if tt.errorCase {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, get, tt.want)
			}
		})
	}
}

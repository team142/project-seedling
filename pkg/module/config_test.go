package module

import (
	"os"
	"testing"
)

func TestConfig_getFileDir(t *testing.T) {
	type fields struct {
		Version              string
		Structs              []string
		OutputDir            string
		Directories          DirectoryConfig
		OutputDirPermissions os.FileMode
		API                  string
		Auth                 bool
		FileName             string
		Generate             DiscoveryFunction
		CreateRouter         bool
		CreateHandler        bool
		CreateMiddleware     bool
		Separator            rune
	}
	type args struct {
		base     string
		fileName string
		dir      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "basic test - no version",
			fields: fields{
				Version:   "",
				Separator: '/',
			},
			args: args{
				fileName: "%s.go",
				base:     "./",
				dir:      "router",
			},
			want: "./router/%s.go",
		},
		{
			name: "basic test - no version - weird separator",
			fields: fields{
				Version:   "",
				Separator: '~',
			},
			args: args{
				fileName: "%s.go",
				base:     ".",
				dir:      "router",
			},
			want: ".~router~%s.go",
		},
		{
			name: "basic test - no base - with version",
			fields: fields{
				Separator: '/',
				Version:   "V1",
			},
			args: args{
				fileName: "%s.go",
				dir:      "router",
				base:     "",
			},
			want: "./router/v1/%s.go",
		},
		{
			name: "basic test - base - with version",
			fields: fields{
				Separator: '/',
				Version:   "V1",
			},
			args: args{
				fileName: "%s.go",
				dir:      "router",
				base:     "./test",
			},
			want: "./test/router/v1/%s.go",
		},
		{
			name: "basic test - no version",
			fields: fields{
				Separator: '/',
				Version:   "",
			},
			args: args{
				fileName: "%s.go",
				base:     ".",
				dir:      "router",
			},
			want: "./router/%s.go",
		},
		{
			name: "basic test - no version",
			fields: fields{
				Separator: '/',
				Version:   "",
			},
			args: args{
				fileName: "file-name.go",
				base:     ".",
				dir:      "router",
			},
			want: "./router/file-name.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Version:              tt.fields.Version,
				Structs:              tt.fields.Structs,
				OutputDir:            tt.fields.OutputDir,
				Directories:          tt.fields.Directories,
				OutputDirPermissions: tt.fields.OutputDirPermissions,
				APIPath:              tt.fields.API,
				Auth:                 tt.fields.Auth,
				FileName:             tt.fields.FileName,
				DiscoverFunction:     tt.fields.Generate,
				CreateRouter:         tt.fields.CreateRouter,
				CreateHandler:        tt.fields.CreateHandler,
				CreateMiddleware:     tt.fields.CreateMiddleware,
				Separator:            tt.fields.Separator,
			}
			if got := c.getFullFilePath(tt.args.base, tt.args.dir, tt.args.fileName); got != tt.want {
				t.Errorf("getFullFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

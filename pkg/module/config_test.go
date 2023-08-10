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

func Test_validateFinalCharacter(t *testing.T) {
	type args struct {
		base string
		sep  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple Test",
			args: args{
				base: "",
				sep:  "/",
			},
			want: "/",
		},
		{
			name: "Simple Test With Ending",
			args: args{
				base: "/",
				sep:  "/",
			},
			want: "/",
		},
		{
			name: "Simple Test with simple base",
			args: args{
				base: "base",
				sep:  "/",
			},
			want: "base/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateFinalCharacter(tt.args.base, tt.args.sep); got != tt.want {
				t.Errorf("validateFinalCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetAPIPath(t *testing.T) {
	type fields struct {
		Version string
		APIPath string
	}
	type args struct {
		base string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Simple Test",
			fields: fields{
				Version: "",
				APIPath: "",
			},
			args: args{
				base: "",
			},
			want: "/",
		},
		{
			name: "Simple Test with base",
			fields: fields{
				Version: "",
				APIPath: "",
			},
			args: args{
				base: "base",
			},
			want: "/base",
		},
		{
			name: "Simple Test With API Path",
			fields: fields{
				Version: "",
				APIPath: "APIPath",
			},
			args: args{
				base: "",
			},
			want: "/APIPath",
		},
		{
			name: "Simple Test With API Path and version",
			fields: fields{
				Version: "V1",
				APIPath: "APIPath",
			},
			args: args{
				base: "",
			},
			want: "/APIPath/v1",
		},
		{
			name: "Simple Test With API Path, version and base",
			fields: fields{
				Version: "V1",
				APIPath: "APIPath",
			},
			args: args{
				base: "base",
			},
			want: "/APIPath/v1/base",
		},
		{
			name: "Simple Test With API Path, version and base",
			fields: fields{
				Version: "V1",
				APIPath: "/APIPath/v1/test",
			},
			args: args{
				base: "base",
			},
			want: "/APIPath/v1/test/base",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Version: tt.fields.Version,
				APIPath: tt.fields.APIPath,
			}
			if got := c.GetAPIPath(tt.args.base); got != tt.want {
				t.Errorf("GetAPIPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAPIPath(t *testing.T) {
	type args struct {
		start string
		end   string
		sep   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple",
			args: args{
				start: "",
				end:   "",
				sep:   "/",
			},
			want: "/",
		},
		{
			name: "Simple No Duplicate",
			args: args{
				start: "/",
				end:   "",
				sep:   "/",
			},
			want: "/",
		},
		{
			name: "Simple start",
			args: args{
				start: "start",
				end:   "",
				sep:   "/",
			},
			want: "/start",
		},
		{
			name: "Simple start and end",
			args: args{
				start: "start",
				end:   "end",
				sep:   "/",
			},
			want: "/start/end",
		},
		{
			name: "Simple sep for start and end",
			args: args{
				start: "/",
				end:   "end",
				sep:   "/",
			},
			want: "/end",
		},
		{
			name: "Simple API Path test with version",
			args: args{
				start: "/APIPath",
				end:   "v1",
				sep:   "/",
			},
			want: "/APIPath/v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAPIPath(tt.args.start, tt.args.end, tt.args.sep); got != tt.want {
				t.Errorf("getAPIPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

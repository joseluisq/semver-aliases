package aliases

import (
	"reflect"
	"testing"
)

func TestFromVersion(t *testing.T) {
	type args struct {
		version string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty string",
			args: args{
				version: "",
			},
			want: []string(nil),
		},
		{
			name: "invalid semver",
			args: args{
				version: "1.0",
			},
			want: []string(nil),
		},
		{
			name: "semver pre-release",
			args: args{
				version: "1.0.0-beta.1",
			},
			want: []string{"1.0.0-beta.1"},
		},
		{
			name: "semver pre-release (prefixed)",
			args: args{
				version: "v1.2.0-beta.0",
			},
			want: []string{"1.2.0-beta.0"},
		},
		{
			name: "valid semver release aliases",
			args: args{
				version: "v1.2.0",
			},
			want: []string{"1", "1.2", "1.2.0"},
		},
		{
			name: "valid semver release aliases",
			args: args{
				version: "0.1.0",
			},
			want: []string{"0", "0.1", "0.1.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromVersion(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromTagNames(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "null slice value",
			args: args{
				tags: nil,
			},
			want: []string(nil),
		},
		{
			name: "empty string slice with no values",
			args: args{
				tags: []string{},
			},
			want: []string(nil),
		},
		{
			name: "string slice with empty values",
			args: args{
				tags: []string{"", "", ""},
			},
			want: nil,
		},
		{
			name: "single release string",
			args: args{
				tags: []string{"1.0.0"},
			},
			want: []string{"1.0.0"},
		},
		{
			name: "single prefixed release string",
			args: args{
				tags: []string{"v2.5.1"},
			},
			want: []string{"2.5.1"},
		},
		{
			name: "multiple prefixed release string values",
			args: args{
				tags: []string{"v5.8.7", "v1.7.9"},
			},
			want: []string{"1.7.9", "5.8.7"},
		},
		{
			name: "mixed release string values",
			args: args{
				tags: []string{"v2.38.1", "v10.18.3", "v6.32.2"},
			},
			want: []string{"10.18.3", "2.38.1", "6.32.2"},
		},
		{
			name: "mixed release string and empty values",
			args: args{
				tags: []string{"5.4.0", "", "v20.8.7"},
			},
			want: []string{"20.8.7", "5.4.0"},
		},
		{
			name: "deduplicate release string values",
			args: args{
				tags: []string{"v1.5.8", "1.8.5", "v1.5.8"},
			},
			want: []string{"1.5.8", "1.8.5"},
		},
		{
			name: "full release string values and alias names",
			args: args{
				tags: []string{"latest", "10.22.0", "v10.0.22", "", "v10.22.0", "v2", "v2.0", "2.0.1", ""},
			},
			want: []string{"10.0.22", "10.22.0", "2", "2.0", "2.0.1", "latest"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromTagNames(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromTagNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

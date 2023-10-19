package internal

import "testing"

func Test_repoHttpUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{"github.com/go-gorm/gen"},
			"https://github.com/go-gorm/gen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRepoUrl(tt.args.url); got != tt.want {
				t.Errorf("repoHttpUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

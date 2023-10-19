package internal

import (
	"os"
	"testing"
)

func Test_localRepoRoot(t *testing.T) {

	localRepoRoot()
	return

	codeDir := "./code"
	tests := []struct {
		name     string
		setup    func()
		want     string
		teardown func()
	}{
		{
			"",
			func() {
				if err := os.Setenv("CODE", codeDir); err != nil {
					t.Fatal(err)
				}
			},
			codeDir,
			func() {
				os.Clearenv()
			},
		},
		{
			"",
			func() {

			},
			"",
			func() {
				os.Clearenv()
			},
		},
		{
			"",
			func() {
				if err := os.Setenv("CODE", codeDir); err != nil {
					t.Fatal(err)
				}
			},
			codeDir,
			func() {
				os.Clearenv()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := localRepoRoot(); got != tt.want {
				t.Errorf("localRepoRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

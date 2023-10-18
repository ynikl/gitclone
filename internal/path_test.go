package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestMakeDstPath(t *testing.T) {

	gorm := "github.com/go-gorm/gen"
	home, err := homedir.Dir()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name       string
		arg        string
		want       string
		tearupFunc func()
	}{
		{
			"",
			gorm,
			filepath.Join(home, gorm),
			func() {
				os.RemoveAll(filepath.Join(home, gorm))
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got, err := makeDstPath(tt.arg); err != nil || got != tt.want {
				t.Fatalf("testcase:%s\t%v\t want %v", tt.name, got, tt.want)
			}
		})
	}

}

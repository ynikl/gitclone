package internal

import "testing"

func Test_localRepoRoot(t *testing.T) {
	tests := []struct {
		name     string
		pre      func()
		want     string
		teardown func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := localRepoRoot(); got != tt.want {
				t.Errorf("localRepoRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

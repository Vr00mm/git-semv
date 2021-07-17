package semv

import (
	"reflect"
	"testing"
)

func Test_username(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := username()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. username() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. username() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_latestCommit(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := latestCommit()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. latestCommit() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. latestCommit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_meta(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := meta()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. meta() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. meta() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

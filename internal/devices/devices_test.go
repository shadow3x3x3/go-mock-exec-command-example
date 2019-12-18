package devices

import (
	"fmt"
	"go-mock-exec-command-example/test"
	"go-mock-exec-command-example/util/cmd"
	"os"
	"strings"
	"testing"
)

func Test_isExists(t *testing.T) {
	fakeCommand := test.MakeFakeCommand("TestLsblkWithName") // Fill in your mock command function name
	cmd.SetupTestCmd(fakeCommand)
	defer cmd.TearDownTestCmd()

	type args struct {
		targetDev string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"exists_dev", args{"sda"}, true},
		{"exists_dev", args{"sda1"}, true},
		{"not_exists_dev", args{"sdb1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExists(tt.args.targetDev); got != tt.want {
				t.Errorf("isExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSize(t *testing.T) {
	fakeCommand := test.MakeFakeCommand("TestLsblkSize")
	cmd.SetupTestCmd(fakeCommand)
	defer cmd.TearDownTestCmd()

	type args struct {
		targetDev string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid_device", args{"sda"}, "1.8T", false},
		{"valid_device", args{"sda1"}, "100M", false},
		{"invalid_device", args{"sda2"}, "", true},
		{"invalid_device", args{"sdg"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSize(tt.args.targetDev)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Mock Commands
func TestLsblkWithName(t *testing.T) {
	if !test.IsTestEnv() {
		return
	}

	// cmdWithParams gets []string{"lsblk", "/dev/XXX"}
	// The last argument in os.Args should be "lsblk /dev/XXX"
	// It appended in the test/test.go MakeFakeCommand()
	cmdWithParams := strings.Split(os.Args[len(os.Args)-1], " ")

	// Check command is valid.
	// Just depends on your command usage.
	cmd := cmdWithParams[0] // lsblk
	if cmd != "lsblk" {
		os.Exit(1) // Mock command fails with status code 1
	}

	table := map[string]bool{
		"/dev/loop1": true,
		"/dev/sda":   true,
		"/dev/sda1":  true,
		"/dev/sdb":   true,
	}

	dev := cmdWithParams[1] // /dev/XXX
	if r, ok := table[dev]; !ok || !r {
		os.Exit(1)
		return
	}

	os.Exit(0)
}

func TestLsblkSize(t *testing.T) {
	if !test.IsTestEnv() {
		return
	}

	// cmdWithParams gets []string{"lsblk", "-lno", "SIZE", "/dev/XXX"}
	cmdWithParams := strings.Split(os.Args[len(os.Args)-1], " ")

	cmd := cmdWithParams[0] // lsblk
	if cmd != "lsblk" {
		os.Exit(1)
	}

	if cmdWithParams[1] != "-lno" || cmdWithParams[2] != "SIZE" {
		os.Exit(1)
	}

	table := map[string]string{
		"/dev/sda":  ` 1.8T`,
		"/dev/sda1": ` 100M`,
	}

	dev := cmdWithParams[3] // /dev/XXX
	r, ok := table[dev]
	if !ok {
		os.Exit(1)
		return
	}

	fmt.Fprintf(os.Stdout, r) // Mock command output
	os.Exit(0)
}

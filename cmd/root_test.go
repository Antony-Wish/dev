package cmd

import (
	"testing"

	"github.com/wish/dev"
	"github.com/wish/dev/test"

	"github.com/spf13/afero"
	"gotest.tools/env"
)

func TestInitializeWithDevConfigSet(t *testing.T) {
	defer env.Patch(t, "DEV_CONFIG", "/home/test/.dev.yaml")()
	appConfig.SetFs(afero.NewMemMapFs())
	test.CreateConfigFile(appConfig.GetFs(), test.BigCoConfig, "/home/test/.dev.yaml")

	Initialize()

	var cmdTests = []struct {
		ProjectName string
		Aliases     []string
	}{
		{"postgresql", []string{"pg", "db"}},
		{"frontend", []string{"shiny"}},
	}

	for _, test := range cmdTests {
		cmd, _, err := rootCmd.Find([]string{test.ProjectName})
		if err != nil {
			t.Errorf("Expected to find '%s' project but got err: '%s'", test.ProjectName, err)
		}
		if cmd == nil {
			t.Errorf("Expected to find '%s' cmd, but got nil", test.ProjectName)
		}
		if cmd.Use != test.ProjectName {
			t.Errorf("Expected cmd to be named '%s', but got '%s'", test.ProjectName, cmd.Short)
		}
		if len(test.Aliases) != len(cmd.Aliases) {
			t.Errorf("Expected to find %d %s aliases, but got %d", len(test.Aliases), test.ProjectName,
				len(cmd.Aliases))
		}
		for _, alias := range test.Aliases {
			if !dev.SliceContainsString(cmd.Aliases, alias) {
				t.Errorf("Expected to find alias '%s' of '%s' cmd", alias, test.ProjectName)
			}
		}

		subCommands := []string{dev.UP, dev.BUILD, dev.PS, dev.SH, dev.UP}
		for _, subCmd := range subCommands {
			sCmd, _, err := cmd.Find([]string{subCmd})
			if err != nil {
				t.Errorf("Expected to find subcommand '%s' of %s but got err: '%s'", subCmd, test.ProjectName, err)
			}
			if sCmd == nil {
				t.Errorf("Expected to find '%s' sub-command of %s, but got nil", subCmd, test.ProjectName)
			}

			if sCmd.Use != subCmd {
				t.Errorf("Expected cmd to be named '%s', but got '%s'", subCmd, sCmd.Short)
			}
		}
	}
}

func TestInitializeWithoutDevConfigSet(t *testing.T) {
	homedir := "/home/test"
	defer env.Patch(t, "DEV_CONFIG", "")() // set to nothing so i can test locally where I set it
	defer env.Patch(t, "HOME", homedir)()

	appConfig.SetFs(afero.NewMemMapFs())
	test.CreateConfigFile(appConfig.GetFs(), test.BigCoConfig, homedir+"/.config/dev/dev.yaml")

	Initialize()

	var cmdTests = []struct {
		ProjectName string
		Aliases     []string
	}{
		{"postgresql", []string{"pg", "db"}},
		{"frontend", []string{"shiny"}},
	}

	for _, test := range cmdTests {
		cmd, _, err := rootCmd.Find([]string{test.ProjectName})
		if err != nil {
			t.Errorf("Expected to find '%s' project but got err: '%s'", test.ProjectName, err)
		}
		if cmd == nil {
			t.Errorf("Expected to find '%s' cmd, but got nil", test.ProjectName)
		}
		if cmd.Use != test.ProjectName {
			t.Errorf("Expected cmd to be named '%s', but got '%s'", test.ProjectName, cmd.Short)
		}
		if len(test.Aliases) != len(cmd.Aliases) {
			t.Errorf("Expected to find %d %s aliases, but got %d", len(test.Aliases), test.ProjectName,
				len(cmd.Aliases))
		}
		for _, alias := range test.Aliases {
			if !dev.SliceContainsString(cmd.Aliases, alias) {
				t.Errorf("Expected to find alias '%s' of '%s' cmd", alias, test.ProjectName)
			}
		}

		subCommands := []string{dev.UP, dev.BUILD, dev.PS, dev.SH, dev.UP}
		for _, subCmd := range subCommands {
			sCmd, _, err := cmd.Find([]string{subCmd})
			if err != nil {
				t.Errorf("Expected to find subcommand '%s' of %s but got err: '%s'", subCmd, test.ProjectName, err)
			}
			if sCmd == nil {
				t.Errorf("Expected to find '%s' sub-command of %s, but got nil", subCmd, test.ProjectName)
			}

			if sCmd.Use != subCmd {
				t.Errorf("Expected cmd to be named '%s', but got '%s'", subCmd, sCmd.Short)
			}
		}
	}
}

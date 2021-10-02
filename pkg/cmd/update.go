package cmd


import (
	"fmt"
	"os"
	"sort"
	"path/filepath"
	cli "github.com/spf13/cobra"
    "os/exec"

	"github.com/6uhrmittag/boilr/pkg/boilr"
	//"github.com/6uhrmittag/boilr/pkg/template"
	"github.com/6uhrmittag/boilr/pkg/util/exit"
	"github.com/6uhrmittag/boilr/pkg/util/validate"
   // "github.com/6uhrmittag/boilr/pkg/util/osutil"

)

// ListTemplates returns a list of templates saved in the local template registry.
func ListTemplatesR() (map[string]bool, error) {
	d, err := os.Open(boilr.Configuration.TemplateDirPath)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	nameSet := make(map[string]bool)
	for _, name := range names {
		nameSet[name] = true
	}

	return nameSet, nil
}

func RunUpdate(templateupdatefile string) {
    cmd, err := exec.Command("/bin/sh", templateupdatefile).Output()
    if err != nil {
    fmt.Printf("error %s", err)
    }
    output := string(cmd)
    //return output
    fmt.Printf("%s", output)


}

// Validate contains the cli-command for validating templates.
var Update = &cli.Command{
	Use:   "update",
	Short: "Updates local templates",
	Run: func(_ *cli.Command, args []string) {


	    MustValidateArgs(args, []validate.Argument{})

		MustValidateTemplateDir()

		templateNames, err := ListTemplatesR()
		if err != nil {
			exit.Error(fmt.Errorf("list: %s", err))
		}

		// For keeping the names in order
		names := []string{}
		for name := range templateNames {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
            var updatefile = filepath.Join(os.Getenv("HOME"), boilr.Configuration.ConfigDirPath, "update_" + name + ".sh")

            if _, err := os.Stat(updatefile); err == nil {
              fmt.Printf("Running Update\n");
              RunUpdate(updatefile)
            } else {
              fmt.Printf("Update file is missing %s\n", updatefile)
            }

		}

    exit.OK("Templates updated")
	},
}

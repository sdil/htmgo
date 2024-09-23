package downloadtemplate

import (
	"flag"
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/run"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func DownloadTemplate(outPath string) {
	cwd, _ := os.Getwd()

	flag.Parse()

	outPath = strings.ReplaceAll(outPath, "\n", "")
	outPath = strings.ReplaceAll(outPath, " ", "-")
	outPath = strings.ToLower(outPath)

	if outPath == "" {
		fmt.Println("Please provide a name for your app.")
		return
	}

	templateName := "starter-template"
	templatePath := filepath.Join("templates", "starter")

	re := regexp.MustCompile(`[^a-zA-Z]+`)
	// Replace all non-alphabetic characters with an empty string
	newModuleName := re.ReplaceAllString(outPath, "")

	tempOut := newModuleName + "_temp_" + time.Now().String()

	install := exec.Command("git", "clone", "https://github.com/maddalax/htmgo", "--depth=1", tempOut)
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	err := install.Run()

	if err != nil {
		println("Error downloading template %v\n", err)
		return
	}

	newDir := filepath.Join(cwd, outPath)

	dirutil.CopyDir(filepath.Join(tempOut, templatePath), newDir, func(path string, exists bool) bool {
		return true
	})

	dirutil.DeleteDir(tempOut)

	process.SetWorkingDir(newDir)

	commands := [][]string{
		{"go", "get", "github.com/maddalax/htmgo/framework"},
		{"go", "get", "github.com/maddalax/htmgo/framework-ui"},
		{"git", "init"},
	}

	for _, command := range commands {
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Dir = newDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Error executing command %s, error: %s\n", strings.Join(command, " "), err.Error())
			return
		}
	}

	_ = util.ReplaceTextInFile(filepath.Join(newDir, "go.mod"),
		fmt.Sprintf("module %s", templateName),
		fmt.Sprintf("module %s", newModuleName))

	_ = util.ReplaceTextInDirRecursive(newDir, templateName, newModuleName, func(file string) bool {
		return strings.HasSuffix(file, ".go")
	})

	fmt.Printf("Setting up the project in %s\n", newDir)
	process.SetWorkingDir(newDir)
	run.Setup()
	process.SetWorkingDir("")

	fmt.Println("Template downloaded successfully.")
	fmt.Println("To start the development server, run the following commands:")
	fmt.Printf("cd %s && htmgo watch\n", outPath)

	fmt.Printf("To build the project, run the following command:\n")
	fmt.Printf("cd %s && htmgo build\n", outPath)
}

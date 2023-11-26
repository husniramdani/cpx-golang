package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// gentle notes
// git reset HEAD^ to uncommit

func main() {
  // get list of new files
  listofFiles := getListOfFiles()
  // define new files and updated files
	newFiles, updatedFiles := filteredFiles(listofFiles)

  fmt.Println("List of new files :", newFiles)
  fmt.Println("LIst of updated files :", updatedFiles)
  fmt.Println()

  // commit new files
  gitAddCmd(newFiles)
  gitCommitCmd("add", newFiles)
  fmt.Println() 
  // commit  updated files
  gitAddCmd(updatedFiles)
  gitCommitCmd("update", updatedFiles)
}

func getListOfFiles() string{
  cmd := exec.Command("git", "status", "--short", "--untracked-files=all")
  output, err := cmd.Output()

  if err != nil {
    fmt.Println("Error :", err)
    return "error";
  }
  return string(output);
}

func filteredFiles(statusOutput string) (string, string) {
	var newFiles, updatedFiles []string
	lines := strings.Split(statusOutput, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "??") {
			newFiles = append(newFiles, strings.TrimSpace(strings.TrimPrefix(line, "??")))
		} else {
      updatedFiles = append(updatedFiles, strings.TrimSpace(strings.TrimPrefix(line, " M")))
    }
	}

	return strings.Join(newFiles, " "), strings.Join(updatedFiles, " ")
}

func gitAddCmd(files string){
  // split the space-separated file names
  fileList := strings.Fields(files)

  // create the git add command with the file names as arguments
  cmd := exec.Command("git", append([]string{"add"}, fileList...)...)
  _, err := cmd.Output()

  if err != nil {
    fmt.Println("Error :", err)
    return 
  }
  fmt.Println("Success git add", files);
}

func gitCommitCmd(key string, files string){
  // define the messages
  messages := fmt.Sprintf("%s %s", key, files)
  // create the git commit command
  cmd := exec.Command("git", "commit", "-m", messages)
  _, err := cmd.Output()

  if err != nil {
    fmt.Println("Error :", err)
    return
  }
  fmt.Println("Success git commit", messages)
}

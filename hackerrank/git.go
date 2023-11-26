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
  listOfFiles := getListOfFiles()
  // define new files and updated files
  fmt.Println(listOfFiles)
	newFiles, updatedFiles, deletedFiles := filteredFiles(listOfFiles)

  // commit new files
  if newFiles != "" {
    fmt.Println("List of new files :", newFiles)
    gitAddCmd(newFiles)
    gitCommitCmd("add", newFiles)
  } else {
    fmt.Println("There are no new files")
  }
  fmt.Println() 
  // commit updated files
  if updatedFiles != "" {
    fmt.Println("List of updated files :", updatedFiles)
    gitAddCmd(updatedFiles)
    gitCommitCmd("update", updatedFiles)
  } else {
    fmt.Println("There are no updated files")
  }
  // commit deleted files
  if deletedFiles != "" {
    fmt.Println("List of deleted files :", deletedFiles)
    gitAddCmd(deletedFiles)
    gitCommitCmd("deleted", deletedFiles)
  } else {
    fmt.Println("There are no deletedFiles files")
  }

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

func filteredFiles(statusOutput string) (string, string, string) {
	var newFiles, updatedFiles, deletedFiles []string
	lines := strings.Split(statusOutput, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "??") {
			newFiles = append(newFiles, strings.TrimSpace(strings.TrimPrefix(line, "??")))
		}
    if strings.HasPrefix(line, " M") {
      updatedFiles = append(updatedFiles, strings.TrimSpace(strings.TrimPrefix(line, " M")))
    }
    if strings.HasPrefix(line, " D") {
      deletedFiles = append(deletedFiles, strings.TrimSpace(strings.TrimPrefix(line, " D")))
    }
	}

	return strings.Join(newFiles, " "), strings.Join(updatedFiles, " "), strings.Join(updatedFiles, " ")
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
  fmt.Printf("Success git commit %s \n\n", messages)
}

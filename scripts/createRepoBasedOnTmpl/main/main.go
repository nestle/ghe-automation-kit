package main

import (
	"fmt"
	"os"

	"github.com/nestle/ghe-automation-kit/creator"
)

func main() {
	// Check if the required arguments are provided
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run scripts/createRepoBasedOnTmpl/main/main.go <repository-name> <organization-name> <repository-visibility> <template-repository>")
		return
	}

	repoName := os.Args[1]
	orgName := os.Args[2]
	repoVisibility := os.Args[3]
	templateRepo := os.Args[4]

	// Create the repository
	err := creator.CreateRepositoryBasedOnTemplate(repoName, orgName, repoVisibility, templateRepo)
	if err != nil {
		fmt.Printf("Failed to create repository: %s\n", err.Error())
		return
	}

	fmt.Printf("Repository '%s' with visibility '%s' successfully created inside '%s' organization based on template '%s'.\n", repoName, repoVisibility, orgName, templateRepo)
}
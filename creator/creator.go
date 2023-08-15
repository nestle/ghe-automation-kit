package creator

import (
	"os"
	"fmt"
	"os/exec"

	"github.com/nestle/ghe-automation-kit/assigner"
	"github.com/nestle/ghe-automation-kit/remover"
)

// Creates a new GitHub repository
func CreateRepository(repoName, orgName, repoVisibility string) error {
	cmd := exec.Command("gh", "repo", "create", orgName+"/"+repoName, "--"+repoVisibility)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Print the error message to the runner console
		fmt.Fprintf(os.Stderr, "error creating GitHub repository: %v\n%s", err, string(output))
		return err // Return the error to handle in the caller function
	}
	return nil
}

// Creates a new GitHub repository based on template
func CreateRepositoryBasedOnTemplate(repoName, orgName, repoVisibility, templateRepo string) error {
	cmd := exec.Command("gh", "repo", "create", orgName+"/"+repoName, "--"+repoVisibility, "--template "+templateRepo)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Print the error message to the runner console
		fmt.Fprintf(os.Stderr, "error creating GitHub repository: %v\n%s", err, string(output))
		return err // Return the error to handle in the caller function
	}
	return nil
}

// Creates a GitHub team and assigns maintainers to it
func CreateTeam(orgName, teamSlug, mainMaintainer, deputyMaintainer string) error {
	// Create the GitHub team using the `gh api` command
	createTeamCommand := exec.Command("gh", "api", "-X", "POST", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "orgs/"+orgName+"/teams", "-f", "name="+teamSlug, "-f", "notification_setting=notifications_enabled", "-f", "privacy=closed")
	output, err := createTeamCommand.CombinedOutput()
	if err != nil {
		// Print the error message to the runner console
		fmt.Fprintf(os.Stderr, "error creating GitHub team: %v\n%s", err, string(output))
		return err // Return the error to handle in the caller function
	}

	// Remove all users from the GitHub team
	err = remover.RemoveUsersFromTeam(orgName, teamSlug)
	if err != nil {
		return fmt.Errorf("error removing users from GitHub team: %v", err)
	}

	// Assign main maintainer to the GitHub team
	err = assigner.AssignUserToTeam(mainMaintainer, orgName, teamSlug, "maintainer")
	if err != nil {
		return fmt.Errorf("error assigning main maintainer: %v", err)
	}

	// Assign deputy maintainer to the GitHub team
	err = assigner.AssignUserToTeam(deputyMaintainer, orgName, teamSlug, "maintainer")
	if err != nil {
		return fmt.Errorf("error assigning deputy maintainer: %v", err)
	}

	return nil
}

// Creates a GitHub team and assigns it to an external group
func CreateTeamWithExternalGroup(orgName, teamSlug, externalGroupName string) error {
	// Create the GitHub team using the `gh api` command
	createTeamCommand := exec.Command("gh", "api", "-X", "POST", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "orgs/"+orgName+"/teams", "-f", "name="+teamSlug, "-f", "notification_setting=notifications_enabled", "-f", "privacy=closed")
	output, err := createTeamCommand.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating GitHub team: %v\n%s", err, output)
	}

	// Remove all users from the GitHub team
	err = remover.RemoveUsersFromTeam(orgName, teamSlug)
	if err != nil {
		return fmt.Errorf("error removing users from GitHub team: %v", err)
	}

	// Assign the team to the external group
	err = assigner.AssignTeamToExternalGroup(orgName, teamSlug, externalGroupName)
	if err != nil {
		return fmt.Errorf("error assigning team to external group: %v", err)
	}

	return nil
}

package assigner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/nestle/ghe-automation-kit/getter"
)

// Assigns a repository to a GitHub team
func AssignRepositoryToTeam(repoName, orgName, teamSlug, repoPermission string) error {
	cmd := exec.Command("gh", "api", fmt.Sprintf("orgs/%s/teams/%s/repos/%s/%s", orgName, teamSlug, orgName, repoName), "-X", "PUT", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "-f", fmt.Sprintf("permission=%s", repoPermission))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Assigns a GitHub team to an external group
func AssignTeamToExternalGroup(orgName, teamSlug, externalGroupName string) error {
	// Get the group ID based on the external group name
	groupID, err := getter.GetGroupID(orgName, externalGroupName)
	if err != nil {
		return fmt.Errorf("error getting group ID: %v", err)
	}

	// Execute the `gh api` command to assign the team to the external group
	cmd := exec.Command("gh", "api", "--method", "PATCH", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "orgs/"+orgName+"/teams/"+teamSlug+"/external-groups", "-F", "group_id="+groupID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error assigning team to external group: %v\n%s", err, output)
	}

	return nil
}

// Assigns a user to a team with the specified role
func AssignUserToTeam(userName, orgName, teamSlug, role string) error {
	// Add the user to the team with the specified role using GitHub API
	cmd := exec.Command("gh", "api",
		"--method", "PUT",
		"-H", "Accept: application/vnd.github+json",
		"-H", "X-GitHub-Api-Version: 2022-11-28",
		"/orgs/"+orgName+"/teams/"+teamSlug+"/memberships/"+userName,
		"-f", "role="+role,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Print the error message to the runner console
		fmt.Fprintf(os.Stderr, "failed to assign user to team as %s: %s\nOutput: %s\n", role, err, string(output))
		return err // Return the error to handle in the caller function
	}

	return nil
}

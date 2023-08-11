package remover

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Removes all users from a GitHub Team
func RemoveUsersFromTeam(orgName, teamSlug string) error {

	// Check if any members exist in the team
	teamMembersCommand := exec.Command("gh", "api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "orgs/"+orgName+"/teams/"+teamSlug+"/members")
	output, err := teamMembersCommand.CombinedOutput()
	if err != nil {
		// Print the error message to the runner console
		fmt.Fprintf(os.Stderr, "error retrieving team members: %s\nOutput: %s\n", err, string(output))
		return err // Return the error to handle in the caller function
	}

	// Parse the response to check if any members are present
	var members []map[string]interface{}
	err = json.Unmarshal(output, &members)
	if err != nil {
		// Print the error message to the runner console
		fmt.Fprintf(os.Stderr, "error parsing team members response: %v", err)
		return err // Return the error to handle in the caller function
	}

	if len(members) > 0 {
		// Remove all members from the team
		for _, member := range members {
			username, ok := member["login"].(string)
			if ok {
				removeMemberCommand := exec.Command("gh", "api", "-X", "DELETE", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "orgs/"+orgName+"/teams/"+teamSlug+"/memberships/"+username)
				output, err = removeMemberCommand.CombinedOutput()
				if err != nil {
					// Print the error message to the runner console
					fmt.Fprintf(os.Stderr, "error removing team member: %v\nOutput: %s\n", err, string(output))
					return err  // Return the error to handle in the caller function
				}
			}
		}
	}

	return nil
}

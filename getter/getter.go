package getter

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Repository struct {
	Name string `json:"name"`
}

func GetRepositories(orgName string, repoLimit string) ([]string, error) {
	cmd := exec.Command("gh", "repo", "list", orgName, "--json", "name", "--limit", repoLimit)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %v", err)
	}

	var repositories []Repository
	err = json.Unmarshal(output, &repositories)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %v", err)
	}

	repoNames := make([]string, len(repositories))
	for i, repo := range repositories {
		repoNames[i] = repo.Name
	}

	return repoNames, nil
}

// Represents an external group in the organization
type ExternalGroup struct {
	GroupID    int    `json:"group_id"`
	GroupName  string `json:"group_name"`
	UpdatedAt  string `json:"updated_at"`
}

// Retrieves the group ID based on the external group name
func GetGroupID(orgName, externalGroupName string) (string, error) {
	// Execute the `gh api` command to retrieve the list of external groups
	cmd := exec.Command("gh", "api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "orgs/"+orgName+"/external-groups")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error retrieving external groups: %v\n%s", err, output)
	}

	// Parse the response to find the group ID based on the external group name
	var response struct {
		Groups []ExternalGroup `json:"groups"`
	}
	err = json.Unmarshal(output, &response)
	if err != nil {
		return "", fmt.Errorf("error parsing external groups response: %v", err)
	}

	for _, group := range response.Groups {
		if group.GroupName == externalGroupName {
			return fmt.Sprintf("%d", group.GroupID), nil
		}
	}

	return "", fmt.Errorf("external group '%s' not found", externalGroupName)
}

// Gets the team ID based on the teamSlug and organization name
func GetTeamID(orgName, teamSlug string) (int, error) {
	cmd := exec.Command("gh", "api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", fmt.Sprintf("orgs/%s/teams/%s", orgName, teamSlug))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to get team ID: %w, Output: %s", err, string(output))
	}

	var teamData map[string]interface{}
	if err := json.Unmarshal(output, &teamData); err != nil {
		return 0, fmt.Errorf("failed to parse team data: %w", err)
	}

	teamID := int(teamData["id"].(float64))
	return teamID, nil
}

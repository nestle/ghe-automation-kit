package filter

func FilterReposByPrefix(repos []string, prefix string) []string {
	filteredRepos := make([]string, 0)

	for _, repo := range repos {
		if hasPrefix(repo, prefix) {
			filteredRepos = append(filteredRepos, repo)
		}
	}

	return filteredRepos
}

func hasPrefix(repo, prefix string) bool {
	if len(repo) < len(prefix) {
		return false
	}

	return repo[:len(prefix)] == prefix
}

# :gear: GitHub Enterprise Automation Kit

The GitHub Enterprise Automation Kit is a collection of Go packages designed to simplify and enhance automation tasks on GitHub Enterprise instances. These packages provide functionalities to interact with the GitHub API, manage repositories, automate workflows, and more.

## Packages

### `assigner`

The `assigner` package streamlines GitHub organization management by offering a straightforward way to assign repositories to teams or external groups, as well as assigning users to teams. Simplify your workflow and maintain efficient assignment processes within your projects.

```go
import "github.com/nestle/ghe-automation-kit/assign"
```

### `creator`

The `creator` package facilitates streamlined creation of resources, making it simple to generate repositories, teams, and other organizational components. Simplify your workflow by easily generating new entities with minimal code.

```go
import "github.com/nestle/ghe-automation-kit/creator"
```

### `filter`

The `filter` package provides a hassle-free solution for refining data sets. Effortlessly narrow down repositories, teams, or users based on specific criteria. Enhance your organization's management process with efficient data filtering.

```go
import "github.com/nestle/ghe-automation-kit/filter"
```

### `getter`

The `getter` package offers a straightforward approach to retrieve information from your GitHub organization. Quickly access repository details, team members, and other essential data. Retrieve what you need with ease using the Getter package.

```go
import "github.com/nestle/ghe-automation-kit/getter"
```

### `remover`

The `remover` package streamlines the removal of repositories, teams, and users from your GitHub organization. Simplify the process of eliminating unwanted components while maintaining control and organization within your projects.

```go
import "github.com/nestle/ghe-automation-kit/remover"
```

## Installation

To use the GitHub Enterprise Automation Kit in your Go project, you can include the required package imports in your code, as shown in the examples above.

Additionally, make sure you have Go 1.20 or later installed.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

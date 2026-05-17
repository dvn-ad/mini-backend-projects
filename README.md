# GitHub User Activity CLI

A simple command-line interface (CLI) to fetch and display the recent activity of a GitHub user.

This project is a solution to the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) challenge from [roadmap.sh](https://roadmap.sh).

## Features

- Fetches recent public events for any GitHub user.
- Displays activities such as:
  - Commits pushed (PushEvent)
  - Issues opened/closed (IssuesEvent)
  - Repositories starred (WatchEvent)
  - Repository/branch creation (CreateEvent)
- Simple and easy-to-use CLI interface.

## Prerequisites

- [Go](https://go.dev/doc/install) (version 1.16 or higher recommended)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/dvn-ad/GitHub-User-Activity.git
   cd GitHub-User-Activity
   ```

2. Build the application:
   ```bash
   go build -o github-activity
   ```

## Usage

Run the application by providing a GitHub username as an argument:

```bash
./github-activity <username>
```

### Example

```bash
./github-activity dvn-ad
```

**Output:**
```text
- Pushed 0 commits to dvn-ad/task-tracker-cli
- Created a new repository/branch in dvn-ad/task-tracker-cli
- Starred google/skills
```

## Implementation Details

- Uses the [GitHub Events API](https://docs.github.com/en/rest/activity/events) to fetch user data.
- Written in Go for efficiency and simplicity.
- Handles basic error cases like invalid usernames or API connection issues.

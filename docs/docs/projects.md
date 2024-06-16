# Projects

The `gitlabctl projects` command provides information about Gitlab projects (i.e. repositories) that your token has access to. This project information is a critical element of automating a variety of static code analysis workflows.

## Usage

```bash
gitlabctl projects --base-url https://gitlab.com/api/v4 --group-id <group id> --output json
```

## Help Text

```bash
$ gitlabctl projects -h
Enumerate Gitlab projects

Usage:
  gitlabctl projects [flags]

Flags:
      --archived          Include archived projects
      --group-id string   Group ID
  -h, --help              help for projects
      --mine              Include only projects owned by the authenticated user. (default true)

Global Flags:
      --base-url string      Base URL for Gitlab API. (e.g. https://gitlab.com/api/v4)
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
      --token string         Gitlab Access Token. Can also be set via GITLAB_TOKEN environment variable
  -v, --verbose              Verbose output
```

# Vulnerabilities

The `gitlabctl vulnerabilities` command allows you to leverage Gitlab's [application security](https://docs.gitlab.com/ee/user/application_security/) capabilities within your automation workflows, ensuring your team has visibility into the entire swathe of vulnerabilities that affect your applications.

## Usage

```bash
gitlabctl vulnerabilities --base-url https://gitlab.com/api/v4 --project <project id> --output json
```

## Help Text

```bash
$ gitlabctl vulnerabilities -h
Enumerate Gitlab vulnerabilities

Usage:
  gitlabctl vulnerabilities [flags]

Aliases:
  vulnerabilities, vulns

Flags:
  -h, --help                 help for vulnerabilities
      --project int          Project ID
      --severities strings   Vulnerability severities. Valid values are 'unknown', 'info', 'low', 'medium', 'high', 'critical'.
      --states strings       Vulnerability states. Valid values are 'detected', 'dismissed', 'resolved'. If no values are provided, 'detected' will be used by default.

Global Flags:
      --base-url string      Base URL for Gitlab API. (e.g. https://gitlab.com/api/v4)
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
      --token string         Gitlab Access Token. Can also be set via GITLAB_TOKEN environment variable
  -v, --verbose              Verbose output
```

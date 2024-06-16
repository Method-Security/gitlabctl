# Capabilities

gitlabctl integrates with the Gitlab API to provide security teams with insights into Gitlab primitives and security concepts, allowing teams to leverage the valuable data that Gitlab is generating about their projects. Each of the below pages offers an in depth look at a specific gitlabctl capability related to a different facet of Gitlab that can be enumerated.

- [Projects](./projects.md)
- [Vulnerabilities](./vulnerabilities.md)

## Top Level Flags

gitlabctl has several top level flags that can be used on any subcommand. These include:

```bash
Flags:
      --base-url string      Base URL for Gitlab API. (e.g. https://gitlab.com/api/v4)
  -h, --help                 help for gitlabctl
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
      --token string         Gitlab Access Token. Can also be set via GITLAB_TOKEN environment variable
  -v, --verbose              Verbose output
```

## Version Command

Run `gitlabctl version` to get the exact version information for your binary

## Output Formats

For more information on the various output formats that are supported by gitlabctl, see the [Output Formats](https://method-security.github.io/docs/output.html) page in our organization wide documentation.

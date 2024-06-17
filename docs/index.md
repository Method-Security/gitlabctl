# gitlabctl Documentation

Hello and welcome to the gitlabctl documentation. While we always want to provide the most comprehensive documentation possible, we thought you may find the below sections a helpful place to get started.

- The [Getting Started](./getting-started/basic-usage.md) section provides onboarding material
- The [Development](./development/setup.md) header is the best place to get started on developing on top of and with gitlabctl
- See the [Docs](./docs/index.md) section for a comprehensive rundown of gitlabctl capabilities

# About gitlabctl

gitlabctl offers security teams a way to tie Gitlab primitives into their security workflows, ensuring that they leverage those primitives within their security automation pipelines. Designed with data-modeling and data-integration needs in mind, gitlabctl can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The types of scans that gitlabctl can conduct are constantly growing. For the most up to date listing, please see the documentation [here](./docs/index.md)

To learn more about gitlabctl, please see the [Documentation site](https://method-security.github.io/gitlabctl/) for the most detailed information.

## Quick Start

### Get gitlabctl

For the full list of available installation options, please see the [Installation](./getting-started/installation.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/gitlabctl`
- `docker run ghcr.io/method-security/gitlabctl`
- Download the latest binary from the [Github Releases](https://github.com/Method-Security/gitlabctl/releases/latest) page
- [Installation documentation](./getting-started/installation.md)

### Authentication

gitlabctl takes advantage of Gitlab's [Personal Access Tokens](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) to authenticate to the Gitlab instance. This token will need to be passed to the gitlabctl command you wish to call, and can be handled in one of two ways.

1. Set a `GITLAB_TOKEN` environment variable with a value equal to the token's value
1. Pass the token in as an argument via the `--token` flag.
   1. Note that this is not recommended in production scenarios as you should avoid having your token logged in your command line history.

### General Usage

```bash
gitlabctl portscan <target>
```

#### Examples

```bash
gitlabctl projects --group-id <group> --mine false --base-url https://gitlab.com/api/v4
```

## Contributing

Interested in contributing to gitlabctl? Please see our organization wide [Contribution](https://method-security.github.io/community/contribute/discussions.html) page.

## Want More?

If you're looking for an easy way to tie gitlabctl into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, visit us [here](https://method.security)

## Community

gitlabctl is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](https://github.com/Method-Security) or our organization wide documentation [here](https://method-security.github.io).

Have an idea for a Tool to contribute? Open a Discussion [here](https://github.com/Method-Security/Method-Security.github.io/discussions).

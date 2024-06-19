<div align="center">
<h1>gitlabctl</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>

gitlabctl offers security teams a way to tie Gitlab primitives into their security workflows, ensuring that they leverage those primitives within their security automation pipelines. Designed with data-modeling and data-integration needs in mind, gitlabctl can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The types of scans that gitlabctl can conduct are constantly growing. For the most up to date listing, please see the documentation [here](https://method-security.github.io/gitlabctl/docs/index.html)

To learn more about gitlabctl, please see the [Documentation site](https://method-security.github.io/gitlabctl/) for the most detailed information.

## Quick Start

### Get gitlabctl

For the full list of available installation options, please see the [Installation](./getting-started/installation.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/gitlabctl`
- `docker run ghcr.io/method-security/gitlabctl`
- Download the latest binary from the [Github Releases](https://github.com/Method-Security/gitlabctl/releases/latest) page
- [Installation documentation](./getting-started/installation.md)

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

[verify]: https://github.com/Method-Security/gitlabctl/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/gitlabctl/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/gitlabctl
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/gitlabctl
[release]: https://github.com/Method-Security/gitlabctl/releases
[releases]: https://github.com/Method-Security/gitlabctl/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/gitlabctl.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/gitlabctl/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/gitlabctl?logo=docker&label=docker%20pulls%20%2F%20gitlabctl
[docker-pull]: https://hub.docker.com/r/methodsecurity/gitlabctl
[license]: https://github.com/Method-Security/gitlabctl/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg

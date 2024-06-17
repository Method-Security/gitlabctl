# Basic Usage

## Binaries

Running as a binary simplifies the process of providing authentication credentials to your gitlabctl command.

You can validate that the binary is working by generating a token for you Gitlab instance and running the following.

```bash
gitlabctl projects --base-url https://gitlab.com/api/v4 --group-id <group id> --output json
```

This will show you all the projects that you own in the provided Group ID.

## Docker

Running a gitlabctl command in Docker should work the same as when ran as a binary, you just need to provide the credential to the container. If you already have your `GITLAB_TOKEN` environment variable set in your local shell, you can then run.

```bash
docker run \
  -e GITLAB_TOKEN=$GITLAB_TOKEN \
  ghcr.io/method-security/gitlabctl \
  projects \
  --base-url https://gitlab.com/api/v4 \
  --group-id <group id> \
  --output json
```

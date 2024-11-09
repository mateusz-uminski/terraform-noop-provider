# terraform-noop-provider

[![verify](https://github.com/mateusz-uminski/terraform-noop-provider/actions/workflows/verify.yml/badge.svg)](https://github.com/mateusz-uminski/terraform-noop-provider/actions/workflows/verify.yml)

This is a minimal Terraform provider that allows users to create, read, update, and delete files in `/tmp` directory (or a custom directory if configured). It serves as an example to help understand provider development in Go.

Besides the README.md further documentation can be found in commits, code comments and nested README files.

Feel free to explore and copy everything you want. Enjoy!


# Usage

## Build the provider

```sh
make build
```

## Lint code

```sh
make lint
```

## Run example terraform code

```sh
make build

export REPO_DIR=$(git rev-parse --show-toplevel)
export TF_CLI_CONFIG_FILE="${REPO_DIR}/.terraformrc"

cd example/
terraform plan
```

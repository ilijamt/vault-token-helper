Vault Token Helper
------------------
[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/vault-token-helper)](https://goreportcard.com/report/github.com/ilijamt/vault-token-helper)
[![Codecov](https://img.shields.io/codecov/c/gh/ilijamt/vault-token-helper)](https://app.codecov.io/gh/ilijamt/vault-token-helper)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ilijamt/vault-token-helper)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilijamt/vault-token-helper)](LICENSE)
[![Release](https://img.shields.io/github/release/ilijamt/vault-token-helper.svg)](https://github.com/ilijamt/vault-token-helper/releases/latest)

A vault token helper that can be used to access multiple Vault instances based on VAULT_ADDR.

For more information about what this is about you can see [Token Helpers](https://www.vaultproject.io/docs/commands/token-helper).

## Install

### Pre-compiled binary

#### manually

Download the pre-compiled binaries from the [releases](https://github.com/ilijamt/vault-token-helper/releases) page.

#### homebrew

```bash
brew tap ilijamt/tap
brew install vault-token-helper
```


#### linux
```bash
mkdir -p ~/bin
export VTH_LATEST_TAG=$(basename $(curl -fs -o/dev/null -w %{redirect_url} https://github.com/ilijamt/vault-token-helper/releases/latest))
export OS_ARCH=$(uname -m)
curl -L https://github.com/ilijamt/vault-token-helper/releases/download/$VTH_LATEST_TAG/vault-token-helper_linux_$OS_ARCH.tar.gz --output - | tar xvz -C ~/bin vault-token-helper
```
Vault Token Helper
------------------

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
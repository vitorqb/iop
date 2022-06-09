# Improved OP

An improved (or simplified) version of the OnePassword (op) cli.

## Installation

```shell
rm -rf ~/.packages/iop-pkgbuild
mkdir -p ~/.packages
cd ~/.packages
git clone git@github.com:vitorqb/iop-pkgbuild.git
cd iop-pkgbuild
GIT_SSH="ssh-personal" GOPRIVATE="github.com/vitorqb/*" makepkg -si
```

## Configuring

A configuration file is expected in `~/.config/iop.yaml`.

Example:
```yaml
# Customize the dmenu command to run to query the user for selecting an item.
# Used for item names and accounts.
DmenuCommand: ["dmenu", "-i"]
```

## Usage (under dev)

```shell
# Copies a password to the clipboard
$ iop copy-password -n <NAME>
```

```shell
# List the name of known passwords
$ iop list-password
```

## Development

### Test

```sh
go test -v ./...
```

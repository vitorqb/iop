# Improved OP

An improved (or simplified) version of the OnePassword (op) cli.

For users like me who just want to have their passwords copied to the clipboard
without having to click buttons or run multiple commands.

## Install

### Installing on Archlinux

```shell
rm -rf ~/.cache/packages/iop && mkdir -p ~/.cache/packages
git clone 'https://github.com/vitorqb/iop-pkgbuild.git' ~/.cache/packages/iop
( cd ~/.cache/packages/iop && makepkg -si )
```

### Manual Install

Make sure `~/.local/bin` is in your `PATH`.

```shell
mkdir -p ~/.local/bin && cd ~/.local/bin
curl -L 'https://github.com/vitorqb/iop/releases/latest/download/iop.tar.gz' | tar -zx
```

## Install: Dependencies

### [dmenu](https://tools.suckless.org/dmenu/) (or alternatives)

In order to query an user to select an item (e.g. an account, a
password to copy, etc.), `iop` requires `dmenu` or a similar command
to be available. Look for `dmenu` in [Configuration](#configuration).

### [pinentry](https://www.gnupg.org/related_software/pinentry/index.html) (or similar)

In order to ask the user for a pin, `iop` requires `pinentry`, which
usually comes with gnupg. Look for `pinentry` in [Configuration](#configuration).

## Usage

```shell
# See help and usage
iop --help

# Set's the active account. Run this before anything else. Interactive.
iop select-account

# Copies a password to the clipboard (interactive)
iop copy-password

# Copies a password to the clipboard. Exact match by name or ID. (non-interactive).
iop copy-password -n Github
```

## Configuration

A configuration file is expected in `~/.config/iop.yaml`.

```yaml
# Customize the dmenu command used to query the user for selecting an item.
# Used for item names and accounts.
DmenuCommand: ["dmenu", "-i"]

# Customize the pinentry command used to query the user for a pin.
PinEntryCommand: ["pinentry-qt"]
```

## Development

### Run
```sh
./scripts/run.sh -h
./scripts/run.sh -- --help
./scripts/run.sh -- copy-password -n Foo
```

### Test
```sh
./scripts/test.sh -h
./scripts/test.sh
./scripts/test.sh -t TestSomethingSpecific
./scripts/test.sh -v
```

### Lint/Format
```sh
./scripts/format.sh -h
./scripts/format.sh
```

### Build
```sh
./scripts/build.sh -h
./scripts/build.sh
```

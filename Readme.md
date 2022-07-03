# Password Manager Wrapper

A wrapper arround password manager to make it easier to copy a 
password to your clipboard with minimal user interaction.

## Install

### Installing on Archlinux

```shell
rm -rf ~/.cache/packages/pmwrap && mkdir -p ~/.cache/packages
git clone 'https://github.com/vitorqb/pmwrap-pkgbuild.git' ~/.cache/packages/pmwrap
( cd ~/.cache/packages/pmwrap && makepkg -si )
```

### Manual Install

Make sure `~/.local/bin` is in your `PATH`.

```shell
mkdir -p ~/.local/bin && cd ~/.local/bin
curl -L 'https://github.com/vitorqb/pmwrap/releases/latest/download/pmwrap.tar.gz' | tar -zx
```

## Install: Dependencies

### [dmenu](https://tools.suckless.org/dmenu/) (or alternatives)

In order to query an user to select an item (e.g. an account, a
password to copy, etc.), `pmwrap` requires `dmenu` or a similar command
to be available. Look for `dmenu` in [Configuration](#configuration).

### [pinentry](https://www.gnupg.org/related_software/pinentry/index.html) (or similar)

In order to ask the user for a pin, `pmwrap` requires `pinentry`, which
usually comes with gnupg. Look for `pinentry` in [Configuration](#configuration).

### [libnotify](https://gitlab.gnome.org/GNOME/libnotify)

In order to send user notifications, we rely on the `notify-send`
command which comes with `libnotify`.

## Usage

```shell
# See help and usage
pmwrap --help

# Set's the active account. Run this before anything else. Interactive.
pmwrap select-account

# Copies a password to the clipboard (interactive)
pmwrap copy-password

# Copies a password to the clipboard. Exact match by name or ID. (non-interactive).
pmwrap copy-password -n Github
```

## Configuration

A configuration file is expected in `~/.config/pmwrap.yaml`.

```yaml
# Customize the dmenu command used to query the user for selecting an item.
# Used for item names and accounts.
DmenuCommand: ["dmenu", "-i"]

# Customize the pinentry command used to query the user for a pin.
PinEntryCommand: ["pinentry-qt"]

# Customize the program used for notify-send
NotifySendCommand: "libnotify"
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

# aget

`git clone` a package from AUR (the unofficial Arch Linux package repository that contains user-submitted packages), without having to remember the URL.

## Installation

One way is to install it from AUR, another way is:

```sh
sudo pacman -S base-devel git go --noconfirm --needed
go get -u github.com/xyproto/aget
sudo install -Dm755 ~/go/bin/aget /usr/bin/aget
```

An alternative to `go get` + `install` is to use `go install`:

```sh
go install github.com/xyproto/aget@latest
```

## Example use

### Download and install a package from AUR

First make sure that `base` and `base-devel` are installed.

Then download the `ld-lsb` package from AUR:

    aget ld-lsb

Build and install it with `makepkg`:

    makepkg -i

### Create a new AUR package

    aget newpackage

Your ssh key must be set up at the [AUR web page](https://aur.archlinux.org) first for this to work.

## General information

* Version: 1.3.3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
* License: BSD-3

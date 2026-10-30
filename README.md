# aget

`git clone` a package from AUR (the unofficial Arch Linux package repository that contains user-submitted packages), without having to remember the URL.

Your ssh key should be set up at the [AUR web page](https://aur.archlinux.org) first for this to work (the HTTPS fallback will be used if not set up).

## Installation

Install `go` and `git`, then:

```sh
go install github.com/xyproto/aget@latest
```

Then just run it as `~/go/bin/aget` OR
* set up an alias with `alias aget=~/go/bin/aget` OR
* add `~/go/bin` to the `PATH` OR
* install it to `/usr/bin` with ie. `sudo install -Dm755 ~/go/bin/aget /usr/bin/aget`.

## Example use

### Download and install a package from AUR

First make sure that `base` and `base-devel` are installed.

Then download the `ld-lsb` package from AUR (`ld-lsb` is just an example):

    aget ld-lsb

Build and install it with `makepkg`:

    makepkg -i

### Create a new AUR package

    aget newpackage

## General information

* Version: 1.4.3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
* License: BSD-3

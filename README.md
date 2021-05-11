# aget

`git clone` a package from AUR (the unofficial Arch Linux package repository that contains user-submitted packages), without having to remember the URL.

## Installation

One way is to install it from AUR, another way is:

```sh
sudo pacman -S go --noconfirm --needed
go get -u github.com/xyproto/aget
sudo install -Dm755 ~/go/bin/aget /usr/bin/aget
```

## Example use

Download the `ld-lsb` package from AUR:

`aget ld-lsb`

It can then be built and installed with `makepkg -i`.

## General information

* Version: 1.2.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
* License: BSD-3

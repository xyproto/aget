# aget

`git clone` a package from AUR (the unofficial Arch Linux package repository that contains user-submitted packages), without having to remember the URL.

## Installation

One way is to install it from AUR, another way is:

`sudo pacman -S go --noconfirm --needed && go get -u github.com/xyproto/aget`

## Example use

Download the `ld-lsb` package from AUR:

`aget ld-lsb`

It can then be built and installed with `makepkg -i`.

## General information

* Version: 1.1.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
* License: MIT

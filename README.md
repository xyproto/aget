aurtic
======

Possibly the fastest way to download and extract a source package from AUR.

Downloads from AUR4 if the package name ends with ".git".

Aurtic is written in Go and intended for use on Arch Linux.


Examples
--------

Download the ld-lsb package from AUR:

`aurtic ld-lsb`

Same thing, but will overwrite existing files:

`aurtic -f ld-lsb`


Installation
------------

One way is to install it from AUR, another way is:
`go get github.com/xyproto/aurtic`

General information
-------------------
* Version: 0.31
* Alexander F Rødseth <xyproto@archlinux.org>

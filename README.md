aurtic
======

Possibly the fastest way to download and extract a source package from AUR.

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

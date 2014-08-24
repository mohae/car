baller
======

creates tarballs out of things

Baller is a package to be used by other programs.

Baller exposes a function to make a tarball out of a variadic list of files or directories passed to it and save it to to a specified destination.

Functional requirements:
    Create an archive in specified format
    Save tarball to specified directory
    Delete archived directory


Patterns:
The filename can be in a pattern:

Supported wishlist:
Date
Datetime
Year(four year only)
month  (2 digit)
day    (2 digit)
Month  (full name)
Time   (ISO 3369)
Hour   24h
Minute 0:60
Second 0:60

Notes: 
* uses semantic versioning
* everything subject to change as this is version .0.0.1
* nothing probably works at the moment as this is version .0.0.1
* expect something to work and the functional requirements to be a bit more set when this hits version 0.1.0
* frásögn is the program to test this out. When it hits version .0.1.0, it will be stable to use. https://github.com/mohae/frasogn
* rancher will implement this package to replace its internal archive package, which baller was initially based on.
* there are no backwords compatibility guarantees until this hits v1.0.0

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

## Settings
Baller has sane defaults that allow it to be run without any additional configuration.

Baller will take the first setting at the highest order of precedence; whether an empty value is interpreted to be a setting override or a missing setting is determined on a per setting basis. As a general  rule, an empty value being interpreted as an override is the exception, not the norm.

The order of precedence for Baller settings, from highest to lowest, are:

* Command-line
* Environment variable
* Configuration file
* Application defaults

All possible settings will be documented in their relevant configuration file. Settings that do not override the application default will be commented out, with its value set to the application's default for that key.

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


Commands to be implemented:
Create, c creates an archive of the destination paths using the passed filename, compression type, and source paths. No deletion of source is done.
Extract, x extracts the archive in the specified path to the target path, if specified.  
Delete, d creates an archive of the destination paths using the passed filename, compression type, and source paths. If the archive is successfully created, the sources are deleted.

Flags to be implemented
-v, --verify true, t, false, f
-d, --datetime true, t, false, f
-m, --format string, Go's datetime format string.
-s, --separator , '-' by default, overrides default


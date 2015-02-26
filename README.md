car
===

Car was created so I could have a command-line tool for creating and extracting compressed archives on Windows.

## About 
Car is a cross platform tool for working with compressed archives. Car is not meant to be a replacement for `tar` and is not necessarily tape/archive focused. Where possible, flags have been made to be consistent with `tar`'s flags, but not all of `tar`'s flags will be supported. Some flags may not be implemented in the exact manner as `tar`'s, but consistency with `tar`, where possible, is the goal.

Even though `zip` is supported, it is minimally supported. Most options are for `tar` archives and compressed files only, not for `zip`.

Car also is a development tool for working on the underlying archiving package, [carchivum](https://github.com/mohae/carchivum).

## Compiling
At this point, no pre-compiled executable is available. One will need to have [Go](https://golang.org) and Git installed in order to compile Car.

Clone the repo:
    git clone https://github.com/mohae/car

Switch to the repo directory:

    cd /path/to/car/repo

Get the dependencies:

    go get -u

Compile

    go build

Move the executable to a location in your PATH.

## Usage
### Create an archive
To create an archive

    car create archive-name.tgz path/to/archive

Car, by default will create a gzip'd tar archive.

To create a zip archive:

    car create -format=zip archive-name.zip path/to/archive

### Extract an archive
Car will check the specified file to see what format it is in and extract the archive using the appropriate format. Currently, it assumes that any compressed files that aren't in the _zip_ format are tarballs, e.g. a file compressed with `gzip` contains a tarball.

The destination path is optional. If a destination is specified, the extracted files will be created within the specified destinatino directory, otherwise they will be extracted to the current working directory.

    car extract source.tgz [path/to/dest] 

### Operations

Most supported operations and their modifiers are not yet supported. These are documented for development purposes. Once `car` is further along, the status of commands and options support will be documented. Until then, assume nothing is working. 

Supported operations:
```
    extract (not implemented) extract files from an archive
    create                create a new archive

```
Operation modifiers:

```
create:
    --owner=UID             force UID as owner for added files
    --group=GID                 force GID as group for added files
    --mode=MASK              force MASK as mode for added files


NOT IMPLEMENTED
extract:
  -D, --remove-files              remove files after adding them to the archive
      --delete-files              alias to -D
  -k, --keep-old-files            don't replace existing files when extracting
      --keep-newer-files          don't replace existing files that are newer
                                than their archive copies
      --overwrite                 overwrite existing files when extracting
      --atime-preserve            don't change access time on dumped files
  -m, --modification-time         don't extract file modified time
      --same-owner                try extracting files with the same ownership
      --no-same-owner             extract files as yourself
      --numeric-owner             always use numbers for user/group names
  -p, --same-permissions          extract permissions information
      --no-same-permissions       do not extract permissions information
      --preserve-permissions      same as -p
```
  
Local file selection:
```
    --exclude=PATTERN           exclude files, given as a PATTERN
    --exclude-ext=[EXTENSIONS]  exclude files with EXTENSIONS
    --exclude-anchored          exclude patterns match file name start
    --include=PATTERN           include files, given as a PATTERN
    --include-ext=[EXTENSIONS]  include files with EXTENSIONS
    --include-anchored          include patterns match file name start
-M, --newer-mtime=DATE          only stores files modified since DATE

NOT IMPLEMENTED:
    --wildcards                 patterns use wildcards
    --no-wildcards              patterns do not use wildcards
    --newer-file=FILENAME       only store files newer than the DATE for
                                FILENAME
```

## Status
Create command for creating a compressed archvie file from a list of sources is working:
  * suppports gzip/tar
  * supports zip
  * supports fullpath/relative path
  * supports logging

Additional functionality will be added as [Carchivum](https://github.com/mohae/carchivum) functionality expands.
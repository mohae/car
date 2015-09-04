car
===

Car was created so I could have a cross platform command-line tool for creating and extracting compressed archives. I use it on various Linux distros and Windows.

## About 
Car is a cross platform tool for working with compressed archives. While Car's default archive format is tar, Car is not meant to be a replacement for `tar` and is not necessarily tape/archive focused. Currently, `car` archives compressed with either `gzip` or `bzip2` are compatible with `tar`. Car can also use other compression formats and does not support all of the formats that `tar` does. 

Car can also create and extract `zip` archives. When creating zip archives, other compression formats are not supported.

Car also is a development tool for working on the underlying archiving package, [carchivum](https://github.com/mohae/carchivum).

Car supports a some flags but it supporting all the flags that `tar` does is not a goal. Additional flags may be added in the future.

## Compiling
### Make sure Go is installed
At this point, no pre-compiled executable is available. One will need to have [Go](https://golang.org) and Git installed in order to compile Car. To install Go, follow the [official install instructions](https://golang.org/doc/install.html). Installing Go using either `brew`, for Mac users, or your package manager, for Linux users, may or may not work and is not recommended.

If you are on Linux, Go 1.5 can be installed using my [Go 1.5 install script](https://gist.github.com/mohae/7e738af18e5041ac3fc4). At time of writing, this is the current version of Go. If there are more recent releases either check [my Gists](https://gist.github.com/mohae) for a version of the script that corresponds with that release or modify your copy of the linked script to download the updated version.

### Clone, get debendencies, and compile
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

### Commands and flag names may change
Please note the `extract` and `create` commands may change until Car 1.0 is released.

### Extract an archive
Car will check the specified file to see what format it is in and extract the archive using the appropriate format. Currently, it assumes that any compressed files that aren't in the _zip_ format are tarballs, e.g. a file compressed with `gzip` contains a tarball.

The destination path is optional. If a destination is specified, the extracted files will be created within the specified destinatino directory, otherwise they will be extracted to the current working directory.

    car extract source.tgz [path/to/dest] 

### Operations
Most supported operations and their modifiers are not yet supported. These are documented for development purposes. Once `car` is further along, the status of commands and options support will be documented. Until then, assume nothing is working. 

Supported operations:

    extract               extract files from an archive
    create                create a new archive

Operation modifiers:

__create:__
  
Local file selection:
```  --exclude-ext=[EXTENSIONS]  exclude files with EXTENSIONS
    --exclude-anchored          exclude patterns match file name start
    --include=PATTERN           include files, given as a PATTERN
    --include-ext=[EXTENSIONS]  include files with EXTENSIONS
    --include-anchored          include patterns match file name start
```
extract:
TODO document flags


__Not Implemented:__
--exclude=PATTERN           exclude files, given as a PATTERN
    
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
  -M, --newer-mtime=DATE          only stores files modified since DATE

    --wildcards                 patterns use wildcards
    --no-wildcards              patterns do not use wildcards
    --newer-file=FILENAME       only store files newer than the DATE for
                                FILENAME

## Status
Currently supported functionality:
  * tar archives, optionally compressed with:
    * gzip (default)
    * lz4
    * bzip2 (extract only)
  * zip archives

Additional functionality will be added as [Carchivum](https://github.com/mohae/carchivum) functionality expands.

## Notes:  
* Currently, directories are created with `0744` permission. There may be support for properly setting directory permissions.
* Symlinks are not followed/archived.
* Hidden files and directories are not supported.

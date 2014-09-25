car
===

Car is a tool for working with compressed archives.

Tarballs are great, but they are not compressed by default, which is why the `c` flag is used with the `tar` command. 

Zip files are great, for Windows users, they are combressed archives by default.

Car understands both, and automatically compresses tarballs, using its default compression format, or specify one of the `car` supported formats.

Car is meant to be a simple tool, which helps me develope other things, including [Quine](https://github.com/mohae/quine), my cli application template, [Contour](https://github.com/mohae/contour), my configuration management package, [my cli fork](https://github.com/mohae/cli) of [mitchellh's cli package](https://github.com/mitchellh/cli), which allows me to work on some additional functionality before proposing the changes on Mitchell's CLI package.

Car also is a development tool for working on the underlying archiving packages.

Lastly, `car` will be cross-platform in the near future, meaning I can have a command-line tool for creating archives and compressing things on Windows.

## Car Options
Car implements POSIX style options, seeking to be consistent with `tar`, when possible. Car does not support everything `tar` does and may support options that `tar` does not support. For `car` operations that are consistent with `tar`, the `tar` options are implemented, at minimum; e.g., `-c` and `--create` for creating a new archive.

Car may implement additional options or additional aliases to existing options,

### Operations

Most supported operations and their modifiers are not yet supported. These are documented for development purposes. Once `car` is further along, the status of commands and options support will be documented. Until then, assume nothing is working. 

Supported operations:
```
    -x, --extract, --get        extract files from an archive
    -c, --create                create a new archive

```
Operation modifiers:

```
-D, --remove-files              remove files after adding them to the archive
    --delete-files              alias to -D
-k, --keep-old-files            don't replace existing files when extracting
    --keep-newer-files          don't replace existing files that are newer
                                than their archive copies
    --overwrite                 overwrite existing files when extracting
    --owner=NAME                force NAME as owner for added files
    --group=NAME                force NAME as group for added files
    --mode=CHANGES              force (symboloc) mode CHANGES for added files
    --atime-preserve            don't change access time on dumped files
-m, --modification-time         don't extract file modified time
    --same-owner                try extracting files with the same ownership
    --no-same-owner             extract files as yourself
    --numeric-owner             always use numbers for user/group names
-p, --same-permissions          extract permissions information
    --no-same-permissions       do not extract permissions information
    --preserve-permissions      same as -p
```
Archive format selection:
```
-z, --gzip, --ungzip            filter the archive through gzip
```
Local file selection:
```
-C, --directory=DIR             change to directory DIR
    --exclude=PATTERN           exclude files, given as a PATTERN
    --exclude-ext=[EXTENSIONS]  exclude files with EXTENSIONS
    --exclude-anchored          exclude patterns match file name start
    --include=PATTERN           include files, given as a PATTERN
    --include-ext=[EXTENSIONS]  include files with EXTENSIONS
    --include-anchored          include patterns match file name start
    --wildcards                 patterns use wildcards
    --no-wildcards              patterns do not use wildcards
-N, --newer=DATE-OR-FILE        only store files newer than DATE or FILE
-M, --newer-mtime=DATE          only stores files modified since DATE
    --after-date=DATE           same as -N
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

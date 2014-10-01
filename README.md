car
===

Car is a tool for working with compressed archives. Car is not meant to be a replacement for `tar` and is not necessarily tape/archive focused. Where possible, flags have been made to be consistent with `tar`'s flags, but not all of `tar`'s flags will be supported. Some flags may not be implemented in the exact manner as `tar`'s, but consistency with `tar`, where possible, is the goal.

Even though `zip` is supported, it is minimally supported. Most options are for `tar` archives and compressed files only, not for `zip`.

Car is meant to be a simple tool, which helps me develope other things, including [Quine](https://github.com/mohae/quine), my cli application template, [Contour](https://github.com/mohae/contour), my configuration management package, [my cli fork](https://github.com/mohae/cli) of [mitchellh's cli package](https://github.com/mitchellh/cli), which allows me to work on some additional functionality before proposing the changes on Mitchell's CLI package. Due to Car's flag requirements, I have a fork of [ogier's pflage](https://github.com/ogier/pflag), [mohae's pflag](https://github.com/mohae/pflag) that adds support for time flags.

Car also is a development tool for working on the underlying archiving package.

Lastly, `car` will be cross-platform in the near future, meaning I can have a command-line tool for creating archives and compressing things on Windows.

## Car Options
Car implements POSIX style options, seeking to be consistent with `tar`, when possible. Car does not support everything `tar` does and may support options that `tar` does not support. For `car` operations that are consistent with `tar`, the `tar` options are implemented, at minimum; e.g., `-c` and `--create` for creating a new archive.

Car may implement additional options or additional aliases to existing options,

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

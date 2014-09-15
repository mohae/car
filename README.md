car
===

Car is a tool for working with compressed archives.

Tarballs are great, but they are not compressed by default, which is why the `c` flag is used with the `tar` command. 

Zip files are great, for Windows users, they are combressed archives by default.

Car understands both, and automatically compresses tarballs, using its default compression format, or specify one of the `car` supported formats.

Car is meant to be a simple tool, which helps me develope other things, including [Quine](https://github.com/mohae/quine), my cli application template, [Contour](https://github.com/mohae/contour), my configuration management package, [my cli fork](https://github.com/mohae/cli) of [mitchellh's cli package](https://github.com/mitchellh/cli), which allows me to work on some additional functionality before proposing the changes on Mitchell's CLI package.

Car also is a development tool for working on the underlying archiving packages.

Lastly, `car` will be cross-platform in the near future, meaning I can have a command-line tool for creating archives and compressing things on Windows.

## Status
Not working, under development.

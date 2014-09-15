<<<<<<< HEAD
Quine
=====

#### README
is under development. Please excuse the mess, sentence fragments, abandoned thoughts, etc.

#### Quine mostly works
When compiled, Quine supports `version` and `help` both as commands and flags. `hello` is supported as a command with one additional bool flag, `lower`, which lowercases the output instead of preserving case. In addition to the `hello` output, Quine also prints out sample log message, some of which will end up in the application log and all of which will be written to stdout. This is just to show logging.

Quine has not been implemented in any applications, yet, and I will not consider it stable until it has as some bugs and annoyances will appear in the process.

## About
Bobby Quine knows software.

Quine is an idiosyncratic CLI application template that seeks to make it easy to integrate your application code into a CLI interface with support for flags, configuration files, and logging. It was created because I found the other options to involve more work than necessary, imo, to get a CLI application started. 

Quine is not designed to be used for server applications where the application must be able to recover from a panic and continue running in most situations. It does not support immediate logging on start-up like one might want from a server application that is to be continuously running as it delays any logging evaluation until the command-line args are processed, in case there are any that affect logging (if this isn't applicable to your program, moving logging to a point where it starts immediately is simple.

Quine was created to minimize the changes needed to create a new CLI application. Mostly, all you need to do is clone https:\\github.com/mohae/quine as your new application, register your application's settings, add the commands, and add the command handlers to the `yourApp/app/` directory. 

No adding of flag filtering or other argument handling, registering a setting as a flag automatically takes care of that.

Quine is both a CLI application template and a minimalist example application. Quine can be compiled and run, with all flags working. To customize to your app, replace the call to ?????? with your application's main package name.

It supports:
* Go Flags
* sub-command aliases
* application defaults
* application configuration file in TOML or JSON
* environment variable support
* logging, defaults to off

### CLI
Quine uses a fork of [Mitchell Hashimoto's cli](https://github.com/mitchellh/cli). It is mostly consistent with the parent repo, with a few customizations for additional features.

### Logging
Quine uses [seelog](https://github.com/cihub/seelog) for its logging library. With it comes colorized stdout output, custom log levels for various logger outputs, custom loggers, custom filters, email alerts, etc. Quine's log configuration file is in `seelog.xml`

### Contour
Quine uses [contour](https://github.com/mohae/contour) for its configuration management. Contour supports multiple named configurations, application settings, configuration files, and flags. Contour automatically handles flag filtering, just pass it all of the received args and it will return the filtered arg list. The flag data is automatically applied to the configuration.


## To Use

Clone the repo:

    git clone https://github.com/mohae/quine yourApp

And change into it:

    $ cd yourApp

## Wishlist
* Environment variable support
* An application to automate the manual changes needed to turn Quine into yourApp
=======
car
======

car (Compressed Archive) is a tool to create compressed archives out of things.

The compressed archives an either be in zip, or tape archive format, tar. The zip format includes compression. Car uses the `.zip` extension for zip archives. Tar does not include compression so car automatically compresses the tarball,  using gzip by default. The compression scheme can be configured or set at runtime. By default, car uses the `.car` extension for the compressed tarballs. This is because car may support compression formats that tar does not. Tar also supports formats that car does not. If you want to use the tar derived extensions, that can be set, which will result in compressed archives that use the tar format and that have a tar compatible compression format will use the tar version of the extension. If the compressoin format used is not compatible with tar, `.car` will be used.

Patterns:
The filename can be in a pattern:

## Settings
Car has sane defaults that allow it to be run without any additional configuration. These built-in defaults can be overridden by the application defaults, which are set in `car.toml` in the application's root directory. The application level settings can be overridden by environment variables. These settings can then be overridden at run time, either by run templates or via the command-line.

Car will take the first setting at the highest order of precedence; whether an empty value is interpreted to be a setting override or a missing setting is determined on a per setting basis. As a general  rule, an empty value being interpreted as an override is the exception, not the norm.

The order of precedence for Baller settings, from highest to lowest, are:

* Command-line/car run template
* Environment variable
* Configuration file
* Application defaults

All possible settings will be documented in their relevant configuration file. Settings that do not override the application default will be commented out, with its value set to the application's default for that key.

Notes: 
* uses semantic versioning
* everything subject to change as this is version .0.0.1
* nothing probably works at the moment as this is version .0.0.1
* expect something to work and the functional requirements to be a bit more set when this hits version 0.1.0
* frásögn is the program to test this out. When it hits version .0.1.0, it will be stable to use. https://github.com/mohae/frasogn
* rancher will implement this package to replace its internal archive package, which car was initially based on.
* there are no backwords compatibility guarantees until this hits v1.0.0


Commands to be implemented:
Create, c creates an archive of the destination paths using the passed filename, compression type, and source paths. No deletion of source is done.
Extract, x extracts the archive in the specified path to the target path, if specified.  
Delete, d creates an archive of the destination paths using the passed filename, compression type, and source paths. If the archive is successfully created, the sources are deleted.

### Flags to be implemented. 
Please note, if a default value is set, that is car's default.
-v, --verify, true by default, true, t, false, f
-d, --datetime, false by default,  true, t, false, f
-f, --format string, RCF922 by default, Go's datetime format string.
-l, --logging  false by default, true, t, false, f, 
-s, --separator, '-' by default, overrides default
-c, --compression, gzip by default, overrides defailt.

>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7

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

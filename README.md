Quine
=====

## In Development

Bobby Quine knows software.

Quine is an idiosyncratic CLI application template that seeks to make it easy to integrate your app into a CLI. It is based on the assumption that your app already knows about its configuration so it doesn't seek to have you reimplement it.

It does, however, make some assumptions about your application that probably aren't true if it's not designed for Quine. The requirements are minimal. Quine assumes that your application is able to tell it its name, version and the sub-commands it implements including flags, arguments, and other options. It also assumes that your application is reading its configuration from environment variables. Quine also assumes that your application is structured in a manner that there is a main application package, with which it will obtain its information from. 

How you architect your application beyond that single requirement is up to you.

This assumption makes it easy to add a CLI to a package, including full logging capabilities. This can help the development process for certain workflows.

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
Quine uses [seelog](http://github.com/cihub/seelog) for its logging library. With it comes colorized stdout output, custom log levels for various logger outputs, custom loggers, custom filters, email alerts, etc. Quine's log configuration file is in `seelog.xml`

## To Use

Clone the repo:

    git clone https://github.com/mohae/quine

And change into it:

    $ cd quine

### Bobby

Bobby Newmark is a hacker.

He is used to turn a `quine` template into your application's CLI harness. When run, `bobby` takes the parent directorie's name and uses that as the application's name. This value is used to replace `quine` with your application's name.

## To Use

To just rename the files:

    $ ./bobby rename

After this has successfully run, clone your application into ---------

To add a basic application directory, add flags

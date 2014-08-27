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


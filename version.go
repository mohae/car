// From github.com/hashicorp/serf/blob/master/version.go
package main

// The git commit that was compiled. This will be filled in by the compiler
var GitCommit string

// The main version number that is being run at the moment.
<<<<<<< HEAD
const Version = "0.1.2"
=======
const Version = "0.0.1"
>>>>>>> 2e926bf1982fb46d62edad497bf67f5e2bff12e7

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = "dev"

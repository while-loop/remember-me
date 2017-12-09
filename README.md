Remember Me
==========

<p align="center">
  <img src="https://cdn.meme.am/instances/48462683.jpg">
  <br><br><br>
  <a href="https://godoc.org/github.com/while-loop/remember-me"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  <a href="https://travis-ci.org/while-loop/remember-me"><img src="https://img.shields.io/travis/while-loop/remember-me.svg?style=flat-square"></a>
  <a href="https://github.com/while-loop/remember-me/releases"><img src="https://img.shields.io/github/release/while-loop/remember-me.svg?style=flat-square"></a>
  <a href="https://coveralls.io/github/while-loop/remember-me"><img src="https://img.shields.io/coveralls/while-loop/remember-me.svg?style=flat-square"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-AGPLv3-blue.svg?style=flat-square"></a>
</p>

Automatic password manager

Remember Me is a proactive solution to keep passwords secure and fresh.
With the push for Autofill services in browsers, password managers, and
even [native mobile OS support](https://developer.android.com/guide/topics/text/autofill.html),
the need to `remember` passwords becomes more obsolete.

Why Should I Use Remember Me
----------------------------

- Proactively password resets
- We do not store __any__ passwords in databases
- Take advantage of Autofill services
- Having too many dag nabbing emails & password combinations!

Update passwords from a given password manager solution at given
intervals.

Dependencies
------------

- [protoc](https://github.com/google/protobuf/releases) >= 3.4.0
- [proto-gen-go](https://github.com/golang/protpbuf)
```go get github.com/golang/protobuf/protoc-gen-go```

Installation
------------

```
$ go get github.com/while-loop/remember-me/remme/...
```

Usage
-----

#### gRPC & Package
[Example](test/main.go)

#### cli
```bash
$ remme -h
```

```
NAME:
   remember-me - Automatic password changer

USAGE:
   remme [global options] command [command options] [arguments...]

VERSION:
   0.0.1

DESCRIPTION:
   Automatic password changer

COMMANDS:
     change, ch  change passwords for a given manager
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Changelog
---------

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/).

[CHANGELOG.md](CHANGELOG.md)

License
-------
rememeber-me is licensed under the GNU Affero General Public License.
See [LICENSE](LICENSE) for details.

Author
------

Anthony Alves

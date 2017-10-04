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
the need to `remember` passwords become more obsolete.

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

Building From Source
------------

Compile protobuf objects
```
$ go get -u github.com/while-loop/remember-me/
$ cd $GOPATH/src/github.com/while-loop/remember-me/
$ go generate ./...
```

Installation
------------

Install binaries
```bash
$ go get -u github.com/while-loop/remember-me/cmd/...
```

Usage
-----

#### gRPC & Package
[Example](test/main.go)

#### cli
```bash
$ remme -h
```

Changelog
---------

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/).

[CHANGELOG.md](CHANGELOG.md)

License
-------
rememeber-me is licensed under the GNU Affero General Public License.
See LICENSE for details.

Author
------

Anthony Alves

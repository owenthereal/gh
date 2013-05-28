# g + h = github

[![Build Status](https://drone.io/github.com/jingweno/gh/status.png)](https://drone.io/github.com/jingweno/gh/latest)

Fast GitHub command line client.

## Overview

gh is a command line client to GitHub. It's designed to run as fast as possible with easy installation across operating systems.
If you like gh, please also take a look at [hub](https://github.com/defunkt/hub). Hub is a reference implementation to gh.

## Motivation

**Fast** 

    $ time hub version > /dev/null
    hub version > /dev/null  0.03s user 0.01s system 93% cpu 0.047 total

    $ time gh version > /dev/null
    gh version > /dev/null  0.01s user 0.01s system 85% cpu 0.022 total

**Muti-platforms**

gh is fully implemented in the Go language with no external dependencies and is designed to run across operating systems.

**Easy installation**

There're no pre-requirements to run gh. Download the binary and go!

## Installation

There are [compiled binary forms of gh](https://drone.io/github.com/jingweno/gh/files) for Darwin, Linux and Windows.

To install gh on OSX with [Homebrew](https://github.com/mxcl/homebrew), you need to run:

    $ brew install https://raw.github.com/jingweno/gh/master/homebrew/gh.rb

To build and install gh locally, you need to have a [Go development environment](http://golang.org/doc/install) and run:

    $ go get github.com/jingweno/gh

## Usage
    
    $ gh help
    Usage: gh [command] [options] [arguments]

    Commands:

        pull-request      Open a pull request on GitHub
        help              Show help
        version           Show gh version

    See 'gh help [command]' for more information about a command.

## Roadmap

* authentication (done)
* gh pull-request (done)
* gh ci-status (in progress)
* gh clone
* gh remote add
* gh fetch
* gh cherry-pick
* gh am, hub apply
* gh fork
* gh check
* gh merge
* gh create
* gh init
* gh push
* gh browse
* gh compare
* gh submodule

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

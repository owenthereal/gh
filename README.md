# gh

[![Build Status](https://drone.io/github.com/jingweno/gh/status.png)](https://drone.io/github.com/jingweno/gh/latest)

Fast GitHub command line client.

## Overview

gh is a command line client to GitHub similar to [hub](https://github.com/defunkt/hub), but designed to be as fast as possible.

## Motivation

**Fast** 

    $ time hub version > /dev/null
    hub version > /dev/null  0.03s user 0.01s system 93% cpu 0.047 total

    $ time gh version > /dev/null
    gh version > /dev/null  0.01s user 0.01s system 85% cpu 0.022 total

**Muti-platforms**

gh is fully implemented in the Go language with no external dependencies and is designed to run across operating systems.

## Installation

There are [compiled binary forms of gh](https://drone.io/github.com/jingweno/gh/files) for Darwin, Linux and Windows.
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

Implementing all features from [hub](https://github.com/defunkt/hub):

* authentication (done)
* hub pull-request (done)
* hub ci-status (in progress)
* hub clone
* hub remote add
* hub fetch
* hub cherry-pick
* hub am, hub apply
* hub fork
* hub check
* hub merge
* hub create
* hub init
* hub push
* hub browse
* hub compare
* hub submodule

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


# g + h = github

[![Build Status](https://drone.io/github.com/jingweno/gh/status.png)](https://drone.io/github.com/jingweno/gh/latest)

Fast GitHub command line client. Current version is [0.5.0](https://drone.io/github.com/jingweno/gh/files).

## Overview

gh is a command line client to GitHub. It's designed to run as fast as possible with easy installation across operating systems.
If you like gh, please also take a look at [hub](https://github.com/defunkt/hub). Hub is a reference implementation to gh.

## Motivation

**Fast** 

    $ time hub version > /dev/null
    hub version > /dev/null  0.03s user 0.01s system 93% cpu 0.047 total

    $ time gh version > /dev/null
    gh version > /dev/null  0.01s user 0.01s system 85% cpu 0.022 total

    $ time hub browse > /dev/null
    hub browse > /dev/null  0.07s user 0.04s system 87% cpu 0.130 total

    $ time gh browse > /dev/null
    gh browse > /dev/null  0.03s user 0.02s system 87% cpu 0.059 total

**Muti-platforms**

gh is fully implemented in the Go language and is designed to run across operating systems.

**Easy installation**

There're no pre-requirements to run gh. Download the [binary](https://drone.io/github.com/jingweno/gh/files) and go!

**Unix**

gh commands are single, unhyphenated words that map to their Unix ancestors’ names and flags where applicable.

## Installation

There are [compiled binary forms of gh](https://drone.io/github.com/jingweno/gh/files) for Darwin, Linux and Windows.

To install gh on OSX with [Homebrew](https://github.com/mxcl/homebrew), run:

    $ brew install https://raw.github.com/jingweno/gh/master/homebrew/gh.rb

## Compilation

To compile gh from source, you need to have a [Go development environment](http://golang.org/doc/install) and run:

    $ go get github.com/jingweno/gh

Note that `go get` will pull down sources from various VCS.
Please make sure you have git and hg installed.

## Upgrade

Since gh is under heavy development, we roll out new releases often.
Please take a look at our [CI server](https://drone.io/github.com/jingweno/gh/files) for the latest built binaries.
We plan to implement automatic upgrade in the future.

To upgrade gh on OSX with Homebrew, run:

    $ brew upgrade https://raw.github.com/jingweno/gh/master/homebrew/gh.rb

To upgrade gh from source, run:

    $ go get -u github.com/jingweno/gh

## Usage
    
    $ gh help
    Usage: gh [command] [options] [arguments]

    Commands:

        pull              Open a pull request on GitHub
        ci                Show CI status of a commit
        browse            Open a GitHub page in the default browser
        compare           Open a compare page on GitHub
        help              Show help
        version           Show gh version

    See 'gh help [command]' for more information about a command.

## Roadmap

* authentication (done)
* gh pull-request (done)
* gh ci-status (done)
* gh browse (done)
* gh compare (done)
* gh fork (in progress)
* gh clone
* gh remote add
* gh fetch
* gh cherry-pick
* gh am, gh apply
* gh check
* gh merge
* gh create
* gh init
* gh push
* gh submodule

## License

gh is released under the MIT license. See LICENSE.md.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

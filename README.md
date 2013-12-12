# g + h = github [![Build Status](https://travis-ci.org/jingweno/gh.png?branch=master)](https://travis-ci.org/jingweno/gh)

![gh](http://owenou.com/gh/images/gangnamtocat.png)

Fast GitHub command line client implemented in Go. Current version is [v0.25.1](https://github.com/jingweno/gh/releases/tag/v0.25.1).

## Overview

gh is a command line client to GitHub. It's designed to run as fast as possible with easy installation across operating systems.
If you like gh, please also take a look at [hub](https://github.com/github/hub). Hub is a reference implementation to gh.

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

There're no pre-requirements to install gh. Download the [binary](https://github.com/jingweno/gh/releases) and go!

## Installation

### Homebrew

The easiest way to install `gh` on OSX is through [Homebrew](https://github.com/mxcl/homebrew).
You can add the [gh Homebrew repository](https://github.com/jingweno/homebrew-gh) with [`brew tap`](https://github.com/mxcl/homebrew/wiki/brew-tap):

    $ brew tap jingweno/gh
    $ brew install gh

### Standalone

`gh` can be easily installed as an executable.
Download the latest [compiled binary forms of gh](https://github.com/jingweno/gh/releases) for Darwin, Linux and Windows.

### Boxen

If you're using [boxen](http://boxen.github.com/), there's a [puppet-gh](https://github.com/boxen/puppet-gh) module to install and set up `gh`.

### Source

To compile gh from source, you need to have a [Go development environment](http://golang.org/doc/install), version 1.1 or better, and run:

    $ go get github.com/jingweno/gh

Note that `go get` will pull down sources from various VCS.
Please make sure you have git and hg installed.

## Upgrade

Since gh is under heavy development, I roll out new releases often.
Please take a look at the [built binaries](https://github.com/jingweno/gh/releases) for the latest built binaries.
I plan to implement automatic upgrade in the future.

### Homebrew

If you installed `gh` with `brew tap jingweno/gh`, you can upgrade it with:

    $ brew upgrade gh

### Source

To upgrade gh from source, run:

    $ go get -u github.com/jingweno/gh

## Aliasing

It's best to use `gh` by aliasing it to `git`.
All git commands will still work with `gh` adding some sugar.

`gh alias` displays instructions for the current shell. With the `-s` flag,
it outputs a script suitable for `eval`.

You should place this command in your `.bash_profile` or other startup
script:

    eval "$(gh alias -s)"

For more details, run `gh help alias`.

## Usage

### gh help
    
    $ gh help
    [display help for all commands]
    $ gh help pull-request
    [display help for pull-request]

### gh init

    $ gh init -g
    > git init
    > git remote add origin git@github.com:YOUR_USER/REPO.git

### gh push

    $ gh push origin,staging,qa bert_timeout
    > git push origin bert_timeout
    > git push staging bert_timeout
    > git push qa bert_timeout

    $ gh push origin
    > git push origin HEAD

### gh checkout

    $ gh checkout https://github.com/jingweno/gh/pull/35
    > git remote add -f -t feature git://github:com/foo/gh.git
    > git checkout --track -B foo-feature foo/feature

    $ gh checkout https://github.com/jingweno/gh/pull/35 custom-branch-name

### gh merge

    $ gh merge https://github.com/jingweno/gh/pull/73
    > git fetch git://github.com/jingweno/gh.git +refs/heads/feature:refs/remotes/jingweno/feature
    > git merge jingweno/feature --no-ff -m 'Merge pull request #73 from jingweno/feature...'

### gh clone

    $ gh clone jingweno/gh
    > git clone git://github.com/jingweno/gh

    $ gh clone -p jingweno/gh
    > git clone git@github.com:jingweno/gh.git

    $ gh clone jekyll_and_hype
    > git clone git://github.com/YOUR_LOGIN/jekyll_and_hype.

    $ gh clone -p jekyll_and_hype
    > git clone git@github.com:YOUR_LOGIN/jekyll_and_hype.git

### gh fetch

    $ gh fetch jingweno
    > git remote add jingweno git://github.com/jingweno/REPO.git
    > git fetch jingweno

    $ gh fetch jingweno,foo
    > git remote add jingweno ...
    > git remote add foo ...
    > git fetch --multiple jingweno foo

    $ gh fetch --multiple jingweno foo
    > git remote add jingweno ...
    > git remote add foo ...
    > git fetch --multiple jingweno foo

### gh cherry-pick

    $ gh cherry-pick https://github.com/jingweno/gh/commit/a319d88#comments
    > git remote add -f jingweno git://github.com/jingweno/gh.git
    > git cherry-pick a319d88

    $ gh cherry-pick jingweno@a319d88
    > git remote add -f jingweno git://github.com/jingweno/gh.git
    > git cherry-pick a319d88

    $ gh cherry-pick jingweno@SHA
    > git fetch jingweno
    > git cherry-pick SHA

### gh remote

    $ gh remote add jingweno
    > git remote add -f jingweno git://github.com/jingweno/CURRENT_REPO.git

    $ gh remote add -p jingweno
    > git remote add -f jingweno git@github.com:jingweno/CURRENT_REPO.git

    $ gh remote add origin
    > git remote add -f YOUR_USER git://github.com/YOUR_USER/CURRENT_REPO.git    

### gh submodule

    $ gh submodule add jingweno/gh vendor/gh
    > git submodule add git://github.com/jingweno/gh.git vendor/gh

    $ gh submodule add -p jingweno/gh vendor/gh
    > git submodule add git@github.com:jingweno/gh.git vendor/gh

    $ gh submodule add -b gh --name gh jingweno/gh vendor/gh
    > git submodule add -b gh --name gh git://github.com/jingweno/gh.git vendor/gh

### gh pull-request

    # while on a topic branch called "feature":
    $ gh pull-request
    [ opens text editor to edit title & body for the request ]
    [ opened pull request on GitHub for "YOUR_USER:feature" ]

    # explicit pull base & head:
    $ gh pull-request -b jingweno:master -h jingweno:feature

    $ gh pull-request -m "title\n\nbody"
    [ create pull request with title & body  ]

    $ gh pull-request -i 123
    [ attached pull request to issue #123 ]

    $ gh pull-request https://github.com/jingweno/gh/pull/123
    [ attached pull request to issue #123 ]

    $ gh pull-request -F FILE
    [ create pull request with title & body from FILE ]

### gh apply

    $ gh apply https://github.com/jingweno/gh/pull/55
    > curl https://github.com/jingweno/gh/pull/55.patch -o /tmp/55.patch
    > git apply /tmp/55.patch

    $ gh apply --ignore-whitespace https://github.com/jingweno/gh/commit/fdb9921
    > curl https://github.com/jingweno/gh/commit/fdb9921.patch -o /tmp/fdb9921.patch
    > git apply --ignore-whitespace /tmp/fdb9921.patch

    $ gh apply https://gist.github.com/8da7fb575debd88c54cf
    > curl https://gist.github.com/8da7fb575debd88c54cf.txt -o /tmp/gist-8da7fb575debd88c54cf.txt
    > git apply /tmp/gist-8da7fb575debd88c54cf.txt

### gh fork

    $ gh fork
    [ repo forked on GitHub ]
    > git remote add -f YOUR_USER git@github.com:YOUR_USER/CURRENT_REPO.git

    $ gh fork --no-remote
    [ repo forked on GitHub ]

### gh create

    $ gh create
    ... create repo on github ...
    > git remote add -f origin git@github.com:YOUR_USER/CURRENT_REPO.git

    # with description:
    $ gh create -d 'It shall be mine, all mine!'

    $ gh create recipes
    [ repo created on GitHub ]
    > git remote add origin git@github.com:YOUR_USER/recipes.git

    $ gh create sinatra/recipes
    [ repo created in GitHub organization ]
    > git remote add origin git@github.com:sinatra/recipes.git

### gh ci-status

    $ gh ci-status
    > (prints CI state of HEAD and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

    $ gh ci-status -v
    > (prints CI state of HEAD, the URL to the CI build results and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

    $ gh ci-status BRANCH
    > (prints CI state of BRANCH and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

    $ gh ci-status SHA
    > (prints CI state of SHA and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)
    
### gh browse

    $ gh browse
    > open https://github.com/YOUR_USER/CURRENT_REPO

    $ gh browse commit/SHA
    > open https://github.com/YOUR_USER/CURRENT_REPO/commit/SHA

    $ gh browse issues
    > open https://github.com/YOUR_USER/CURRENT_REPO/issues

    $ gh browse -p jingweno/gh
    > open https://github.com/jingweno/gh

    $ gh browse -p jingweno/gh commit/SHA
    > open https://github.com/jingweno/gh/commit/SHA

    $ gh browse -p resque
    > open https://github.com/YOUR_USER/resque

    $ gh browse -p resque network
    > open https://github.com/YOUR_USER/resque/network

### gh compare

    $ gh compare refactor
    > open https://github.com/CURRENT_REPO/compare/refactor

    $ gh compare 1.0..1.1
    > open https://github.com/CURRENT_REPO/compare/1.0...1.1

    $ gh compare -u other-user patch
    > open https://github.com/other-user/REPO/compare/patch

### gh release (beta)

    $ gh release
    > (prints a list of releases of YOUR_USER/CURRENT_REPO)


### gh issues (beta)

    $ gh issues
    > (prints a list of issues for YOUR_USER/CURRENT_REPO)


## Release Notes

See [Releases](https://github.com/jingweno/gh/releases).

## Roadmap

See [Issues](https://github.com/jingweno/gh/issues?labels=feature&page=1&state=open).

## script/bootstrap

This script will get all the dependencies ready so you can start hacking on gh.

```
$ ./script/bootstrap
```

## script/build

This script will build gh. It will also perform script/bootstrap, which gets all dependencies and all that jazz.

```
$ ./script/build
```

## script/release

This script will cross-compile gh and prepare for release.

```
$ ./script/release
```

## script/test

For your convenience, there is a script to run the tests.

```
$ ./script/test
```

## Contributors

See [Contributors](https://github.com/jingweno/gh/graphs/contributors).

## License

gh is released under the MIT license. See [LICENSE.md](https://github.com/jingweno/gh/blob/master/LICENSE.md).

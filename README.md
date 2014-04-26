g + h = github [![Build Status](https://travis-ci.org/jingweno/gh.png?branch=master)](https://travis-ci.org/jingweno/gh)
==============

![gh](http://owenou.com/gh/images/gangnamtocat.png)

Fast GitHub command line client implemented in Go.
Current version is [v2.0.0](https://github.com/jingweno/gh/releases/tag/v2.0.0).
[Moving forward gh will be known as GitHub CLI](https://github.com/github/hub/issues/475).

Overview
--------

`gh` is a command line client to GitHub. It's designed to run as fast as possible with easy installation across operating systems.
If you like gh, please also take a look at [hub](https://github.com/github/hub). Hub is a reference implementation to gh.

Motivation
----------

**Muti-platforms**

gh is fully implemented in the Go language and is designed to run across operating systems.

**Easy installation**

There're no pre-requirements to install gh. Download the [binary](https://github.com/jingweno/gh/releases) and go!

**Fast**

    $ hub version
    git version 1.8.2.3
    hub version 1.11.1

    $ gh version
    git version 1.8.2.3
    gh version 1.0.0

    $ time hub version > /dev/null
    hub version > /dev/null  0.03s user 0.01s system 91% cpu 0.042 total

    $ time gh version > /dev/null
    gh version > /dev/null  0.00s user 0.01s system 86% cpu 0.014 total

    $ time hub browse > /dev/null
    hub browse > /dev/null  0.07s user 0.04s system 87% cpu 0.050 total

    $ time gh browse -u >> /dev/null
    gh browse -u >> /dev/null  0.01s user 0.02s system 81% cpu 0.036 total

Installation
------------

### Homebrew

The easiest way to install `gh` on OSX is through [Homebrew](https://github.com/mxcl/homebrew).
You can add the [gh Homebrew repository](https://github.com/jingweno/homebrew-gh) with [`brew tap`](https://github.com/mxcl/homebrew/wiki/brew-tap):

    $ brew install jingweno/gh/gh
    $ brew install --build-from-source jingweno/gh/gh # build gh from source
    $ brew install --HEAD jingweno/gh/gh # build gh HEAD from source

### Standalone

`gh` can be easily installed as an executable.
Download the latest [compiled binary forms of gh](https://github.com/jingweno/gh/releases) for Darwin, Linux and Windows.

### Boxen

If you're using [boxen](http://boxen.github.com/), there's a [puppet-gh](https://github.com/boxen/puppet-gh) module to install and set up `gh`.

### Source

To install gh from source, you need to have a [Go development environment](http://golang.org/doc/install), version 1.1 or better, and run:

    $ git clone https://github.com/jingweno/gh.git
    $ cd gh
    $ script/install

Note that `go get` will pull down sources from various VCS.
Please make sure you have git and hg installed.

Update
------

`gh` comes with a command to self update:

    $ git selfupdate

### Autoupdate

`gh` checks every two weeks for newer versions and prompts you for update if there's one.
A timestamp is stored in `~/.config/gh-update` for the next update time.

You can enable to always update automatically by answering `always` when a new version is released, or setting the global git config `gh.autoUpdate`:

```
$ git config --global gh.autoUpdate always
```

You can also disable completely automatic updates by answering `never` when a new version is released, or setting the global git config `gh.autoUpdate`:

```
$ git config --global gh.autoUpdate never
```

### Homebrew

If you installed `gh` with `brew tap jingweno/gh`, you can update it with:

    $ brew upgrade gh

### Source

To update gh from source, run:

    $ cd GH_SOURCE_DIR
    $ git pull origin master
    $ script/install

Aliasing
--------

It's best to use `gh` by aliasing it to `git`.
All git commands will still work with `gh` adding some sugar.

`gh alias` displays instructions for the current shell. With the `-s` flag,
it outputs a script suitable for `eval`.

You should place this command in your `.bash_profile` or other startup
script:

    eval "$(gh alias -s)"

For more details, run `gh help alias`.

Commands
--------

Assuming you've aliased gh as `git`, the following commands now have:

### git init

    $ git init -g
    > git init
    > git remote add origin git@github.com:YOUR_USER/REPO.git

### git push

    $ git push origin,staging,qa bert_timeout
    > git push origin bert_timeout
    > git push staging bert_timeout
    > git push qa bert_timeout

    $ git push origin
    > git push origin HEAD

### git checkout

    $ git checkout https://github.com/jingweno/gh/pull/35
    > git remote add -f -t feature git://github:com/foo/gh.git
    > git checkout --track -B foo-feature foo/feature

    $ git checkout https://github.com/jingweno/gh/pull/35 custom-branch-name

### git merge

    $ git merge https://github.com/jingweno/gh/pull/73
    > git fetch git://github.com/jingweno/gh.git +refs/heads/feature:refs/remotes/jingweno/feature
    > git merge jingweno/feature --no-ff -m 'Merge pull request #73 from jingweno/feature...'

### git clone

    $ git clone jingweno/gh
    > git clone git://github.com/jingweno/gh

    $ git clone -p jingweno/gh
    > git clone git@github.com:jingweno/gh.git

    $ git clone jekyll_and_hype
    > git clone git://github.com/YOUR_LOGIN/jekyll_and_hype.

    $ git clone -p jekyll_and_hype
    > git clone git@github.com:YOUR_LOGIN/jekyll_and_hype.git

### git fetch

    $ git fetch jingweno
    > git remote add jingweno git://github.com/jingweno/REPO.git
    > git fetch jingweno

    $ git fetch jingweno,foo
    > git remote add jingweno ...
    > git remote add foo ...
    > git fetch --multiple jingweno foo

    $ git fetch --multiple jingweno foo
    > git remote add jingweno ...
    > git remote add foo ...
    > git fetch --multiple jingweno foo

### git cherry-pick

    $ git cherry-pick https://github.com/jingweno/gh/commit/a319d88#comments
    > git remote add -f jingweno git://github.com/jingweno/gh.git
    > git cherry-pick a319d88

    $ git cherry-pick jingweno@a319d88
    > git remote add -f jingweno git://github.com/jingweno/gh.git
    > git cherry-pick a319d88

    $ git cherry-pick jingweno@SHA
    > git fetch jingweno
    > git cherry-pick SHA

### git remote

    $ git remote add jingweno
    > git remote add -f jingweno git://github.com/jingweno/CURRENT_REPO.git

    $ git remote add -p jingweno
    > git remote add -f jingweno git@github.com:jingweno/CURRENT_REPO.git

    $ git remote add origin
    > git remote add -f YOUR_USER git://github.com/YOUR_USER/CURRENT_REPO.git

### git submodule

    $ git submodule add jingweno/gh vendor/gh
    > git submodule add git://github.com/jingweno/gh.git vendor/gh

    $ git submodule add -p jingweno/gh vendor/gh
    > git submodule add git@github.com:jingweno/gh.git vendor/gh

    $ git submodule add -b gh --name gh jingweno/gh vendor/gh
    > git submodule add -b gh --name gh git://github.com/jingweno/gh.git vendor/gh

### git pull-request

    # while on a topic branch called "feature":
    $ git pull-request
    [ opens text editor to edit title & body for the request ]
    [ opened pull request on GitHub for "YOUR_USER:feature" ]

    # explicit pull base & head:
    $ git pull-request -b jingweno:master -h jingweno:feature

    $ git pull-request -m "title\n\nbody"
    [ create pull request with title & body  ]

    $ git pull-request -i 123
    [ attached pull request to issue #123 ]

    $ git pull-request https://github.com/jingweno/gh/pull/123
    [ attached pull request to issue #123 ]

    $ git pull-request -F FILE
    [ create pull request with title & body from FILE ]

### git apply

    $ git apply https://github.com/jingweno/gh/pull/55
    > curl https://github.com/jingweno/gh/pull/55.patch -o /tmp/55.patch
    > git apply /tmp/55.patch

    $ git apply --ignore-whitespace https://github.com/jingweno/gh/commit/fdb9921
    > curl https://github.com/jingweno/gh/commit/fdb9921.patch -o /tmp/fdb9921.patch
    > git apply --ignore-whitespace /tmp/fdb9921.patch

    $ git apply https://gist.github.com/8da7fb575debd88c54cf
    > curl https://gist.github.com/8da7fb575debd88c54cf.txt -o /tmp/gist-8da7fb575debd88c54cf.txt
    > git apply /tmp/gist-8da7fb575debd88c54cf.txt

### git fork

    $ git fork
    [ repo forked on GitHub ]
    > git remote add -f YOUR_USER git@github.com:YOUR_USER/CURRENT_REPO.git

    $ git fork --no-remote
    [ repo forked on GitHub ]

### git create

    $ git create
    ... create repo on github ...
    > git remote add -f origin git@github.com:YOUR_USER/CURRENT_REPO.git

    # with description:
    $ git create -d 'It shall be mine, all mine!'

    $ git create recipes
    [ repo created on GitHub ]
    > git remote add origin git@github.com:YOUR_USER/recipes.git

    $ git create sinatra/recipes
    [ repo created in GitHub organization ]
    > git remote add origin git@github.com:sinatra/recipes.git

### git ci-status

    $ git ci-status
    > (prints CI state of HEAD and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

    $ git ci-status -v
    > (prints CI state of HEAD, the URL to the CI build results and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

    $ git ci-status BRANCH
    > (prints CI state of BRANCH and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

    $ git ci-status SHA
    > (prints CI state of SHA and exits with appropriate code)
    > One of: success (0), error (1), failure (1), pending (2), no status (3)

### git browse

    $ git browse
    > open https://github.com/YOUR_USER/CURRENT_REPO

    $ git browse commit/SHA
    > open https://github.com/YOUR_USER/CURRENT_REPO/commit/SHA

    $ git browse issues
    > open https://github.com/YOUR_USER/CURRENT_REPO/issues

    $ git browse -p jingweno/gh
    > open https://github.com/jingweno/gh

    $ git browse -p jingweno/gh commit/SHA
    > open https://github.com/jingweno/gh/commit/SHA

    $ git browse -p resque
    > open https://github.com/YOUR_USER/resque

    $ git browse -p resque network
    > open https://github.com/YOUR_USER/resque/network

### git compare

    $ git compare refactor
    > open https://github.com/CURRENT_REPO/compare/refactor

    $ git compare 1.0..1.1
    > open https://github.com/CURRENT_REPO/compare/1.0...1.1

    $ git compare -u other-user patch
    > open https://github.com/other-user/REPO/compare/patch

### git release (beta)

    $ git release
    > (prints a list of releases of YOUR_USER/CURRENT_REPO)

    $ git release create TAG
    > (creates a new release for the given tag)

### git issues (beta)

    $ git issue
    > (prints a list of issues for YOUR_USER/CURRENT_REPO)

    $ git issue create
    > (creates an issue for the project that "origin" remote points to)

Configuration
-------------

### GitHub OAuth authentication

`gh` will prompt for GitHub username & password the first time it needs
to access the API and exchange it for an OAuth token, which it saves in `~/.config/gh`.
You could specify the path to the config by setting the `GH_CONFIG` environment variable.

### HTTPS instead of git protocol

If you prefer using the HTTPS protocol for GitHub repositories instead
of the git protocol for read and ssh for write, you can set
"gh.protocol" to "https".

    # default behavior
    $ git clone jingweno/gh
    < git clone >

    # opt into HTTPS:
    $ git config --global gh.protocol https
    $ git clone jingweno/gh
    < https clone >

### GitHub Enterprise

By default, `gh` will only work with repositories that have remotes which
point to `github.com`. GitHub Enterprise hosts need to be whitelisted to
configure `gh` to treat such remotes same as `github.com`:

    $ git config --global --add gh.host my.git.org

The default host for commands like `init` and
`clone` is still `github.com`, but this can be affected with the `GITHUB_HOST` environment variable:</p>

    $ GITHUB_HOST=my.git.org git clone myproject

### Crash reports

`gh` includes automatic crash reporting in case that something unexpected happens.
It will ask you if you want to report the error to us if the program terminates suddenly, and
then it will open an issue on your behalf under [the crash report issues](https://github.com/jingweno/gh/issues?labels=Crash+Report&page=1&state=open).

`gh` doesn't send any information about the command that you ran.
Check [some examples](https://github.com/jingweno/gh/issues?labels=Crash+Report&state=closed) of the information included by default, you can always modify it before the issue is open.

You can enable to always send crash reports with the default information by answering `always` when a crash error happens, or setting the global git config `gh.reportCrash`:

```
$ git config --global gh.reportCrash always
```

You can also disable completely crash report notifications by answering `never` when a crash report happens, or setting the global git config `gh.reportCrash`:

```
$ git config --global gh.reportCrash never
```

Release Notes
-------------

See [Releases](https://github.com/jingweno/gh/releases).

Roadmap
-------

See [Issues](https://github.com/jingweno/gh/issues?labels=feature&page=1&state=open).

Development
-----------

[Godep](https://github.com/kr/godep) is used to lock down all the
dependencies. You need to use the wrapper scripts to build/install/test `gh`:

### script/bootstrap

This script will get all the dependencies ready so you can start hacking on gh.

```
$ ./script/bootstrap
```

### script/build

This script will build gh. It will also perform script/bootstrap, which gets all dependencies and all that jazz.

```
$ ./script/build
```

### script/install

This script will build and install gh. It will also perform script/bootstrap, which gets all dependencies and all that jazz.

```
$ ./script/install
```

### script/package

This script will cross-compile gh and package release for current platform.
It executes the `package` [gotask](https://github.com/jingweno/gotask) [task](https://github.com/jingweno/gh/blob/master/gh_task.go).

```
$ ./script/package
```

### script/test

For your convenience, there is a script to run the tests.

```
$ ./script/test
```

See [this guide](https://github.com/jingweno/gh/blob/master/CONTRIBUTING.md) on how to submit a pull request.

Contributors
------------

See [Contributors](https://github.com/jingweno/gh/graphs/contributors).

License
-------

gh is released under the MIT license. See [LICENSE.md](https://github.com/jingweno/gh/blob/master/LICENSE.md).

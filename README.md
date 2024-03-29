[![Build Status](https://travis-ci.org/warrensbox/jscheck.svg?branch=master)](https://travis-ci.org/warrensbox/jscheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/jscheck)](https://goreportcard.com/report/github.com/warrensbox/jscheck)
[![CircleCI](https://circleci.com/gh/warrensbox/jscheck/tree/release.svg?style=shield&circle-token=59e89260bc775454057c7e775977510f03130ae1)](https://circleci.com/gh/warrensbox/jscheck)

# JSCheck

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/jscheck/smallerlogo.png" alt="drawing" width="120" height="130"/>

<!-- ![gopher](https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/logo.png =100x20) -->

The `jscheck` command line tool looks for json files and check's its validity.
Sometimes not all our json files end with a `.json` extension. `jscheck` tries to figure out if your file has is a json format before linting.
The installation is minimal and easy.

## Installation

`jscheck` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/).


```ruby
brew install warrensbox/tap/jscheck
```

### Linux

Installation for other linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/jscheck/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/jscheck/releases)

## How to use:
### Without args
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/jscheck/jscheck_v1.gif" alt="drawing" style="width: 210px;"/>

1.  By typing the command `jscheck` on your terminal, it walks through all your current directory and sub-directory to determine any errors in json formatted files.

The most recently selected versions are presented at the top of the dropdown.

### Specify a directory 
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/jscheck/jscheck_v2.gif" alt="drawing" style="width: 170px;"/>

1. You can also supply the desired directory to walk through as an argument on the command line.
2. For example, `jscheck -d dirname` for walk through `dirname`.

### Execute as part of jenkins job

```
    sh """\
        #!/bin/bash 
        eval "\$(chef shell-init bash)"
        echo install jscheck
        wget https://raw.githubusercontent.com/warrensbox/jscheck/release/install.sh 
        chmod 755 install.sh
        ./install.sh -b installs

        ./installs/jscheck
        """.stripIndent()
```

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/jscheck/issues)

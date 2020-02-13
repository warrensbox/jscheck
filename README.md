[![Build Status](https://travis-ci.org/warrensbox/jscheck.svg?branch=master)](https://travis-ci.org/warrensbox/jscheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/jscheck)](https://goreportcard.com/report/github.com/warrensbox/jscheck)
[![CircleCI](https://circleci.com/gh/warrensbox/jscheck/tree/master.svg?style=shield&circle-token=55ddceec95ff67eb38269152282f8a7d761c79a5)](https://circleci.com/gh/warrensbox/jscheck)

# JSCheck

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/smallerlogo.png" alt="drawing" width="120" height="130"/>

<!-- ![gopher](https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/logo.png =100x20) -->

The `jscheck` command line tool looks for json files and check's it's validity.
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
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/jscheck/jscheck_v1.gif" alt="drawing" style="width: 180px;"/>

1.  By typing the command `jscheck` on your terminal, it walk through all your directory let lets you know if there is an error in your json format files.
2.  When you don't specify a directory, it will check your current and child directories.


The most recently selected versions are presented at the top of the dropdown.

### Supply directory on command line
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/jscheck/jscheck_v2.gif" alt="drawing" style="width: 170px;"/>

1. You can also supply the desired directory to walk through as an argument on the command line.
2. For example, `jscheck -d dirname` for walk through `dirname`.

**Execute as part of jenkins job**

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
[![Build Status](https://travis-ci.org/warrensbox/jscheck.svg?branch=master)](https://travis-ci.org/warrensbox/jscheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/jscheck)](https://goreportcard.com/report/github.com/warrensbox/jscheck)
[![CircleCI](https://circleci.com/gh/warrensbox/jscheck/tree/master.svg?style=shield&circle-token=55ddceec95ff67eb38269152282f8a7d761c79a5)](https://circleci.com/gh/warrensbox/jscheck)

# JSCheck

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/smallerlogo.png" alt="drawing" width="120" height="130"/>

<!-- ![gopher](https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/logo.png =100x20) -->

The `jscheck` command line tool lets you look for json files and check's it's validity.
Sometimes we not all our json files end with a `.json` extension. `jscheck` tries to figure out if your json file has is a json file and then runs a lint check.
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
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch.gif" alt="drawing" style="width: 180px;"/>

1.  You can switch between different versions of terraform by typing the command `tfswitch` on your terminal.
2.  Select the version of terraform you require by using the up and down arrow.


The most recently selected versions are presented at the top of the dropdown.

### Supply directory on command line
<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch-v4.gif" alt="drawing" style="width: 170px;"/>

1. You can also supply the desired version as an argument on the command line.
2. For example, `tfswitch 0.10.5` for version 0.10.5 of terraform.
3. Hit **Enter** to switch.

**Execute part jenkins job**

```
cdtfswitch(){
  builtin cd "$@";
  cdir=$PWD;
  if [ -f "$cdir/.tfswitchrc" ]; then
    tfswitch
  fi
}
alias cd='cdtfswitch'
```

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/jscheck/issues)
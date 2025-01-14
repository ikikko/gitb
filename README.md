# gitb [![Build Status](https://travis-ci.org/vvatanabe/gitb.svg?branch=master)](https://travis-ci.org/vvatanabe/gitb) [![Coverage Status](https://coveralls.io/repos/github/vvatanabe/gitb/badge.svg?branch=master)](https://coveralls.io/github/vvatanabe/gitb?branch=master)

A command line tool for using Backlog's git comfortably.

## Description

`gitb` command helps to use Backlog's git comfortably. For example, can open PR, issue, branches, tags, etc in the browser with one action. 

Also, `gitb` wraps all `git` commands, therefore can execute `git-command` using `gitb` like a `gitb fetch`, `gitb pull`,  `gitb push`, and more...

Suffix B has multiple meanings. Backlog, Browser, B-Dash.

## Installation

Built binaries are available on Github releases:  
https://github.com/vvatanabe/gitb/releases

This package can be installed with the go get command too:

`$ go get github.com/vvatanabe/gitb`

## Usage

### Pull Request

Related to Backlog Pull Requests for the current repository.

__COMMANDS:__

`gitb pr [-s <STATE>]`

&emsp;Open the pull request list page in the current repository.

`gitb pr show`

&emsp;Open the pull request page related to the current branch.

`gitb pr add [-b <BASE>]`

&emsp;Open the page to create pull request with the current branch.


__OPTIONS:__

`-s, --state <STATE>`

&emsp;Filter pull requests by STATE. Values: "open" (default), "closed", "merged", "all".

`-b, --base <BASE>`

&emsp;BASE is base branch name. Default is empty.

### Issue

Related to Backlog Issues for the current repository.

__COMMANDS:__

`gitb issue [-s <STATE>]`

&emsp;Open the issue list page in the current project.

`gitb issue show`

&emsp;Open the issue page related to the current branch.

`gitb issue add`

&emsp;Open the page to create issue in the current project.

__OPTIONS:__

`-s, --state <STATE>`

&emsp;Filter issues by STATE. Values: "all", "open", "in_progress", "resolved", "closed", "not_closed" (default).

### Browse

Open other git page (e.g. branch, tree, tag, and more...) in current repository.

__COMMANDS:__

`gitb browse branch`

&emsp;Open the branch list page in the current repository.

`gitb browse tag`

&emsp;Open the tag list page in the current repository.

`gitb browse tree`

&emsp;Open the tree page in the current branch.

`gitb browse history`

&emsp;Open the history page in the current branch.

`gitb browse network`

&emsp;Open the network page in the current branch.

`gitb browse repo`

&emsp;Open the repository list page in the current project.

## Alias 

Please write an alias to .XXXrc (.bashrc, .zshrc, config.fish) if you want to use `gitb <command>` as `git <command>`.

### Bash

```
eval "$(gitb alias -s)"
```

### Zsh

```
function git(){gitb "$@"}
```

### Fish

```
function git
  gitb $argv
end
```

## Acknowledgments

Inspired by [github.com/github/hub](https://github.com/github/hub)

## Bugs and Feedback

For bugs, questions and discussions please use the GitHub Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)

## Author

[vvatanabe](https://github.com/vvatanabe)
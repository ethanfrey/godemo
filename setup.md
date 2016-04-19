Getting the shell set up
========================

First, download and install go in a directory.  This path will be known as GOPATH from now on. Downloads are available here https://golang.org/dl/ (use at least 1.6 for this intro)

Now, set up either `~/.bashrc` or `~/.zshrc` with the following lines:

```
export GOPATH=/home/user/my/path
export PATH=$PATH:$GOPATH/bin
# this is a bit more contreversial, but go install doesn't work without it...
export GOBIN=$GOPATH/bin
```

Make a directory and make sure `go run <file>.go` works. Yippee!!


Adding Tools
============

There are a lot of cool tools to make development nicer, here are some:

`go doc`: type any package or class and get help on it.  kind of like man pages and always one stroke away at your keyboard, great place to look before googling
`go fmt`: Reformat your code with the go styling guide.  Makes it easy to just write without worrying about spacing and have everything nice and standard later


Others are useful when integrated with an editor

`golint`: Code linting
`go vet`: Deeper analysis for potential problems, before compiling
`go oracle`: View context info, nice when looking in a large codebase you don't know
`godepgraph`: Build a dependency graph to understand what depends on what, when anaylsing large projects


To get all these going:

```
go get -u golang.org/x/tools/cmd/vet
go get -u github.com/golang/lint/golint
go get -u golang.org/x/tools/cmd/oracle
go get -u golang.org/x/tools/cmd/gotype
```

If you want the dependency graph:

```
go get github.com/kisielk/godepgraph
brew install graphviz
```


Getting Sublime3 Configured
===========================

Okay, now we got all the command line tools set up, how to get Sublime working like a lean, mean machine?  There is some help here: http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/ but I will make a slightly different take on it.

First, install the following with package control:

* https://github.com/DisposaBoy/GoSublime
* https://github.com/waigani/GoOracle
* https://github.com/SirReal/SublimeLinter-contrib-golint
* https://github.com/sirreal/SublimeLinter-contrib-gotype

Edit Preferences -> Package Settings -> GoSublime -> Settings User:

```
{

    // you may set specific environment variables here
    // e.g "env": { "PATH": "$HOME/go/bin:$PATH" }
    // in values, $PATH and ${PATH} are replaced with
    // the corresponding environment(PATH) variable, if it exists.
    "env": {"GOPATH": "$HOME/work", "PATH": "$GOPATH/bin:$PATH" },
    // TODO: is this needed?? or just a restart? or bash stuff??
    "paths":
    {
        "osx":
        [
            "$HOME/work/bin"
        ]
    },
    "fmt_cmd": ["goimports"],

    // enable comp-lint, this will effectively disable the live linter
    "comp_lint_enabled": true,

    // list of commands to run
    "comp_lint_commands": [
        // run `golint` on all files in the package
        // "shell":true is required in order to run the command through your shell (to expand `*.go`)
        // also see: the documentation for the `shell` setting in the default settings file ctrl+dot,ctrl+4
        {"cmd": ["golint *.go"], "shell": true},

        // this actually checks for valid syntax and function calls
        {"cmd": ["gotype *.go"], "shell": true},

        // run go vet on the package
        // {"cmd": ["go", "vet"]},

        // run `go install` on the package. GOBIN is set,
        // so `main` packages shouldn't result in the installation of a binary
        // {"cmd": ["go", "install"]}
    ],

    "on_save": [
        // run comp-lint when you save,
        // naturally, you can also bind this command `gs_comp_lint`
        // to a key binding if you want
        {"cmd": "gs_comp_lint"}
    ]
}
```

Note: Ensure you update the GOPATH value to match the one configured earlier.

I work on other languages with other formating (python and javascript), so it is nice to make some go syntax settings.
Open up a .go file and edit Preferences -> Settings - More -> Syntax Specific

```
{
    "tab_size": 2,
    "shift_tab_unindent": true,
    "use_tab_stops": true,
    "translate_tabs_to_spaces": false
}
```

You can add this to keybindings to trigger the linter not just on save:

```
    { "keys": ["alt+l"], "command": "gs_comp_lint"}
```


Testing Code
============


Some (untested) ideas on debuggers
==================================

* Looks nice, generates special debug code hooks
    http://blog.mailgun.com/introducing-a-new-cross-platform-debugger-for-go/
    https://github.com/mailgun/godebug
* In development but can attach to running (production) programs
    https://github.com/derekparker/delve
* Good, old gdb :(
    https://golang.org/doc/gdb


Other Links
===========

Maybe you want git integration as well (install package GitGutter)

Cross-Compiling: http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5

Interesting if you want to cross-compile for mobile platforms (alpha/beta):

```
go get golang.org/x/mobile/cmd/gomobile
```

A build tool in alpha, but maybe useful for larger build tasks:
https://getgb.io/


Thoughts to Investigate
=======================

GoGenerateTools

* https://github.com/golang/go/wiki/CodeTools
  * json-to-go, gonerics, gorename
* https://github.com/golang/go/wiki/GoGenerateTools
  * for example, adding strings for enums...



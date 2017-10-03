# cheat

cheat generates the files for creating a [Dash][1] cheat sheet. It requires
[cheatset][2].

## Installation

Use Go to install it from this repo: 

    go get github.com/lorin/cheat

## Usage

    cheat <projectname>

For example:

    cheat foo

This will create the following directory and files:

```
foo-cheat-sheet
├── .gitignore
├── Makefile
├── README.md
└── foo.rb
```

[1]: https://kapeli.com/dash
[2]: https://github.com/Kapeli/cheatset

# `tree`

Wraps `$ tree` with a pipable version.

## Usage

```bash
# Regular tree still works:
$ tree
# => 
# .
# ├── tmp
# └── src
#     └── linux_amd64
#         └── github.com
#             └── homburg
#                 ├── cli.a
#                 └── envoke.a

# Pipable tree
$ find pkg | tree "/"
# .
# └── pkg
#     └── linux_amd64
#         └── github.com
#             └── homburg
#                 ├── cli.a
#                 └── envoke.a

$ git ls-files | tree
# .travis.yml
# LICENSE
# Makefile
# README.md
# cmd
# └── tree
#     ├── README.md
#     └── main.go
# data.csv
# files.txt
# tree.go
# tree_test.go
```

# LICENSE

[LICENSE.md](LICENSE.md)

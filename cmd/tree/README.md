# tree

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
```

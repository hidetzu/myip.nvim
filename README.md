# myip.nvim
myip.nvim is the nvim plugin using [myip](https://github.com/hidetzu/myip).

## Required

- [golang](https://go.dev/)
- [go-client](https://github.com/neovim/go-client)

## Installing

with [packer.nvim](https://github.com/wbthomason/packer.nvim)

```lua
use {
    'hidetzu/myip.vim',
    run ="make install"
}
```

## Useage

it' example. 

use ex, print LocalAddress()
```vim
:echo LocalAddress()
```

use ex, print GlobalAddress()
```vim
:echo GlobalAddress()
```

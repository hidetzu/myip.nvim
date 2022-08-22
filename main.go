package main

import (
	"github.com/hidetzu/myip"
	"github.com/neovim/go-client/nvim/plugin"
)

func localAddress(args []string) (string, error) {
    return myip.GetLocalIP()
}

func globalAddress(args []string) (string, error) {
    return myip.GetRemoteAddr(), nil
}

func main() {
    plugin.Main(func(p *plugin.Plugin) error {
        p.HandleFunction(&plugin.FunctionOptions{Name: "LocalAddress"}, localAddress)
        p.HandleFunction(&plugin.FunctionOptions{Name: "GlobalAddress"}, globalAddress)
        return nil
    })
}

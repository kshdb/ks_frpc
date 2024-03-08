package main

import (
	"fmt"

	frpc "github.com/kshdb/ks_frpc"
)

func main() {
	_content := fmt.Sprintf(`
[common]
server_addr = %s
server_port = %s
token=%s

[api]
type = %s
local_ip = %s
local_port = %s
remote_port = %s
`, "test2.api.cnwtn.com", "7005", "AvbuYer!35.com", "tcp", "127.0.0.1", "8077", "18077")
	err := frpc.RunContent(_content)
	panic(err)
}

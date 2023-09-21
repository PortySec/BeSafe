package scan

import (
	"net"
)

func Scanner(host string, port string) bool {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true

}

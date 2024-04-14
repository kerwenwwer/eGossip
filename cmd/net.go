package cmd

import "fmt"

// write
func write(nodeList *NodeList, addr string, port int, data []byte) {
	udpWrite(nodeList, addr, port, data)
}

// listen
func listen(nodeList *NodeList, mq chan []byte) {
	if nodeList.Protocol == "UDP" || nodeList.Protocol == "TC" {
		udpListen(nodeList, mq)
	} else if nodeList.Protocol == "XDP" {
		xdpListen(nodeList, mq)
	} else {
		fmt.Println("Protocol not supported, only UDP, TC and XDP.")
	}
}

package main

import (
	"mychain/chainDemo/core"
)

func main() {
	bc := core.NewBlockChain()
	bc.SendData("Send 1 BTC To Jacky")
	bc.SendData("Send 1 EOS To Jack")
	bc.Print()
}

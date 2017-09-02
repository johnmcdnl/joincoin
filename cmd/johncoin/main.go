package main

import (
	"johncoin"
)

func main() {
	bc := johncoin.NewBlockchain()
	defer bc.Close()

	cli := johncoin.CLI{BlockChain: bc}
	cli.Run()
}

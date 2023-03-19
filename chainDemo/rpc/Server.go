package main

import (
	"encoding/json"
	"io"
	"mychain/chainDemo/core"
	"net/http"
)

var blockChain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func blockchainGetHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(blockChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(b))
}
func blockchainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockchainGetHandler(w, r)
}

func main() {
	blockChain = core.NewBlockChain()
	run()
}

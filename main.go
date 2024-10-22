package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rishavmehra/bittorrent-v2/bencode"
)

func main() {
	cmd := os.Args[1]

	if cmd == "decode" {
		bencodeValue := os.Args[2]

		decoded, err := bencode.DecodeBencode(bencodeValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("unknown command: " + cmd)
		os.Exit(1)
	}
}

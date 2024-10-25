package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rishavmehra/bittorrent-v2/bencode"
	"github.com/rishavmehra/bittorrent-v2/torrentfile"
)

func main() {
	cmd := os.Args[1]

	switch cmd {
	case "decode":
		bencodeValue := os.Args[2]

		decoded, err := bencode.DecodeBencode(bencodeValue)
		if err != nil {
			fmt.Println(err)
		}
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
		return
	case "info":
		filePath := os.Args[2]
		torrent, err := torrentfile.NewTorrentFile(filePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Tracker URL: %s\n", torrent.Announce)
		fmt.Printf("Lenght: %d\n", torrent.Info.Length)
		fmt.Printf("name: %s\n", torrent.Info.Name)
		infoHash, err := torrent.InfoHash()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Info Hash: %s\n", infoHash)
		}
		return
	default:
		fmt.Println("Unknown command:", cmd)
	}
	os.Exit(1)

}

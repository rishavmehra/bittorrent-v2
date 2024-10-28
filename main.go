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
			os.Exit(1)
		}
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
		return
	case "info":
		filePath := os.Args[2]
		torrent, err := torrentfile.NewTorrentFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Tracker URL: %s\n", torrent.Announce)
		fmt.Printf("Lenght: %d\n", torrent.Info.Length)
		fmt.Printf("name: %s\n", torrent.Info.Name)
		infoHash, err := torrent.InfoHash()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("Info Hash: %x\n", infoHash)
		}
		fmt.Printf("Pieces Length:%d\n", torrent.Info.PieceLength)
		fmt.Printf("Pieces Hash: \n")
		for i := 0; i < len(torrent.Info.Pieces); i = i + 20 {
			fmt.Printf("%x\n", torrent.Info.Pieces[i:i+20])
		}
		return

	case "peers":
		filePath := os.Args[2]
		file, err := torrentfile.NewTorrentFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tracker := &torrentfile.TrackerReq{PeerID: "00000000000000000000", Port: 1234, Compact: 1}
		response, err := file.GetTrackerResponse(tracker)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for i := 0; i < len(response.Peers); i++ {
			fmt.Println(response.Peers[i])
		}
	default:
		fmt.Println("Unknown command:", cmd)
	}
	os.Exit(1)
}

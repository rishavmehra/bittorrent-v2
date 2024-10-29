package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/rishavmehra/bittorrent-v2/bencode"
	"github.com/rishavmehra/bittorrent-v2/handshake"
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

	case "handshake":
		filePath := os.Args[2]
		peerInfo := os.Args[3]

		file, err := torrentfile.NewTorrentFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		tracker := &torrentfile.TrackerReq{PeerID: "00000000000000000000", Port: 1234, Compact: 1}
		peerID := []byte(tracker.PeerID)
		infoHash, err := file.InfoHash()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		msg := handshake.NewHandshake([20]byte(peerID), infoHash)
		var buf bytes.Buffer
		err = binary.Write(&buf, binary.BigEndian, msg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		conn, err := net.Dial("tcp", peerInfo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		writer.Write(buf.Bytes())
		err = writer.Flush()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader := bufio.NewReader(conn)
		var handshakeRes handshake.Handshake
		err = binary.Read(reader, binary.BigEndian, &handshakeRes)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Peer ID: %x\n", handshakeRes.PeerID)
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

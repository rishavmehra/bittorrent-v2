package torrentfile

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rishavmehra/bittorrent-v2/bencode"
)

type TrackerReq struct {
	PeerID     string
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Compact    int // compact representation of peer addresses in the tracker request ip and port (6 bytes)
}

type TrackerRes struct {
	Interval int
	Peers    []TrackerPeer
}

type TrackerPeer struct {
	IP   net.IP
	Port int
}

func (torrent *TorrentFile) newTrackerRequestURL(trackerReq *TrackerReq) (string, error) {
	infoHash, err := torrent.InfoHash()
	if err != nil {
		return "", err
	}

	trackerParams := url.Values{}
	trackerParams.Set("info_hash", string(infoHash[:]))
	trackerParams.Set("peer_id", trackerReq.PeerID)
	trackerParams.Set("port", strconv.Itoa(trackerReq.Port))
	trackerParams.Set("uploaded", strconv.Itoa(trackerReq.Uploaded))
	trackerParams.Set("downloaded", strconv.Itoa(trackerReq.Downloaded))
	trackerParams.Set("left", strconv.Itoa(torrent.Info.Length))
	trackerParams.Set("compact", strconv.Itoa(trackerReq.Compact))

	trackerRequestURL := fmt.Sprintf("%s?%s", torrent.Announce, trackerParams.Encode())
	return trackerRequestURL, nil
}

func (torrent *TorrentFile) GetTrackerResponse(trackerReq *TrackerReq) (*TrackerRes, error) {
	trackerReqURL, err := torrent.newTrackerRequestURL(trackerReq)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(trackerReqURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	trackerRes, _, err := bencode.DecodeBencodeDic(string(body))
	if err != nil {
		return nil, err
	}

	response := &TrackerRes{Peers: make([]TrackerPeer, 0)}
	peers, err := getInfoValue(trackerRes, "peers", "")
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(peers); i += 6 {
		response.Peers = append(response.Peers, TrackerPeer{IP: net.IPv4(peers[i], peers[i+1], peers[i+2], peers[i+3]), Port: int(binary.BigEndian.Uint16([]byte(peers[i+4:])))})
	}

	response.Interval, err = getInfoValue(trackerRes, "interval", response.Interval)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func (peer TrackerPeer) String() string {
	return fmt.Sprintf("%s:%d", peer.IP, peer.Port)
}

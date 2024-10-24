package torrentfile

import (
	"fmt"
	"os"

	"github.com/rishavmehra/bittorrent-v2/bencode"
)

type TorrentFile struct {
	FilePath string
	Announce string
	Info     TorrentFileInfo
}

type TorrentFileInfo struct {
	Length      int
	Name        string
	PieceLength int
	Pieces      string
}

func getInfoValue[T comparable](info map[string]interface{}, key string, valueType T) (T, error) {
	value, ok := info[key]
	if !ok {
		return valueType, fmt.Errorf("TorrentFile.info.%s:no \"%s\" field in torrent file", key, key)
	}
	val, ok := value.(T)
	if !ok {
		return valueType, fmt.Errorf("TorrentFile.info.%s:invalid\"%s\" field in torrent file", key, key)
	}
	return val, nil
}

func NewTorrentFile(filePath string) (*TorrentFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	stat.Size()
	torrent := &TorrentFile{FilePath: filePath}

	buf := make([]byte, stat.Size())
	size, err := file.Read(buf)
	if err != nil {
		return nil, err
	}
	if len(buf) != size {
		return nil, fmt.Errorf("did not read full torrent file, file size: %d, read: %d", size, len(buf))
	}

	d, _, err := bencode.DecodeBencodeDic(string(buf))
	if err != nil {
		return nil, err
	}

	value, ok := d["announce"]
	if !ok {
		return nil, fmt.Errorf("\"announce\" not found in the torrent file")
	}
	announce, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("\"announce\" in torrent file is not BencodeString")
	}
	torrent.Announce = announce

	value, ok = d["info"]
	if !ok {
		return nil, fmt.Errorf("\"info\" not found in the torrent file")
	}
	info, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("\"info\" in torrent file is not BencodeDict")
	}

	torrent.Info.Length, err = getInfoValue(info, "length", torrent.Info.Length)
	if err != nil {
		return nil, err
	}

	torrent.Info.Name, err = getInfoValue(info, "name", torrent.Info.Name)
	if err != nil {
		return nil, err
	}

	torrent.Info.PieceLength, err = getInfoValue(info, "piece length", torrent.Info.PieceLength)
	if err != nil {
		return nil, err
	}

	torrent.Info.Pieces, err = getInfoValue(info, "pieces", torrent.Info.Pieces)
	if err != nil {
		return nil, err
	}

	return torrent, nil
}

# BitTorrent Implementation in Go

This project is a BitTorrent protocol implementation in Go, aimed at providing essential components for torrent file management and peer-to-peer file sharing.

### Prerequisites

- Go 1.18 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/rishavmehra/bittorrent-v2.git
   cd bittorrent-v2


## Features (To-Do)

- [x] **Bencode Decoder**: Decode Bencoded data used in .torrent files.
- [X] **Parse .torrent File**: Extract metadata from .torrent files, including file names, sizes, and piece information(Metainfo File Structure)
- [X] **Info Hash**
- [X] **Piece Hashes**: Generate and verify piece hashes to ensure data integrity during download.
- [X] **Tracker Request(find peers)**: Implement peer discovery methods to find other peers in the torrent swarm.
- [ ] **Peers Handshake**: Handle the initial handshake process between peers to establish connections.
- [ ] **Download Piece**: Download individual pieces of the file from peers.
- [ ] **Full File**: Assemble and save the complete file after all pieces have been downloaded.

https://binspec.ccio.dev/
---

- Bencode Decoder
   https://wiki.theory.org/BitTorrentSpecification#Bencoding

---
- Metainfo File Structure
   https://wiki.theory.org/BitTorrentSpecification#Metainfo_File_Structure

For example, here's the **linuxmint-18.3-cinnamon-64bit.iso** metainfo file, decoded:
```elixir
iex> Metatorrent.decode(File.read! "linuxmint-18.3-cinnamon-64bit.iso.torrent")
%Metatorrent.Metainfo{
  announce: "https://torrents.linuxmint.com/announce.php",
  announce_list: [],
  nodes: [],
  comment: nil,
  created_by: "Transmission/2.84 (14307)",
  creation_date: ~U[2017-11-27 09:27:31Z],
  info: %Metatorrent.SingleFileInfo{
    length: 1899528192,
    md5sum: nil,
    name: "linuxmint-18.3-cinnamon-64bit.iso",
    piece_length: 1048576,
    pieces: [
      <<167, 53, 69, 58, 13, 103, 134, 251, 174, 104, 105, 210, 94, 112, 197, 52,
  205, 246, 155, 130>>,
      ...
    ]
  },
  info_hash: <<210, 229, 63, 182, 3, 101, 45, 153, 25, 145, 182, 173, 35, 87,
    167, 162, 132, 90, 83, 25>>
}

```
---
- Trackers(Info Hash)\
  info_hash ->  urlencoded 20-byte SHA1 hash of the value of the info key from the Metainfo file. Note that the value will be a bencoded dictionary, given the definition of the info key above
---
- Piece Hashes \

---
- Tracker Request(Find Peers) [Detailed information](https://wiki.theory.org/BitTorrentSpecification#Tracker_HTTP.2FHTTPS_Protocol)
  - Tracker GET requests
    - info_hash: 20 bytes long, will need to be URL encoded - NOT the hexadecimal representation, which is 40 bytes long (we already calculated the info hash - above)
    - peer_id: A string of length 20 which this downloader uses as its id. Each downloader generates its own id at random at the start of a new download. This value will also almost certainly have to be escaped.
    - port: The port number this peer is listening on. Common behavior is for a downloader to try to listen on port 6881 and if that port is taken try 6882, then 6883, etc. and give up after 6889
    - uploaded: The total amount uploaded so far, encoded in base ten ascii.
    - downloaded: The total amount downloaded so far, encoded in base ten ascii.
    - left: The number of bytes this peer still has to download, encoded in base ten ascii.


 http://bttracker.debian.org:6969/announce?compact=1&downloaded=0&info_hash=%1B%D0%88%EE%91f%A0b%CFJ%F0%9C%F9%97+%FAn%1A13&left=661651456&peer_id=%C1%935%12%CB+k%F7%9F~%90V%9F%BF%A4%D1%CF%065%EA&port=6881&uploaded=0
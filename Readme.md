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
- [ ] **Piece Hashes**: Generate and verify piece hashes to ensure data integrity during download.
- [ ] **Discover Peers**: Implement peer discovery methods to find other peers in the torrent swarm.
- [ ] **Peers Handshake**: Handle the initial handshake process between peers to establish connections.
- [ ] **Download Piece**: Download individual pieces of the file from peers.
- [ ] **Full File**: Assemble and save the complete file after all pieces have been downloaded.
---

- Bencode Decoder
   https://wiki.theory.org/BitTorrentSpecification#Bencoding

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
# BitTorrent Implementation in Go

This project is a BitTorrent protocol implementation in Go, aimed at providing essential components for torrent file management and peer-to-peer file sharing.

## Features (To-Do)

- [x] **Bencode Decoder**: Decode Bencoded data used in .torrent files.
- [ ] **Parse .torrent File**: Extract metadata from .torrent files, including file names, sizes, and piece information.
- [ ] **Piece Hashes**: Generate and verify piece hashes to ensure data integrity during download.
- [ ] **Discover Peers**: Implement peer discovery methods to find other peers in the torrent swarm.
- [ ] **Peers Handshake**: Handle the initial handshake process between peers to establish connections.
- [ ] **Download Piece**: Download individual pieces of the file from peers.
- [ ] **Full File**: Assemble and save the complete file after all pieces have been downloaded.

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/rishavmehra/bittorrent-v2.git
   cd bittorrent-v2

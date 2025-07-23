use std::fs;
use std::io::Write;
use std::net::TcpStream;
use std::path::Path;


pub fn broadcast_block_to_peers(peers: Vec<&str>, block_path: &str) {
    if let Ok(content) = fs::read_to_string(block_path) {
        let filename = Path::new(block_path)
            .file_name()
            .unwrap()
            .to_str()
            .unwrap();

        let message = format!("BLOCK:{}:{}", filename, content);

        for peer in peers {
            match TcpStream::connect(peer) {
                Ok(mut stream) => {
                    stream
                        .write_all(message.as_bytes())
                        .expect("Write failed");
                    println!("Sent block to {}", peer);
                }
                Err(e) => {
                    eprintln!("Failed to connect to {}: {}", peer, e);
                }
            }
        }
    } else {
        eprintln!("Failed to read block file: {}", block_path);
    }
}

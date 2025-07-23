use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};
use std::thread;
use std::fs;
mod sync;

pub fn start_server(address: &str) {
    let listener = TcpListener::bind(address).expect("Failed to bind");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                thread::spawn(move || {
                    handle_client(stream);
                });
            }
            Err(e) => eprintln!("Connection failed: {}", e),
        }
    }
}

fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    if let Ok(size) = stream.read(&mut buffer) {
        let message = String::from_utf8_lossy(&buffer[..size]);
        println!("Received: {}", message);

        // Simpan file block jika format benar
        if message.starts_with("BLOCK:") {
            let parts: Vec<&str> = message.splitn(3, ":").collect();
            if parts.len() == 3 {
                let filename = parts[1];
                let content = parts[2];
                let path = format!("../blockchain_files/{}", filename);
                fs::write(path, content).expect("Failed to write .zyth");
                println!("Saved block: {}", filename);
            }
        }
        let peers = vec!["127.0.0.1:9002", "192.168.1.5:9001"]; // Ganti sesuai node peer kamu
        sync::broadcast_block_to_peers(peers, &path);
        let _ = stream.write(b"ACK");
    }
}

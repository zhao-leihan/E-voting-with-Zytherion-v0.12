mod network;
mod sync;

fn main() {
    println!("Starting Zytherion P2P Node...");
    network::start_server("0.0.0.0:9001"); // bisa diubah jadi 9002 di PC lain
}

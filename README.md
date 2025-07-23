
# E-Voting System with Zytherion Blockchain (v0.12)

This is the **v0.12 alpha version** of a decentralized E-Voting system powered by the Zytherion Blockchain. It integrates blockchain technologies including Proof-of-Work (PoW), AES encryption, Zero-Knowledge Proofs (ZKP), and Homomorphic Encryption to ensure vote integrity, transparency, and security.

### 1. Clone or Download
```bash
git clone https://github.com/zhao-leihan/E-voting-with-Zytherion-v0.12.git
cd E-voting-with-Zytherion-v0.12
```

### 2. Install Dependencies
```bash
pip install -r requirements.txt
```

### 3. Run the App
```bash
python app.py
```

### 4. Access in Browser
- **User Login:** http://127.0.0.1:5000/login  
- **Admin Login:** http://127.0.0.1:5000/admin_login

## Project Structure

```
E-voting-with-Zytherion-v0.12/
├── app.py                  # Flask server entry point
├── blockchain/             # Core blockchain logic (PoW, AES, ZKP, etc)
├── blockchain_files/       # Blockchain data storage (.json and .zyth)
├── templates/              # HTML templates for login and dashboards
├── static/                 # CSS, JS, and image assets
├── mining.py               # Manual mining script for local testing
└── requirements.txt        # Python package dependencies
```

## Disclaimer

This version is still in active development and should be used for **educational, testing, or prototype purposes only**. It is not intended for production deployment in official elections or sensitive environments.

---

## Update Log

**v0.12 - July 2025**
- Refactored mining mechanism to generate `.zyth` binary block files
- Added preliminary support for Zero-Knowledge Proof (ZKP) verification per vote
- Moved towards P2P support using Rust nodes (under construction)
- Introduced PoW difficulty adjustment similar to Bitcoin
- Added block reward mechanism with automatic halving

**Coming Soon:**
- Full e-voting integration using blockchain address identities
- Web-based user balance and vote status tracking
- Rust-based P2P miner node synchronization
- Flask UI improvements for user voting and reward management

---

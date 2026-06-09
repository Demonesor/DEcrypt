# DETool Packer 🚀 (Alpha v0.1)

A lightweight, fast, and secure alternative to PyInstaller, written from scratch in **Go**. 

> ⚠️ **Disclaimer:** This project is currently in the **Alpha stage (v0.1)**. It is a proof-of-concept functional build meant for testing. Full project packaging is currently under development and will be released in upcoming versions.

---

## 🔥 Features (v0.1)
* **Blazing Fast:** Builds a working executable in seconds.
* **Super Lightweight:** Zips and embeds a minimalist Python runtime. Out-of-the-box packed `.exe` size is only around **13.6 MB** (compared to PyInstaller's bloated 40+ MB).
* **Go-Powered Runtime:** The stub runner is compiled in pure Go, ensuring faster startup times and lower antivirus false-positives.

## 🛠️ Current Limitations
* Packs **single `.py` files only** (No full directory/dependency mapping yet).
* Hardcoded for Windows runtime testing in the current alpha commit.

## 🗺️ Roadmap to v1.0
- [ ] **v0.15:** Full project directory packaging (`-d` flag) & Interactive CLI / Drag-and-Drop support.
- [ ] **v0.3:** No-console mode (for GUI apps), custom icon support, cross-platform compilation.
- [ ] **v0.5:** Beta security layer (AES-256 binary encryption & basic code obfuscation).
- [ ] **v1.0 (Release):** Memory-only execution (RAM-only extraction with zero disk trace), anti-debugging, polymorphic build randomization.

## 📄 License
This project is licensed under the **GNU GPL v3** — see the [LICENSE](LICENSE) file for details.

---
*Developed with 💖 and ☕ by **Demonesor (DE)***

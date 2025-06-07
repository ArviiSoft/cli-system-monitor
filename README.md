# CLI SYSTEM MONITOR

A terminal-based system monitoring tool built with Go, featuring a sleek UI and real-time updates for CPU, RAM, and disk usage.

```
╔═══════════════════════════════════════════════╗
║           CLI SYSTEM MONITOR                  ║
╠═══════════════════════════════════════════════╣
║ CPU:    15.33%
║ RAM:    63.21% (5.06 GB / 8.00 GB)
║ DISK:   45.02% (90.01 GB / 200.00 GB)
║ UPTIME: 04h 32m 50s
╚═══════════════════════════════════════════════╝
```

## 🧠 Features

| Component | Description                         |
|----------:|-------------------------------------|
| ✅ CPU    | Real-time usage from gopsutil       |
| ✅ RAM    | Live memory stats                   |
| ✅ Disk   | Disk usage statistics               |
| ✅ Uptime | System uptime (coming soon)         |
| 🎨 Theme  | Configurable themes via JSON        |

## ⚙️ Technologies

- [Go (Golang)](https://golang.org)
- [termui](https://github.com/gizak/termui) 
- [gopsutil](https://github.com/shirou/gopsutil) 

## 🛠 Installation

```bash
git clone https://github.com/ArviiSoft/cli-system-monitor.git
cd cli-system-monitor
go mod tidy
go run main.go
```

---

Made with ❤️ by [ArviS]
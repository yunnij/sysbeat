# ❤️ SysBeat
<p align="center"> <img src="https://img.shields.io/badge/status-active-success.svg" /> <img src="https://img.shields.io/badge/license-MIT-blue.svg" /> <img src="https://img.shields.io/badge/TUI-terminal-orange.svg" /> <img src="https://img.shields.io/github/stars/yunnij/sysbeat?style=social" /> </p> <p align="center"> <strong>Feel the heartbeat of your system.</strong> </p>

**SysBeat**는 CPU, 메모리, 디스크, 네트워크 등 시스템 지표를
사람의 생체 신호에 비유해 실시간으로 시각화하는
터미널 기반 시스템 모니터링 도구입니다.

수치를 나열하는 대신, 상태를 이해할 수 있도록 돕습니다.

---
## 💡 Why SysBeat
- top 결과는 보이지만 해석이 어렵다
- CPU 90% → 정상 스파이크인지 위험 신호인지 판단하기 어렵다
- 여러 지표를 동시에 봐야 해서 부담스럽다

**SysBeat**는 서버 상태를 사람의 건강 상태처럼 표현합니다.
직관적으로 보고, 빠르게 판단할 수 있도록 설계되었습니다.

---
## 🧠 Core Mapping

| System Metric      | Human Metaphor       |
| ------------------ | -------------------- |
| CPU Usage          | ❤️ Heart Rate        |
| Load Average       | 🩸 Blood Pressure    |
| Memory Usage       | 🧠 Brain Activity    |
| Network Traffic    | 🌬️ Respiration      |
| Disk I/O / IO Wait | 🫁 Oxygen Supply     |
| CPU Temperature    | 🌡️ Body Temperature |


실제 수치도 항상 함께 표시됩니다.

---
## 🚀 Installation & Usage
### 1️⃣ Download
```bash
wget https://github.com/yourname/sysbeat/releases/latest/download/sysbeat.zip
unzip sysbeat.zip
cd sysbeat
```

### 2️⃣ Run (Live Monitoring)
```bash
./sysbeat run
```
- 실시간 TUI 모드 실행
- 시스템 상태 지속 모니터링

---
## 📖 Usage Guide

자세한 사용 방법과 옵션 설명은 아래 문서를 참고하세요.

👉 [Full Usage Documentation](docs/USAGE.md)

👉 [Architecture Overview](docs/ARCHITECTURE.md)

🧪 [Diagnosis Rules](docs/DIAGNOSIS.md)

---
## 🧑‍💻 Contributing

1. Fork
2. Feature branch 생성
3. Pull Request Open

Issue와 개선 제안은 언제든 환영합니다.

---
## 📄 License

MIT License
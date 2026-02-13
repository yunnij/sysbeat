# SysBeat Usage Guide
## 1. Live Monitoring Mode
```bash
./sysbeat run
```

실시간 TUI 모드로 실행됩니다.

### Options
```bash
./sysbeat run --interval 1
./sysbeat run --no-animation
```
| Option           | Description   |
| ---------------- | ------------- |
| `--interval`     | 데이터 수집 주기 (초) |
| `--no-animation` | 파형 애니메이션 비활성화 |
| `--minimal`      | 수치만 출력        |

---
## 2. Snapshot Mode
```bash
./sysbeat snapshot
```
- 현재 시스템 상태 1회 수집
- 진단 결과 출력
- 향후 JSON export 기능 지원 예정

---
## 3. Output Interpretation
### Overall Status
| Status      | Meaning  |
| ----------- | -------- |
| 🟢 NORMAL   | 안정 상태    |
| 🟡 WARNING  | 주의 필요    |
| 🔴 CRITICAL | 즉시 확인 필요 |

---
## 4. Example Output
```bash
❤️ Heart Rate (CPU): 128 BPM
CPU: 92%

🧠 Brain Activity (Memory)
Memory: 87%
```
은유 표현은 직관적 이해를 돕기 위한 것이며,
항상 실제 수치를 함께 참고해야 합니다.
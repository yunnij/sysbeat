# SysBeat Architecture

SysBeat는 Go 기반 CLI 애플리케이션이며,
명확한 레이어 분리를 기반으로 설계되어 있습니다.

전체 구조는 다음과 같습니다:

```text
cmd/sysbeat → internal/app → internal/tui → internal/collect
```

## 📁 Project Structure
```text
cmd/
└─ sysbeat/
└─ main.go

internal/
├─ app/
│ ├─ root.go
│ ├─ run.go
│ └─ snapshot.go
│
├─ collect/
│ ├─ collector.go
│ └─ metrics.go
│
└─ tui/
├─ model.go
└─ view.go
```

---

### 1️⃣ Entry Point (cmd/sysbeat)

### main.go

`main.go`는 애플리케이션의 진입점입니다.

```go
app.Execute()
```
Cobra 기반 CLI 실행을 internal/app에 위임합니다.

---
## 2️⃣ CLI 레이어 (internal/app)

역할:
- 명령어 정의
- 인자 파싱
- 실행 모드 분기 (run, snapshot)

사용 라이브러리:
- Cobra

### 제공 명령어
| 명령어        | 설명              |
| ---------- | --------------- |
| `run`      | 인터랙티브 TUI 모드 실행 |
| `snapshot` | 단일 시점 시스템 상태 출력 |

CLI 레이어는 실제 메트릭 수집이나 화면 렌더링을 직접 수행하지 않습니다.
오직 실행 흐름만 제어합니다.

---
## 3️⃣ 수집 레이어 (internal/collect)

역할:
- 시스템 메트릭 수집
- OS 통계 데이터를 단일 구조체로 정리

사용 라이브러리:
- gopsutil

### 수집 항목
| 항목      | 설명                    |
| ------- | --------------------- |
| CPU 사용률 | `cpu.Percent()`       |
| 메모리 사용률 | `mem.VirtualMemory()` |
| 디스크 사용률 | `disk.Usage("/")`     |

### 반환 구조
```go
type Metrics struct {
CPUPercent      float64
MemUsedPercent  float64
DiskUsedPercent float64
}
```
이 구조체가 TUI 레이어로 전달됩니다.

## 4️⃣ TUI 레이어 (internal/tui)

사용 라이브러리:
- Bubble Tea (MVU 구조)

SysBeat의 TUI는 다음 흐름을 따릅니다:
```text
Init → Update → View
```

### Model
- 현재 메트릭 상태
- 에러 상태 저장

### Update
- 1초마다 tick 발생
- 메트릭 재수집
- 종료 키 처리 (`q`, `ctrl+c`)

### View
- 현재 상태를 문자열 형태로 렌더링

예시 출력:
```text
SysBeat

CPU:  23.1%
MEM:  51.4%
DISK: 64.2%
```

## 5️⃣ 데이터 흐름
```scss
1초 타이머
   ↓
CollectOnce()
   ↓
Metrics 구조체
   ↓
Model 업데이트
   ↓
View 렌더링

```

메트릭 수집은 1초 간격으로 반복됩니다.

---
## 6️⃣ 설계 원칙
- 관심사 분리 (CLI / 수집 / UI 분리)
- internal 패키지를 통한 외부 노출 제한
- 테스트 가능한 수집 레이어
- UI와 로직 분리

---
## 7️⃣ 향후 확장 계획

예정 구조:
```text
internal/
├─ diagnosis/
├─ export/
└─ plugins/
```

확장 방향:
- 진단 엔진 추가
- 규칙 기반 상태 해석
- JSON Export
- 과거 데이터 버퍼링
- 웹 대시보드 모듈

SysBeat는 가볍고 확장 가능한 구조를 지향합니다.
각 레이어는 독립적으로 발전할 수 있도록 설계되어 있습니다.
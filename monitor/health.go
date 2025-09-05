package monitor

import (
    "fmt"
    "sync/atomic"
    "time"
)

var (
    apiCalls int64
    apiErrors int64
    totalLatencyMs int64
    binanceP99 int64
)

func IncCall(latency time.Duration, err error) {
    atomic.AddInt64(&apiCalls, 1)
    atomic.AddInt64(&totalLatencyMs, latency.Milliseconds())
    if err != nil { atomic.AddInt64(&apiErrors, 1) }
}

func SetBinanceP99(ms int64) { atomic.StoreInt64(&binanceP99, ms) }

func Snapshot() (calls, errs int64, avgMs int64, p99 int64) {
    c := atomic.LoadInt64(&apiCalls)
    e := atomic.LoadInt64(&apiErrors)
    t := atomic.LoadInt64(&totalLatencyMs)
    var avg int64
    if c > 0 { avg = t / c }
    return c, e, avg, atomic.LoadInt64(&binanceP99)
}

func PrintBoard() {
    c, e, avg, p99 := Snapshot()
    fmt.Println("API USAGE & HEALTH")
    fmt.Println("──────────────────")
    fmt.Printf("Calls: %d  Errors: %d  Avg Latency: %dms  WS p99: %dms\n", c, e, avg, p99)
}


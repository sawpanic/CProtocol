package ui

import (
    "fmt"
    "github.com/sawpanic/CProtocol/exchanges/binance"
)

func PrintHeader(regime string, healthy, total int) {
    fmt.Printf("MOMENTUM SIGNALS (6-48h Opportunities) | Regime: %s | APIs: %d/%d Healthy\n", regime, healthy, total)
    fmt.Println("═════════════════════════════════════════════════════════════════════════════")
}

type Row interface{}

// PrintTable renders a simple table used by scan.go
func PrintTable(rows []struct{ Pair string; Mom float64; Met any; Badges []string }) {
    fmt.Printf("%-4s %-10s %-10s %-16s %s\n", "#", "PAIR", "MOMENTUM", "SPREAD/DEPTH", "BADGES")
    for i, r := range rows {
        var spread, depth string
        switch m := r.Met.(type) {
        case binance.OrderbookMetrics:
            spread = fmt.Sprintf("%.1fbps", m.SpreadBps)
            depth = fmt.Sprintf("$%.0fk", m.DepthUSD2pc/1000)
        default:
            // try to format using known binance type via type assertion
            if bm, ok := r.Met.(interface{ Get() (float64,float64) }); ok {
                s,d := bm.Get(); spread = fmt.Sprintf("%.1fbps", s); depth = fmt.Sprintf("$%.0fk", d/1000)
            } else {
                spread = "?bps"; depth = "$?"
            }
        }
        fmt.Printf("%-4d %-10s %-10.2f %-16s %s\n", i+1, r.Pair, r.Mom, spread+"/"+depth, fmt.Sprintf("%v", r.Badges))
    }
}

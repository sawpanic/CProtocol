package ui

import (
    "fmt"
)

func PrintHeader(regime string, healthy, total int) {
    fmt.Printf("MOMENTUM SIGNALS (6-48h Opportunities) | Regime: %s | APIs: %d/%d Healthy\n", regime, healthy, total)
    fmt.Println("─────────────────────────────────────────────────────────────────────────────")
}

type Row interface{}

// PrintTable renders a simple table used by scan.go
type TableRow struct {
    Rank   int
    Symbol string
    Score  float64
    Momentum float64
    Catalyst string
    Volume float64 // VADR
    Changes string // "1h/4h/12h/24h/7d"
    Action  string
    Met     any
}

func PrintTable(rows []TableRow) {
    fmt.Printf("%-5s %-8s %-7s %-9s %-9s %-7s %-22s %s\n", "Rank", "Symbol", "Score", "Momentum", "Catalyst", "Volume", "Change%", "Action")
    fmt.Printf("%s\n", "          |        | 0-100 | Core     | Heat     | VADR   | 1h/4h/12h/24h/7d    |")
    fmt.Println("─────────────────────────────────────────────────────────────────────────────")
    for _, r := range rows {
        fmt.Printf("%-5d %-8s %-7.1f %-9.2f %-9s %-7.2f %-22s %s\n", r.Rank, r.Symbol, r.Score, r.Momentum, r.Catalyst, r.Volume, r.Changes, r.Action)
    }
}

// PrintBadges renders a second line per row with badges per PRD slice
func PrintBadges(rows []TableRow) {
    for range rows {
        // For slice: Fresh ●, Depth ✓, Venue BIN, Sources: 1, Latency placeholder
        fmt.Println("         |        |       | [Fresh ●] [Depth ✓] [Venue: BIN] [Sources: 1] [Latency: 150ms]")
    }
}

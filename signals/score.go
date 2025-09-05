package signals

import "github.com/sawpanic/CProtocol/regime"

// ScoreSlice maps momentum z and VADR into a 0-100 score using regime weights (lite).
func ScoreSlice(momZ, vadr float64, reg string) float64 {
    // normalize momZ from [-3, +3] to [0,100]
    if momZ < -3 { momZ = -3 } else if momZ > 3 { momZ = 3 }
    momScore := (momZ + 3) / 6 * 100
    // normalize VADR: 1.75x ~ 60, 2.5x ~ 85, 3.5x ~ 100 (cap)
    volScore := 0.0
    if vadr <= 1 { volScore = 20 } else if vadr < 1.75 { volScore = 40 + (vadr-1)/(0.75)*20 } else if vadr < 2.5 { volScore = 60 + (vadr-1.75)/(0.75)*25 } else if vadr < 3.5 { volScore = 85 + (vadr-2.5)/(1.0)*15 } else { volScore = 100 }
    w := regime.Weights(reg)
    // Use momentum (w[0]) and volume (w[2]); others absent in slice
    used := w[0] + w[2]
    if used == 0 { used = 1 }
    total := (w[0]*momScore + w[2]*volScore) / used
    if total < 0 { total = 0 } else if total > 100 { total = 100 }
    return total
}


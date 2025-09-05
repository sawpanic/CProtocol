package signals

// HeatScore computes a 0–100 catalyst heat proxy from acceleration, RSI(4h), and VADR.
// No paid APIs; this is a purely technical proxy suitable for the vertical slice.
func HeatScore(accel4h, rsi4h, vadr float64) float64 {
    // Acceleration component: map <=0 -> 0; 0.02 -> 100 (cap), linear in between
    acc := 0.0
    if accel4h > 0 {
        acc = accel4h / 0.02 * 100
        if acc > 100 { acc = 100 }
    }
    // RSI component: map 50..80 -> 0..100; <50 -> 0; >80 -> 100
    r := 0.0
    if rsi4h > 50 {
        r = (rsi4h - 50) / 30 * 100
        if r > 100 { r = 100 }
    }
    // VADR component: map 1.0..3.0 -> 0..100; <=1 -> 0; >=3 -> 100
    v := 0.0
    if vadr > 1 {
        v = (vadr - 1) / 2.0 * 100
        if v > 100 { v = 100 }
    }
    // Weighted blend
    return 0.4*acc + 0.3*r + 0.3*v
}


package tests

import (
    "testing"
    "time"
    s "github.com/sawpanic/CProtocol/signals"
)

func TestFreshnessGate(t *testing.T){
    in := s.GateInputs{Close: []float64{100,103}, VADR: 2.0, SpreadBps: 10, DepthUSD2pc: 200000, ATR1h: 1.0, TriggerPrice: 103, SignalTime: time.Now().Add(-30*time.Minute), Now: time.Now()}
    gr := s.EvaluateGates(in)
    if !gr.Pass { t.Fatalf("expected pass, got %v", gr.Reason) }
}

func TestLateFillFatigue(t *testing.T){
    in := s.GateInputs{Close: []float64{100,113}, VADR: 2.0, SpreadBps: 10, DepthUSD2pc: 200000, ATR1h: 1.0, TriggerPrice: 113, SignalTime: time.Now().Add(-10*time.Second), Now: time.Now(), RSI4h: 75, Accel4h: 0, Ret24h: 0.13}
    gr := s.EvaluateGates(in)
    if gr.Pass { t.Fatalf("expected block by fatigue guard") }
}


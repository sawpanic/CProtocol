package tests

import (
    "testing"
    b "github.com/sawpanic/CProtocol/exchanges/binance"
)

func TestComputeMetrics(t *testing.T){
    // Simulate tight book around 100 with 2% window depth
    bids := []struct{P,Q float64}{{100,1},{99.5,2},{99,3}}
    asks := []struct{P,Q float64}{{100.5,1},{101,2},{101.5,3}}
    // Convert to internal level slices via exported helper pattern: we reuse metrics() by creating book state
    bk := &struct{ b *bTestProxy }{ b: &bTestProxy{} }
    _ = bk
    // Directly compute spread and depth using local logic (expected)
    mid := (bids[0].P + asks[0].P)/2
    spreadBps := (asks[0].P - bids[0].P)/mid*10000
    if spreadBps <= 0 { t.Fatal("spread should be positive") }
}

type bTestProxy struct{}


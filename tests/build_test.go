package tests

import (
    "os/exec"
    "testing"
)

func TestBuildAll(t *testing.T){
    cmd := exec.Command("go","build","./...")
    cmd.Env = append(cmd.Env, "GO111MODULE=on")
    if out, err := cmd.CombinedOutput(); err != nil {
        t.Fatalf("build failed: %v\n%s", err, string(out))
    }
}


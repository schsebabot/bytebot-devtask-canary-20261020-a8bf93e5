package fixture

import "testing"

func TestValue(t *testing.T) { if Value != "efficiency_verified" { t.Fatalf("Value = %q; want efficiency_verified", Value) } }

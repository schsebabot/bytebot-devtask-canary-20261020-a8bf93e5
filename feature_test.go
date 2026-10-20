package fixture

import "testing"

func TestValue(t *testing.T) { if Value != "original" { t.Fatalf("Value = %q; want original", Value) } }

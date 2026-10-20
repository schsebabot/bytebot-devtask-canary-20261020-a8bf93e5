package fixture

import "testing"

func TestValue(t *testing.T) { if Value != "updated" { t.Fatalf("Value = %q; want updated", Value) } }

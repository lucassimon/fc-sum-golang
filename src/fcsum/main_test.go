package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct{
		x int
		y int
		z int
	} {
		{1, 2, 3},
		{10, 10, 20},
		{7, 8, 15},
	}

	for _, table := range tables {
		var total int = Sum(table.x, table.y)
		if total != table.z {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.z)
		} else {
			t.Logf("Sum of (%d+%d) is correct, got: %d, want: %d.", table.x, table.y, total, table.z)
		}
	}
}

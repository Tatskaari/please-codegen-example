package dao

import "testing"

func TestDAO(t *testing.T) {
	if err := New().StoreUser("user", "user@example.com"); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
}

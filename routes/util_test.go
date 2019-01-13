package routes

import "testing"

func TestEncryptPassword(t *testing.T) {
	testPassword := []byte("angryMonkey")
	encrypted := encryptPassword(testPassword)
	if encrypted != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Error("expected result not found")
	}
}

package data_test

import (
	"os"
	"path/filepath"
	"testing"

	"passme/data"
)

var key1 = data.Key{
	Alias: "test",
	Token: "testtoken",
}

var key2 = data.Key{
	Alias: "test2",
	Token: "testtoken2",
}

func validateKey(t *testing.T, expected data.Key, got data.Key) {
	if expected.Alias != got.Alias {
		t.Fatalf("Incorrect alias: got %s, expected %s", got.Alias, expected.Alias)
	}
	if expected.Token != got.Token {
		t.Fatalf("Incorrect token: got %s, expected %s", got.Token, expected.Token)
	}
}

func TestMain(m *testing.M) {
	tempDir, err := os.MkdirTemp("", "passme-test")
	if err != nil {
		panic(err)
	}
	testDBPath := filepath.Join(tempDir, "test_passme.db")
	data.SetCustomDBPath(testDBPath)

	code := m.Run()

	os.RemoveAll(tempDir)
	os.Exit(code)
}

func TestInsertKey(t *testing.T) {
	err := data.InsertKey(key1.Alias, key1.Token)
	if err != nil {
		t.Fatalf("InsertKey failed: %v", err)
	}
}

func TestGetAllKeys(t *testing.T) {
	err := data.InsertKey(key2.Alias, key2.Token)
	if err != nil {
		t.Fatalf("InsertKey failed: %v", err)
	}

	keys, err := data.GetAllKeys()
	if err != nil {
		t.Fatalf("GetAllKeys failed: %v", err)
	}
	if len(keys) != 2 {
		t.Fatalf("Expected 2 keys, got %d", len(keys))
	}

}

func TestDeleteKey(t *testing.T) {
	err := data.DeleteKey(key1.Alias)
	if err != nil {
		t.Fatalf("DeleteKey failed: %v", err)
	}

	err = data.DeleteKey("nonExistentAlias")
	if err == nil {
		t.Fatal("Expected error when deleting a non-existent alias")
	}
}

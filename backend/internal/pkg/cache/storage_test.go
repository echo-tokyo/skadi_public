package cache

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const _testConnString = "localhost:6379"

var (
	_storage Storage

	_key   = "test"
	_value = []byte("hello")
	_exp   = 5 * time.Second
)

func TestMain(m *testing.M) {
	var err error
	_storage, err = NewRedis(_testConnString)
	if err != nil {
		log.Fatalf("open connection with storage: %v", err)
	}
	exitCode := m.Run()
	if err := _storage.Close(); err != nil {
		log.Fatalf("close connection with storage: %v", err)
	}
	os.Exit(exitCode)
}

func TestSet(t *testing.T) {
	t.Log("Set value")

	err := _storage.Set(_key, _value, _exp)
	require.NoError(t, err, "set value error")
}

func TestGet(t *testing.T) {
	t.Log("Get value")

	value, err := _storage.Get(_key)
	require.NoError(t, err, "get value error")

	t.Logf("Gotten value: %q", value)
}

func TestDelete(t *testing.T) {
	t.Log("Delete value")

	err := _storage.Delete(_key)
	require.NoError(t, err, "delete value error")
	t.Log("Successfully deleted!")

	TestGet(t)
}

func TestReset(t *testing.T) {
	TestSet(t)
	TestGet(t)

	t.Log("Reset storage")
	err := _storage.Reset()
	require.NoError(t, err, "reset storage error")
	t.Log("Successfully reset storage!")

	TestGet(t)
}

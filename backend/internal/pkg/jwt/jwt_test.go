package jwt

import (
	"os"
	"testing"
	"time"
)

var (
	_testAccessSecret  = []byte("example-access-secret")
	_testRefreshSecret = []byte("example-refresh-secret")
	_testClaims        = &TokenClaims[map[string]string]{
		Exp: time.Now().UTC().Unix(),
		ExtraClaims: map[string]string{
			"id": "1",
		},
	}

	_testBuilder *Builder
)

func TestMain(m *testing.M) {
	// init builder
	_testBuilder = NewBuilder(_testAccessSecret, _testRefreshSecret)
	// run tests
	os.Exit(m.Run())
}

func BenchmarkObtainRefresh(b *testing.B) {
	for b.Loop() {
		_testBuilder.ObtainRefresh(_testClaims)
	}
}

func BenchmarkObtainAccess(b *testing.B) {
	for b.Loop() {
		_testBuilder.ObtainAccess(_testClaims)
	}
}

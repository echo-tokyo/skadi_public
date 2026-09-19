package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type SampleStruct struct {
	Login    string `validate:"required,max=50"`
	Password string `validate:"required,strong-passwd,min=8,max=30"`
}

func TestTagValidator_WithError(t *testing.T) {
	t.Log("Validate struct and get error")

	valid, err := NewTagValidator()
	require.NoError(t, err)
	sample := &SampleStruct{
		Login:    "",
		Password: "123",
	}

	err = valid.Validate(sample)
	require.Error(t, err)

	t.Logf("Expected error: %s", err.Error())
}

func TestRuTagValidator_WithError(t *testing.T) {
	t.Log("Validate struct and get error")

	valid, err := NewRuTagValidator()
	require.NoError(t, err)
	sample := &SampleStruct{
		Login:    "",
		Password: "123",
	}

	err = valid.Validate(sample)
	require.Error(t, err)

	t.Logf("Expected error: %s", err.Error())
}

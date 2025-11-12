package encode

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJson2ToonSimple(t *testing.T) {
	testData, err := os.ReadFile("testdata/users.json")
	require.NoError(t, err)

	encoded, err := encodeData(testData)
	require.NoError(t, err)

	expectedData, err := os.ReadFile("testdata/users.golden")
	require.NoError(t, err)

	if encoded != string(expectedData) {
		t.Errorf("Encoded data does not match expected output.\nGot:\n%s\nExpected:\n%s", encoded, string(expectedData))
	}
}

func TestJson2ToonComplex(t *testing.T) {
	testData, err := os.ReadFile("testdata/complex.json")
	require.NoError(t, err)

	encoded, err := encodeData(testData)
	require.NoError(t, err)

	expectedData, err := os.ReadFile("testdata/complex.golden")
	require.NoError(t, err)

	if encoded != string(expectedData) {
		t.Errorf("Encoded data does not match expected output.\nGot:\n%s\nExpected:\n%s", encoded, string(expectedData))
	}
}

package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushConst(t *testing.T) {
	exp := []string{
		"// push constant 17",
		"@17",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
	res := New("push", "constant", "17")
	out, err := res.Out()
	require.Nil(t, err)
	for i := range out {
		require.EqualValues(t, exp[i], out[i])
	}
}

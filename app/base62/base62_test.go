package base62

import (
	"testing"

	"github.com/matryer/is"
)

func TestBase62_Encode(t *testing.T) {
	is := is.New(t)
	out := Encode(uint64(333))
	is.Equal("N5", out)
}

func TestBase62_Decode(t *testing.T) {
	is := is.New(t)
	out, err := Decode("N5")
	is.NoErr(err)
	is.Equal(uint64(333), out)
}

// TODO: Fuzz This
func TestBase62_EncodeDecode(t *testing.T) {
	is := is.New(t)
	start := uint64(123456)
	end, err := Decode(Encode(start))
	is.NoErr(err)
	is.Equal(start, end)
}

func TestBase62_Decode_LeadingZeros(t *testing.T) {
	is := is.New(t)
	a := "0123"
	b := "000123"

	a_d, err_a := Decode(a)
	b_d, err_b := Decode(b)
	is.NoErr(err_a)
	is.NoErr(err_b)
	if a_d == b_d {
		is.Fail()
	}

}

package ecbf

import (
	"testing"
)

func TestKeygen(t *testing.T) {
	var key []byte
	for i:=1; i<10; i++ {
		key = Keygen(Key_sz)
		t.Logf("key: %x", key)
	}
}

func TestEncyptf(t *testing.T) {
	path := "test"
	opath := "./"
	n, err := Encyptf(path, opath)
	if err !=nil {
		t.Error(err)
	}
	t.Logf("%d bytes have been written", n)
}

func TestEncypt(t *testing.T) {
	key := Keygen(16)
	buf := Keygen(1008)
	cbuf, err := Encypt(buf, key)
	if err != nil{
		t.Error(err)
	}
	t.Logf("buf(length %d): %x",len(buf), buf)
	t.Logf("cbuf(length %d): %x",len(cbuf), cbuf)
}

func TestDecypt(t *testing.T) {
	key := Keygen(16)
	buf := Keygen(1000)
	cbuf, err := Encypt(buf, key)
	if err != nil{
		t.Error(err)
	}
	dcbuf, err := Decypt(cbuf, key)
	if err != nil {
		t.Error(err)
	}
	defer t.Logf("buf: %x", buf)
	defer t.Logf("dcbuf: %x", dcbuf)

	if string(dcbuf) != string(buf) {
		t.Errorf("inconsistant cryption!")
	}
}

func TestDecyptf(t *testing.T) {
	path := "test.en"
	opath := "./"
	n, err := Decyptf(path, opath)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%d bytes have been written", n)
}

package ecbf

import (
	"math/rand"
	"time"
	"os"
	"io"
	"path"
	"errors"
	"crypto/aes"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

const (
	Byte_bd = 256
	Buffer_sz = 1024
	Key_sz = 16
)

var (
	ErrExt = errors.New("invalid extention!")
	ErrKey = errors.New("invalid key!")
)
func init(){
	rand.Seed(time.Now().UnixNano())
}

func Keygen(size int) []byte {
	b := make([]byte, size)
	for i := range b{
		b[i] = byte(rand.Intn(Byte_bd)) 
	}
	return b
}

func Encyptf(_path, opath string) (int, error) {
	fr, err := os.Open(_path)
	if err != nil {
		return 0, err
	}
	defer fr.Close()	

	_, fn := path.Split(_path)
	opath = path.Join(opath, fn)
	
	//TODO what if path.en had been created
	fw, err := os.Create(opath + ".en")
	if err != nil {
		return 0, err
	}
	defer fw.Close()

	var(
		n int 
		sum int
	)
	b := make([]byte, Buffer_sz)
	key := Keygen(Key_sz)
	for {
		n, err = fr.Read(b)
		//TODO remove fw if it is incomplete
		if err != nil && err != io.EOF {
			return 0, err
		}
		if (n == 0 || err == io.EOF) {
			if err = Keyf(opath + ".en.key", key); err != nil {
				return 0, err
			}
			return sum, nil
		}
		sum += n
		cbuf, err := Encypt(b[:n], key)
		if err != nil {
			return 0, err
		}
		fw.Write(cbuf)
	} 
}

func Keyf(_path string, key []byte) error {
	fw, err := os.Create(_path)
	if err != nil {
		return err
	}
	defer fw.Close()

	fw.Write(key[:Key_sz])
	return nil
}

func Decyptf(_path, opath string) (int, error) {
	if ext := path.Ext(_path); ext != ".en" {
		return 0, ErrExt
	} 
	
	fr, err := os.Open(_path)
	if err != nil {
		return 0, err
	}
	defer fr.Close()

	fk, err := os.Open(_path + ".key")
	if err != nil {
		return 0, err
	}
	defer fk.Close()
	
	k := make([]byte, Buffer_sz)
	key := make([]byte, Key_sz)
	m, err := fk.Read(k)
	if err != nil && err != io.EOF {
		return 0, err
	}
	if m != Key_sz {
		return 0, ErrKey 
	}
	key = k[:len(key)]

	_, fn := path.Split(_path)
	opath = path.Join(opath,fn)
	fw, err := os.Create(opath + ".de")
	if err != nil {
		return 0, err
	}
	defer fw.Close()

	var (
		n int 
		sum int
	)
	b := make([]byte, Buffer_sz)
	for {
		n, err = fr.Read(b)
		if err != nil && err !=io.EOF {
			return 0, err
		}
		if n == 0 || err == io.EOF {
			return sum, nil
		}
		sum += n
		dcbuf, err := Decypt(b[:n], key)
		if err != nil {
			return 0, err
		}
		fw.Write(dcbuf)  
	}
}

func Encypt(buf, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBEncrypter(block)

	if len(buf) % mode.BlockSize() > 0 {
		padder := padding.NewPkcs7Padding(mode.BlockSize())
		buf, err = padder.Pad(buf) 
		if err != nil {
			return nil, err
		}
	}
	cbuf := make([]byte, len(buf))
	mode.CryptBlocks(cbuf, buf)
	return cbuf, nil
}

func Decypt(cbuf, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBDecrypter(block)
	buf := make([]byte, len(cbuf))
	mode.CryptBlocks(buf, cbuf)
	if len(cbuf) < Buffer_sz && isPad(buf) {
		padder := padding.NewPkcs7Padding(mode.BlockSize())
		buf, err = padder.Unpad(buf) 
		if err != nil {
			return nil, err
		}
	}
	return buf, nil	
}

func isPad(buf []byte) bool {
	pad := buf[len(buf)-1]
	if int(pad) > Key_sz {
		return false
	}
	for _, v := range buf[len(buf)-int(pad) : len(buf)-1] {
		if v != pad {
			return false
		}
	}
	return true
}
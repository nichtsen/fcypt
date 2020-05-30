package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"github.com/nichtsen/fcypt/ecbf"
)

const dName = "output"
var (
	flagEncrypt = flag.Bool("en", false, "Encrypt files")
	flagDecrypt = flag.Bool("de", false, "Decrypt files")
	flagDir = flag.String("d", "", "Directory, \".\" for current directory")
	flagPath = flag.String("fp", "", "A certain file path to handle")
)

type Files struct {
	root	string 
	files	[]string
}

type crypto struct {
	cfiles		Files
	isEncypt	bool
	isBatch		bool
}

func main() {
	flag.Parse()
	count := 0
	flag.Visit(func(*flag.Flag){
		count += 1
	})
	if count == 0 {
		flag.Usage()
		os.Exit(1)		
	}

	if (*flagDecrypt && *flagEncrypt) || (*flagDir != "" && *flagPath != "") {
		log.Fatal("Invalid flag!")
	}
	
	if *flagDir != "" { 
		if fi, err := os.Stat(*flagDir); os.IsNotExist(err) || !fi.IsDir() {
			log.Fatal("Invalid Directory!")
		}
	}
	
	os.Mkdir(dName, 01)

	c := &crypto {
		Files{"", make([]string, 0)},
		*flagEncrypt, 
		false,
	}
	if *flagPath == "" {
		c.isBatch = true
		c.Batch(*flagDir) //root directory
	} else {
		c.Batch(*flagPath)
	}
}

func (f *Files) collect() error {
	err := filepath.Walk(f.root, func(_path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if _path == "." || _path == ".." || strings.Trim(_path, " ") == "" || info.IsDir()  {
			return nil
		}
		_path = strings.Replace(_path, "\\", "/", -1)
		f.files = append(f.files, _path)
		return nil 
	})
	if err != nil {
		return err
	}
	return nil
}

func (f *Files) String() {
	for i, _ := range f.files {
		log.Printf("%d. %s", i, f.files[i])
	}
}

func logs(mode, path string, n int){
	log.Printf("%s: %d bytes have been written to path ./%s ", mode, n, path)
}

func (c *crypto) Crypt(path string) (int, string, error) {
	if c.isEncypt {
		mode := "Encryption"
		n, err := ecbf.Encyptf(path, dName)
		return n, mode, err
	} else {
		mode := "Decryption"
		n, err := ecbf.Decyptf(path, dName)
		return n, mode, err
	}
}

func (c *crypto) Batch(_path string) {
	if c.isBatch {
		c.cfiles.root = _path
		if err := c.cfiles.collect(); err !=nil {
			log.Fatal(err)
		}
		c.cfiles.String()
		for i := range c.cfiles.files {
			n, mode, err := c.Crypt(c.cfiles.files[i])
			if err != nil {
				log.Fatal(err)
			}
			logs(mode, path.Join(dName,c.cfiles.files[i]), n)
		} 
	} else {
		n, mode, err := c.Crypt(_path)
		if err!= nil {
			log.Fatal(err)
		}
		logs(mode, path.Join(dName, _path), n)
	}
}
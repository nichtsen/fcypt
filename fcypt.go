package main

import (
	"flag"
	"log"
	"os"
	"github.com/nichtsen/fcypt/ecbf"
)

const dName = "output"
var (
	flagEncrypt = flag.Bool("en", false, "Encrypt files")
	flagDecrypt = flag.Bool("de", false, "Decrypt files")
	flagDir = flag.String("d", ".", "Directory, current as default" )
	flagPath = flag.String("fp", "", "A certain file path to handle")
)

func main() {
	flag.Parse()

	if !*flagDecrypt && !*flagEncrypt {
		flag.Usage()
	} 

	if *flagDecrypt && *flagEncrypt {
		log.Fatal("Invalid flag!")
	} 
	
	os.Mkdir(dName, 01)
	if *flagEncrypt {
		if *flagPath != "" {
			n, err := ecbf.Encyptf(*flagPath, dName)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Encryption: %d bytes have been written to path ./%s ", n, dName)
		}
	} else if *flagDecrypt {
		if *flagPath != "" {
			n, err := ecbf.Decyptf(*flagPath, dName)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Decryption: %d bytes have been written to path ./%s ", n, dName)
		}
	}


}
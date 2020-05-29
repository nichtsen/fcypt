# fcypt
> Quick file encryption/decryption via AES(ECB) 

## Installation
```
git get github.com/nichtsen/fcypt.git
cd $GOPATH/src/github.com/nichtsen/fcypt
go build 
```

## filename extension:

* __.key:__ 128-bits key,which should be kept separately with your encrypted file
* __.en:__ encrypted file
* __.de:__ decrypted file that the same as your original file after the extension being stripped


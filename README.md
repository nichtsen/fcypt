# fcypt
> Quick file encryption/decryption via AES(ECB) 

## Installation
```
go get github.com/nichtsen/fcypt
cd $GOPATH/src/github.com/nichtsen/fcypt
go build 
```

## filename extension:

* __.key:__ 128-bits key,which should be kept separately from your encrypted file
* __.en:__ encrypted file
* __.de:__ decrypted file that is the same as your original file after the extension being stripped

## language
[python implementation](https://github.com/nichtsen/symk-fcrypto)


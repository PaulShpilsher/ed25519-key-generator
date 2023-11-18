# public-key Ed25519 generator written in Go

**A simple public/private key generator using Edwards-curve Digital Signature Algorithm (EdDSA)**

## Usage

```sh
$ ed25519-key-generator [-out=filename]
```

Output is the pair of private and public key files:  ```filename.prv``` and ```filename.pub``` respectively.\
If command line option *-out* omitted the default filename is *"key.ed25519"*


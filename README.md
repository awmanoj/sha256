# sha256

## Build

```
$ git clone https://github.com/awmanoj/sha256.git 
$ go build
```

## Run 

```
$ ./sha256 <message> <salt>
```

## Examples

```
parallels$ ./sha256 "hello world" "ironman" # with salt 
SHA256('hello worldironman') = 8df09365de1079d79d9c3265397cf8a649bf8a65408ccc4e247839a8cbc11841
DoubleSHA256('hello worldironman') = 8142d966bd89e740d45b12de4fc592b10cb9cf8f220329ffddc848477fc90e0b
parallels$ 
parallels$ ./sha256 "hello world" ""  # no salt
SHA256('hello world') = b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
DoubleSHA256('hello world') = 049da052634feb56ce6ec0bc648c672011edff1cb272b53113bbc90a8f00249c

```
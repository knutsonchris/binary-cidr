# binary-cidr
Get the binary representation of an IP address with CIDR notation

`$ go get github.com/knutsonchris/binary-cidr`

`$ binary-cidr 192.168.2.2/16`
```
Binary representation
11000000.10101000.00000010.00000010

Subnet mask
11111111.11111111.00000000.00000000

Network ID
11000000.10101000.00000000.00000000

Broadcast
11000000.10101000.11111111.11111111

```

The network bits and host bits will be colorized in your terminal :)



## MAC
Get the local mac address.



## Install
```bash
$ go get -u github.com/xjh22222228/mac
```


## Usage
```go
// Fetch the computer's MAC address
macAddr, err := mac.Default()

// Fetch the computer's MAC address for a specific interface
macAddr, err := mac.AddrWithIface("en0")

// Validate that an address is a MAC address
isMac := mac.IsIsMAC("ac:de:48:00:11:22")
```


## LICENSE
[MIT](./LICENSE)



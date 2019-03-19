Ginga-Client-Go
===============
分布式锁 go版client

![image](ginga.jpeg)

Installation
------------

Install Ginga-Client-Go using the "go get" command:

    go get github.com/FastSchnell/Ginga-Client-Go/ginga


Usage
-----
```go
import "Ginga-Client-Go/ginga"

c := ginga.Client{
    Token: "test_token",
    Endpoint: "0.0.0.0:1903",
    Nonce: "test_nonce",
}

err := c.Lock()
if err != nil {
    fmt.Println(err.Error())
}

defer c.Unlock()

```

Server
------------
[Ginga Server](https://github.com/FastSchnell/Ginga)

License
-------

Ginga-Client-Go is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
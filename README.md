# Googletrans

This is Golang version of [py-googletrans](https://github.com/ssut/py-googletrans).

Googletrans is a **free** and **unlimited** Golang library that implemented Google Translate API.
This uses the [Google Translate Ajax API from Chrome extensions](https://chrome.google.com/webstore/detail/google-translate/aapbdbdomjkkjkaonfhkkikfgjllcleb) to make calls to such methods as detect and translate.

## Download from Github
```shell script
GO111MODULE=on go get github.com/dreamsxin/go-googletrans
```

## Quick Start Example

### Simple translate
```go
package main

import (
	"fmt"
	"github.com/dreamsxin/go-googletrans"
)

func main() {
	t := translator.New()
	result, err := t.Translate("你好，世界！", "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

### Using proxy
```go
c := translator.Config{
    Proxy: "http://PROXY_HOST:PROXY_PORT",
}
t := translate.New(c)
```

### Using custom service urls or user agent
```go
c := translator.Config{
    UserAgent: []string{"Custom Agent"},
    ServiceUrls: []string{"translate.google.com.hk"},
}
t := translate.New(c)
```

See [Examples](./examples) for more examples.

## Special thanks

* [py-googletrans](https://github.com/ssut/py-googletrans)

## License
This SDK is distributed under the [The MIT License](https://opensource.org/licenses/MIT), see [LICENSE](./LICENSE) for more information.

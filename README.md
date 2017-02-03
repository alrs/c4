
# C4 - The Cinema Content Creation Cloud
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![GoDoc](https://godoc.org/github.com/etcenter/c4go?status.svg)](https://godoc.org/github.com/etcenter/c4)
[![Stories in Ready](https://badge.waffle.io/etcenter/c4.png?label=ready&title=Ready)](https://waffle.io/etcenter/c4)
[![Build Status](https://travis-ci.org/etcenter/c4.svg?branch=master)](https://travis-ci.org/etcenter/c4)
[![Coverage Status](https://coveralls.io/repos/github/etcenter/c4/badge.svg?branch=master)](https://coveralls.io/github/etcenter/c4?branch=master)

## News
See the **Important Notification** for information on breaking changes in this release if you have C4 IDs stored that were generated prior to **February 7, 2017**

## C4 - Go
This package implements the C4 framework. See [Framework.md](https://www.github.com/etcenter/c4/Framework.md) for more details.

# Videos and Links
  - [C4 Framework Universal Asset ID](https://youtu.be/ZHQY0WYmGYU)
  - [The Magic of C4](https://youtu.be/vzh0JzKhY4o)

- [C4 ID Whitepaper](http://www.cccc.io/downloads/c4id_latest.pdf)

- Web: http://www.cccc.io/
- Mailing list: https://groups.google.com/forum/#!forum/c4-framework
- Twitter: https://twitter.com/CinemaC4

## Example
The following shows how to generate a C4 ID.

```go
package main

import (
  "fmt"
  "io"
  "os"

  // import 'asset' asset identification
  "github.com/etcenter/c4/asset"
)

func main() {
  file := "main.go"
  f, err := os.Open(file)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  // create a ID encoder.
  e := asset.NewIDEncoder()
  // the encoder is an io.Writer
  _, err = io.Copy(e, f)
  if err != nil {
    panic(err)
  }
  // ID will return an *asset.ID.
  // Be sure to be done writing bytes before calling ID()
  id := e.ID()
  // use the *asset.ID String method to get the c4id string
  fmt.Printf("C4id of \"%s\": %s\n", file, id.String())
  return
}

```


Output:

```bash
>go run main.go 
C4id of "main.go": c44jVTEz8y7wCiJcXvsX66BHhZEUdmtf7TNcZPy1jdM6S14qqrzsiLyoZRSvRGcAMLnKn4zVBvAFimNg14NFKp46cC
```


### Command line tool
See [c4 command line tool](https://github.com/etcenter/c4/tree/master/cmd/c4)

C4 is the Cinema Content Creation Cloud.  This repo contains the c4 command line interface,
and the c4 demon.  We are in the process of rolling out the following features:

- [x] Identify any file or block of data.
- [x] Identify folders and arbitrarily complex filesystems.
- [x] Key/c4 id store
- [ ] Threaded multi-target copy and id
- [ ] Optimized remote file sync ("the rsync killer").
- [ ] Dependency graph/workflow language
- [ ] Fuse based file system.
- [ ] PKI Security Model

---

## Important Notification
**The following only applies to those who've generated and stored C4 IDs prior to this notification.**

If you have been using C4 IDs prior to  be aware there has been a breaking change as part of the standardization process.

The base 58 character set has been changed.

Old: "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
New: "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

This change was necessary to insure identical sorting order for both the C4 ID string and the underlying hash data. Consistent sorting order simplifies many things, and is essential for correctly computing the C4 ID of a collections. The new character set is the same as the base 58 character set for bitcoin addresses.

This means that any C4 ID generated prior to the February 7th 2017 release will need to be regenerated or transformed to the new character set.  

Regenerating the IDs is the preferred approach.  However, this release includes functions to make it easier to directly transform the characters of exiting IDs into the new character set.  Care must be taken to perform the transformation from old character set to new only once. 


### License
This software is released under the MIT license.  See [LICENSE](./LICENSE) for more information.






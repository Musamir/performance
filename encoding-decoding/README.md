[![Test statuses: ](https://github.com/Musamir/performance/workflows/Test%20statuses/badge.svg??branch=master)](https://github.com/Musamir/performance/actions)

## Contents
- [Configurations](#Configurations)
- [Benchmarks](#Benchmarks)
  - [Encoding and Decoding](#Encoding and Decoding)
    - [Encoding](#Encoding)
    - [Decoding](#Decoding)
    - [Encoding&Decoding](#Encoding&Decoding)
  - [Marshaling and Unmarshalling](#Marshaling and Unmarshalling)
    - [Marshaling](#Marshaling)
    - [Unmarshalling](#Unmarshalling)
    - [Marshaling&Unmarshalling](#Marshaling&Unmarshalling)
  - [Quick start](#quick-start)

### Configurations

- **go version** : go 1.17.1 <br>
- **goos**    : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;windows<br>
- **goarch**  : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;amd64<br>
- **pkg**     : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;performance/encoding-decoding<br>
- **cpu**     : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz<br>

### Benchmarks


#### Encoding and Decoding

- ##### Encoding

|       lib      | operations | ns/op | B/op | allocs/op |
|:--------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json  |    1494144 | 826.3 |  208 |         3 |
| encoding/gob   |     278217 |  4983 | 1408 |        26 |

- ##### Decoding
|       lib      | operations | ns/op | B/op | allocs/op |
|:--------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json  |     347662 |  3428 | 1008 |        16 |
| encoding/gob   |      45146 | 25728 | 7568 |       209 |

- ##### Encoding&Decoding

|       lib      | operations | ns/op | B/op | allocs/op |
|:--------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json  |     288964 |  4306 | 1168 |        18 |
| encoding/gob   |      41682 | 30389 | 8936 |       234 |

<br>

#### Marshaling and Unmarshalling
- ##### Marshaling

|       lib       | operations | ns/op | B/op | allocs/op |
|:---------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json   |    1795035 | 705.6 |   64 |         1 |
| mailru/easyjson |    2573928 | 510.9 |  240 |         2 |
| protobuf/proto  |    4131890 | 288.2 |   48 |         1 |

- ##### Unmarshalling

|       lib       | operations | ns/op | B/op | allocs/op |
|:---------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json   |     428583 |  3000 |  248 |        12 |
| mailru/easyjson |    1204302 | 841.8 |   12 |         6 |
| protobuf/proto  |    1395759 | 935.4 |  312 |        13 |

- ##### Marshaling&Unmarshalling

|       lib       | operations | ns/op | B/op | allocs/op |
|:---------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json   |     342834 |  3757 |  312 |        13 |
| mailru/easyjson |     996850 |  1292 |  252 |         8 |
| protobuf/proto  |    1000000 |  1497 |  360 |        14 |

- #### quick-start
1. In case you want to test on your machine you need the go compiler version 1.13+ (
   if you don’t have, you can download the compiler on the official website [Go](https://golang.org/) and set your environment).
2. Open your command line in the program folder and use the below Go command to run benchmarks.

```sh
$ go test --bench=. --benchmem ./...
```

If you find any bugs, I would appreciate it if you let me know. Also, if you have any questions, I would love to answer
them. My contacts for communication: (mail - **mirovmusamir@gmail.com**), (telegram - **@MusamirMirov**)
[![Test statuses: ](https://github.com/Musamir/performance/workflows/Test%20statuses/badge.svg??branch=master)](https://github.com/Musamir/performance/actions)

**goos**    : &nbsp;&nbsp;&nbsp;windows<br>
**goarch**  : amd64<br>
**pkg**     : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;performance/encoding-decoding<br>
**cpu**     : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz<br>

##### Encoding

|       lib      | operations | ns/op | B/op | allocs/op |
|:--------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json  |    1525567 | 791.8 |  208 |         3 |
| encoding/gob   |     279340 |  4396 | 1408 |        26 |
| protobuf/proto |    3985990 | 305.9 |   48 |         1 |


##### Decoding

|       lib      | operations | ns/op | B/op | allocs/op |
|:--------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json  |     383084 |  3450 | 1008 |        16 |
| encoding/gob   |      49855 | 23928 | 7568 |       209 |
| protobuf/proto |    1409575 | 885.5 |  312 |        13 |


##### Encoding&Decoding

|       lib      | operations | ns/op | B/op | allocs/op |
|:--------------:|:----------:|:-----:|:----:|:---------:|
| encoding/json  |     235290 |  5045 | 1544 |        27 |
| encoding/gob   |      41368 | 29258 | 9120 |       237 |
| protobuf/proto |     923396 |  1337 |  472 |        15 |


If you find any bugs, I would appreciate it if you let me know. Also, if you have any questions, I would love to answer them.
My contacts for communication: (mail - **mirovmusamir@gmail.com**), (telegram - **@MusamirMirov**)
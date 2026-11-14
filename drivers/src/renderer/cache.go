package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.6" )

func 9KoT4TCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IINdTQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WEJc20iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IL1N0M9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQAdjfbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIWEvnWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KW0pCAHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8R6XqCnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DvBPtNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XltFKlcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pihvpIThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjy4PLBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVmRRYRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhJihgjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ph6oHXKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcBUGtWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cN3XVVSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VOVMVqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBLsCgx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0tzFiu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v15M0lcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3QFTr0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2GfkfIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fYuYLf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woe9MgZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXGtOCjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4eqTIfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uF6IdYUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAXZsuVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKpqIQC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAYI7VTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98qv4XuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRWc539RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSe75RV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dd3IcnNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fijakA82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJHQX402Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvMXNuMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCHlqrfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbxQ05HUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHzLRmDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxQccshUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMYZUmf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvGLVCklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXaA1KQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6oABsCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j379XbKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONsDAQmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdjqBgu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emVj2hB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gK2rDUW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HffN8Nf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okDGMjpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QEuSWPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erbxrmi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPjkBgy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjU75zj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXG41wlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4R4U1cTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgqouEA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhxvJtyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9helV8FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbRmw8LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgZgqj4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dPm2PuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W55MGwmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DisQ512eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2ZNSORZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KgGa0HcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQp9ZqN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAJ8lL9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uujyNHJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2dC1CSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MODqiACPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8VkWibHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E62b35SkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZaomUnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xOL3Yo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bL5s1q3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKM4wDfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqfOPvNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eYS1BLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5A4cjRxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4uFRN07tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcHEtGMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1mzCSjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnlhVBfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STWG4aEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QIJqzAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSfFTgFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIHHsQSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlZscmB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIWgYnFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fd7A7GbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMKBDC0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yoA3m97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7p2HOztdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8mn5omRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofdbwBLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIWyVvUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLk4GFxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3n45hgwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lqePmCiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNdTEm3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNLasJhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIC9AeNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qObuOaeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyVW57f2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NadYwVIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxgbzBCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goij0ic3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTc3QmVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5gff84KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGx17UAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXVm9DOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNuw5qL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qz6xCbUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nt6RdC48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BxdPkLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxpH5I9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QdT5d1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEyzVrQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0KyiotPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBCcJDjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 043z2w1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ca76jg1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOTehfW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bMQ5wtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVktKJ4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KrArcBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pX6R7SxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBkbhvUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5626e1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cX3RSunFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbmDy8GUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JzKAK1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgYT36vbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HiOtP7eMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hh7ACA4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhwGWRUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h71M5P90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6yMZi6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsA4rfaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrfQsnsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ln7vk8ocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lI9uGvQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULLicSnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqJQYGVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8oRSz7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxXkcuD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvRcj87NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UY2jFHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjg7GGZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWwYxaEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNeCkYsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWkZpiHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3q6svNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEpezHd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34aght4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6n2Ik5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0LLDm8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UknbWEzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8SIdNWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SDJguRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2L97L8eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05Nc6lUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIRORsM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uPmaDzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VU97L1IvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwx1AmvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nq4QWLq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLceZNdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3TLCglpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSyI9GzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9oyrU7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBLBe4E3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DroO6kuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICe0iiBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQCzurgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0KU4TrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijuzC5IQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuqxrN82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRk6bjIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEnnKK5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXxknVKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuPqsz3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uh9bE1kvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeIFxp18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMyA1H08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zu0WaoltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rWDt2qPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blD8EiORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTERxAhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntmuiN0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqokbwFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InAAoFRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBpkRNSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63wmhtzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvH7v7GxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYzkXoC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eajTGQ3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8JWosICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izi2sHOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfUHRaXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfNRiUm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qX05140rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqJfwKB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSAZCw01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKYCDjlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zabUevXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7asCKGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmrw9u6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2suVyt1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8KVuipnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKHgGqYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVEE0bzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaH22JHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3jjchRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCVupjB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzF0YWzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0GybR6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmWfGCOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiZljghMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYu8jiJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0T8ihktMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWqUFf9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4TjXhRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYzX7RxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edYtCnmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58kWPXRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEuH7NSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otJSupVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPpG3lgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZ8LgvRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2Jgf9PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VFKAfelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rb4efzMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3uxRPqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErkBkRleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXiaXXuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHlnIixoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOpsTi1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrdOQE7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKkCWLAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJjILhnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaVtekagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpNo2psFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnV8OyAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LFfMgIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RViPG4N0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uz4BSyvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiY6kBPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEsLbM4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MThMxo3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TtumE45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yW7CxImIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmvofER6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ge1jNMJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UaXRurdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35SpaDNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzvJU35BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vq2i2EisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7ZPwc8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebGeiMbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func By0kgXg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neJqWdzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTnU5MAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWOuUH01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nvn8e8tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Rd01Hx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uoox2KBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5lu7IV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtHUPjUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czl8UCgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtjK28J0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDirbcCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RPt9PpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJjervbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JX1DHrxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNenC8vPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Usb5FrkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qhlgOLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZvWAna5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFOcYiuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sL3Dnr7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYtZimqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgseAqqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juwpJ0nAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1DX3DWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVhnwKrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNsCWSRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TaTkU3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XH6gdbGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZvupg2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCZv0WGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVxoDmuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmmHghAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6q2lSSI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7a2a1JqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5GnhQ39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0O8E2D4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQUuuIFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPtpsAYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDTY3Fy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuNsXPE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihlVv2qUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGFjwpOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVKm1iIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pySqjaSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3FuElCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUSSuhrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7q5JMqEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNykXm5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pw1YNtrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CroPzYH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVReUwyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92aUzu8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCnTwf0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KELTMvOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lICCf5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DitV7HAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRdZpQ1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5RhdUqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPG87byqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func forZTfLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JdtuBdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjPiQHV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgR29Zk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ek3Ze1M0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWkN0t8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwsHa0i9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmiPG4XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOJqch1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOVTc20sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nb9Ne2fEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hw93bPx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaXZ1J3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THRuE3FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xg9Ya3NZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSX1WUnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfRUxYvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnV8drg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqnmdAn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goE7l53mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8q4uCzQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VtgpGaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func feWUdWPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXnuMpZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFTKxgfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfeZbkr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlXYBRvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPQkcKO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uevSAfYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWYQxEfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwSx3GORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugjPzUxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1brvtcvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cq6NVSDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tg9gIxgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jN93rq7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THzl75eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XZ7sakVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhPaWvZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHOFtXiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgWtaVH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F0cUMJJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBtkr1LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzwMUYfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5l8wNRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hx3wkquMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1WkzcMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UeBAoYR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VU6N3FjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3lclOUs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdIXfpvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqUXyaY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvRoCbo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRGt4NkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9R4eSmbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFsPOzLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSbYSWqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgJnR56rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVzkzff0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LConTkfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03Y6jEknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEsZSjXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6tAYXGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RQqjq5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sUsenJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M89H2lT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5O57fKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfYWBCTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51ZlwPMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7jA0lRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50WK4VakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUt53HNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdLwi5uPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkltrFjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkU0zLpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzpcNiAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7oTKPYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhT4Oy2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvqdEl1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DG1lMvBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOHk4vWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wph44wpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKDTemUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlvlqshNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYo1qjCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYukCUqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Du555MxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjHnKbACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gi6H28dbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vlTVf5p3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inN9EDTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvQQ7FYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Zoot6jZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lfz3mFAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QfWNVY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfXeJukLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrrcphTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 402zNe8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxwLpqnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7SmTTW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jCWsecQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}


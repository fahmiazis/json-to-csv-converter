package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.3" )

func JwglxHSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWKzP9AhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjyuzNlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0hyZkvbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0i6mKxPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9M7sAnptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNFDHqKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbJTIZfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvrMtckuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38FFtLq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPkqH3GVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JibSMvqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAalGW7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pjNXNTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNR2QTD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4A2liWsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PkwgGJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mu5FkvwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALyRw0qlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4LlwuGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rO8a4SMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XK9SjLr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrUw1LTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jr0eyDEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66o44un7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLxG07g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2m1aePRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5yqNVeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dzhz1IJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uj04qlzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuB6DWVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDAlGMfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEuDTdxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1kAZka1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFzsfomrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0flxp0LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WjSn54SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNfxv1hWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usMyKIvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6x4iuSqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8L4cJkDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPX2CVPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seEsw8d4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrzMtKrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4sicqdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs1exv3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YnMDVMjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HV94cwAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9cAifIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qegY4hlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeGHNqmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0gOnKMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7n2QuFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0nTk1ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cr0EOK7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VN4gYwTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRdBpVbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPt1Z1RYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4NXKzbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOvQqdSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkTGXQz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7lIfb21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVtKVzAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmsIRUQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5AsHggrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UUcFuhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEWm1V9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zblTksurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvrIdkG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNYdsu2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADLbavruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crTmDM9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5VSCVq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTTSfG3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPsTR2jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZ7BxClBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQo8VPOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NisAHy9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XS7PfYc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saTC0hFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dg0kie1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 547d8AerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGFCEZjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESSLv7UkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tg6lJB5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZemu0T7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ljx5I4e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZGj7GWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPfDyq2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2pnzGjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5VYD5EVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xF3InDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INgbZnzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLGmG6jpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWHUWANZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kJ4CT0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWipZzUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F4wqCbW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4E3o4hlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQL1x7aaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfbKF7OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBVXcweyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbo6MKcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1cfUGRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQ50YDnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbY3od60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBcS6WEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCAXOYxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4w1b6CRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUAsdsiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znmBurA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcB5vxLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBylLaIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcKgsHBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JF2Erv3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BMLeUItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aaCxAtcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zx8z6ao4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tZNvTOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgTa1VnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IMm8cGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhyxXfyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rpyXhd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgGdIX5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sA9giN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4zcPFpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PaCdBYF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlcYJE2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfxfAYzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjTlmgZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYlZn8JaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQMOMjIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0k2yxV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSiKeorlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaWCk5KeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqNbyMJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FInuaUsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tATjkYVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4g862PLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWq32koEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PM9KgMuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNggy86SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtXFNJgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IDOlcKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSvB58CnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsPTveLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgCSCbFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvcEopDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fF67CAzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDKu5Z2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVpSS1SbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuqxFTYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpDToEk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lj5c3TRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYyfQrMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1jwOpIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBxVWhDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7KlUBPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raD37FZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5zj8mDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYdnDbPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUQqb8W0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vt6oaPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPls1ODiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1S9INhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ti2D0tKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iW5vZqDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyQwb19hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmYYkbfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7os52OkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fk8SGPplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DutnlyyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQr8xQnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRqF4zPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZblBv4bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSi9cnOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaXhH8o9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ED6lmVE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ENdxBWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3YnLkp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhY7kvOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mt6E9X50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqwMMy6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXiv9Om7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrcATVIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHa555TxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BiDxePwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKCu0NdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kXlTrPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UI4zUmuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wer1ekqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRRMmLU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmJIo9qcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFnLdQiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F409p3JkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxULyo8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4jpiLvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNex0wWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eyfRjXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sau9cpkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bu9KNoWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4wVgKkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3y1hdv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qRNVjTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ub5jBIyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3atdjAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjBT2laIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nF1jND9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiclKNGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRFtz7DrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBlu6C5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ae3QmWBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcVTjgjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XaAHVbnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLhaIUlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDsN1rZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXyftWGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NhH2PZKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wn5RlYvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7B17maZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vX2bnUISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltptQdMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZqN15xFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fc1rGsVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeAXYDeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jl3hL3XzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uP4vqG9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVtCPw1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAY1IUo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F77Yz5uIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hn8S2C3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvodTo25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZjZD1iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OII50udWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbza55tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyLe3PvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62ogSmifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2hY2R0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FPsC1C5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9I2EocvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YFdVx4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLs7k4yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoiCuYk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDl6NOp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peYNa62AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlhwOs64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbVAxPq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGOtOep8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLn2BEQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLAkrsKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzbwFUg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1CQatC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oT4j6XxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RtzgYD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASM0s2KjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtdyQdcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jENBOqc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTCMNaSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkRxER3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xS7OSxGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MbcyiGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2kJ6QIXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dp9ljsf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEcGBYV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfPWyaCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnKragY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brX2fxlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9V1BbxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xlQQ0T0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euXrKN6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4YQii2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGItSqdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1llNjz3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0MWRRSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxBRJtc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYnpmD0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0M3lAjyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6M43d9yWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VvnUCRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TBrNe2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdFU7NCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGVSDub1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9AmW2YsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQOxi4BRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYCKHyJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGLIVMARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTwO3VHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCVvdUYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOehwxmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nRewNDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZheVCxTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goSSfZlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UDEZUhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r41xIs2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yqob9s0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXkwH6XUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJj2h2DUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xr0jf7tJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSmrZvIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mssQf5cOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRRhNPDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ho8CGnwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9GjIpuvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEXgGxTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWG1vbKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5WwgWo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGxF1LqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7booVzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Xzd1eKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoPy6cCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1wJ1YrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yhPz5mdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQCMKtGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRLwNRwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbm3Yzt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DaWWfwP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZlnDmjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBUMEwx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5e53JXIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z94lRA0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNqxLkLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTm1OQjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWXReMflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OlRFrhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUYH5cHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ei1nDYmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8jgCqbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pnbUSxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5TXFLANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xU8oAVNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 055HdV0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMbkUkQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7B4D2YUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5N9Qfjc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSg7VB19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLUQpE9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBcsTjLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQT7Is02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soyTMyuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opRDqp4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7H8rNyn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ad5OTYP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ivrvb81LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCnxkFXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TjW4EAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONAddnrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjX0FkAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3xV5yxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rr8GnpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0AJQAZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2HbRgWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7AUAM5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpZ4NBdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYHSUaBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNaHhA8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSIAkhXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wX0mmzbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbtwktuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VncRPJzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrLqr7ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2YAZZ0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJv80Cs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOpo1klMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4QMPVuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2BkGxKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2w6T24aBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyljdET9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UayPMlCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOD60lcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkhFOl5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7SZl0yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08gpPjNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clOUvsYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tb6MtnLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haPOo9SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 09J6DC5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilblkfEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrwmxRdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZQpPzBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OU1FFvvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4ZLIwIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRgki2nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDLyyR3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oz9AVzNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLRscJ9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhs2wApHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXOgx6SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9u8onk1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ao6sNGi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l0gDgi8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYePG3M9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kR3uBD90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MGkBvsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIWv5YaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vy58shBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4aLqK4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXJIflgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqCAohLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxsO32iOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK45UDXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lojy5BKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XjVHd8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tuBLNJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rDZA7EYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Mg1ebYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hyRFzpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHsJ46EpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haxUIJJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqyAc49ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yWHjSCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7tNxt85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8TGXgaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGhU1bKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xv3ak7zBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msVnRgDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceDzSkmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6xS0IklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfT4rVKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5mw0magWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JfGTdCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZvr7KL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYG9TtBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kVEVBfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Rr19xT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9YTVfYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVTjnvjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZdCTwD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rK5WJSPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mE7TBTK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5hLGTWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oklfiaPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7onr2G6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRNzLRNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVtfog98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlYEB8I0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Yqug13QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gGRf0pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlCQjVclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6N4rvIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLCCW3unWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zl9lvooNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wofriiVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKkbahlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuLJQ5bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GrMLh6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNSmFFCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjC0yO0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wo8L7j4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IP6JNKTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlrnjzojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOBeom2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVNgnxoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hq9xKsDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pst9BEk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlBAPYPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8JSm42tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29XOOgQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbxUpK7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnvP8sWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laKL44OwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIlmNsR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1kO8386Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKWo1ZgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZdbab4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYohbcWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1cfryy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QAvjKA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkSnVqBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfPMlVhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5UtM1p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOPblD52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GQlRzaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aP6WiiTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSmotWsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auFTClGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1U5Wlka3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w54ZLgIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCsTmqriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJUlQ4ZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPFgQ7yqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRYEFmMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZbbo1BuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BALs1qQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c30YuxQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c73jvdvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8lxSTE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNT6ZBPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DT4LwEy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fn5YpkBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdohCLVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQrulZbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLDBFm4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJRa2cTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNbLl1xoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIIK1uvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXXn4ZOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRpFi3G2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twrmqVWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oikK8QRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzJdCmvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTw9nbViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Um7y64pwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMAUU2vHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1SfeCfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gubQgzLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyoYajs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjyxXZxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLxDvJPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebNCnXm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPZyZZkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOnsfaXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEdofTSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvRddEy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDzryindWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oic9bCyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26ymj9JAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHeeY1frWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0cyx0dPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fl0pscf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqO4GbF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyGXiUA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXcfF47jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txfmiv84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUgjgVdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YF9nOBdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Wf679rhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UTyg4mNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZOdRkozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbBhXqu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5WzVMpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAyDtsQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ee0SFPQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSw1qdW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQvbLLX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuP0gbvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJBe6pCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9j3gJXfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqL2xTufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EBy8rkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAcsbpDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnC76nD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpa80zaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1UD3IUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XI7Zh9zJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMp3vDKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoXQuBfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOwaYrQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWl4h6hXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sM63VxCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCnvPkh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzhsGVoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDepZ5cWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4jtiIqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oBzL9ODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAz8MkrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bs2GYcwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCQslrd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zTa7rRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMR808qRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dooka3lSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3WyJ69TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUTD9c5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYQqX5hRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1A8egQXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aiIe4GIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5HnYbUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8aeaD6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnBiKMjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uDGDdnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPHI0WANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ALMGrF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYkkr3dhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pt4e11OzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQTJePf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRenubsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kE2ZPAPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6WUbmcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwGjiFwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vm24ghwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEcbwVQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxxSHSSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZ1hsOEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8JfzsUtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1bX23QdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhcnqoTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BI1hlVBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLtlMe28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48NlcmWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnGAvVYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzuttlzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func np2LoH4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtXjoSlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSbcStdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8y2qom53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baNjb7tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwii5L9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLvgtwskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvvsQ9bEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0GVUMiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHN68TwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcCbKqVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEyKQXj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqd2wRb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3klGwRPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHphkhkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nNhqojrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gj93WnsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uNh6VzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlIZpIvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRmWni6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXI3jMN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBuAbSEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILWZzzIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0A1iqoH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyHN3AXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teeR02SSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lisaBjhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6168vZ0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRW7n4Z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9fEA0pEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtP9FASOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfMfVum9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLyrCL2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fM7w6jctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkl7WzoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUJfRfY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8LlnQDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVWZkm7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mevpOK88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b48oYflcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBJm9NHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdVJsoZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDcoMarOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9EqwLECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZjnSMZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nc9EwJs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STJnwBXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBZCU52SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmeENwtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htJh0NG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cos5LHDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkKQexqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuGoNVhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpyGddcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhzcNzNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jxh1Xi82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13LVsoxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2N3szcDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVCci2dXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEYyciI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gp9MUjJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQjFle6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajtLXNkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viL4PAbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hif87UIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NqOrpesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYWYgy9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOzjIBdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3iWjVpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iym2qEJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ae9mkVTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fT34u211Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZh15xihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SufvHTCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVlpHcszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQSIoUKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Gi3reG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtqogrmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAOYtEVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8vzpmVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWgkC3FRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Im4lgqYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmBtwGEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYdm2SSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqHAnawWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BeAyyd43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Znpb7Ep1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6GlvNY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNCgyYrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTGiFSQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9qPlCPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HYrahuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UC5PNPSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIZdwQCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6i45ogOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnFcO9uzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BU9Wr4R5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWMpUHi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQDEucEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQwsfnVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qu3SiR55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMKLjm3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLfVpgiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhI1nIGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjclXZMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ndFOfx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjySeH3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLQHUAJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oe0o03fxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNAwcL8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJiaJNmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eByUjbxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHMHsp5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhInNquTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDJ1sEdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ug8hcDUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5qEuwv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbWZcHPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWLWxmS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfQRE5gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHgyTYdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWq92yITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAEtJ8mSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeQkkH5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFNtrxCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZctJ6HGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNwBC5gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Hh4mzMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7S9zeM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmVfhXj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8lg7DWUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcAG2SPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beJGDe3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uf6pxjKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSGmkwZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Sn2OONGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EvZDQvdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IM0AGJp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPFruoFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBIZVIBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4o4Y4zKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOZUUHKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gm9kRzIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxWa7IwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyxD6S6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqXIXMdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOk0vPtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2nAgKfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYSfcl9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmueTkg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHdN0OSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7rRHCNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iq1qg8BuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ou4aB40eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdghs1vMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YGGBU2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RE7psQ1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLWhhcHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MmlNPNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vNprkbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiTnlNqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGewCFJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGxZXmyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXKqTLXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfQu8SoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgBTSABXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}


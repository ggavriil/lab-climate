package dht;

// #cgo CFLAGS: -I${SRCDIR}/../cdht
// #cgo LDFLAGS: -L${SRCDIR}/../lib -ldht -lgpio
// #include "pi_dht_read.h"
import "C"

import (
    "fmt"
    "time"
)

var isInit bool = false;

var DHT_RESPONSE = map[int]string {
    -1: "ERRR_TIMEOUT",
    -2: "ERROR_CHECKSUM",
    -3: "ERROR_ARGUMENT",
    -4: "ERROR_GPIO",
     0: "SUCCESS",
}

func DhtRead() (float32, float32, int) {
    var res C.temp_data_t = C.get_temp_data();
    for(int(res.return_code) != 0) {
        fmt.Println(int(res.return_code));
        time.Sleep(time.Second * 2);
        res = C.get_temp_data();
    }
    return float32(res.temperature), float32(res.humidity), int(res.return_code);
}

package gpio

// #cgo CFLAGS: -I../cgpio
// #cgo LDFLAGS: -L${SRCDIR}/../lib -lgpio
// #include "gpio.h"
import "C"

var isInit bool = false;

func InitGpio ()   {
    if(!isInit) {
        C.init_gpio();
    }
    isInit = true;
}

func InputPin(p int) {
    if(!isInit) {
        return;
    }
    C.input_pin(C.int(p));
}

func OutputPin(p int) {
    if(!isInit) {
        return;
    }
    C.output_pin(C.int(p));
}

func IsSet(p int) bool {
    if(!isInit) {
        return false;
    }
    return bool(C.is_set(C.int(p)));
}

func SetPin(p int) {
    if(!isInit) {
        return;
    }
    C.set_pin(C.int(p));
}

func ClearPin(p int) {
    if(!isInit) {
        return;
    }
    C.clear_pin(C.int(p));
}

func WritePin(p int, v bool) {
    if(!isInit) {
        return;
    }
    var vb C._Bool = false;
    if(v) {
        vb = true;
    }
    C.write_pin(C.int(p), vb);
}


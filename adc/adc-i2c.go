package adc

import (
    "encoding/binary"
    "golang.org/x/exp/io/i2c"
    "time"
    "log"
    "os"
)

const CONFIG_OS_SINGLE     uint16 = 0x8000;
const CONFIG_MODE_SINGLE   uint16 = 0x0100;
const CONFIG_MUX_OFFSET    uint16 = 12;

var CONFIG_GAIN = map[int]uint16 {
    0:  0x0000, //2/3
    1:  0x0200,
    2:  0x0400,
    4:  0x0600,
    8:  0x0800,
    16: 0x0A00,
};

var CONFIG_DATA_RATE = map[int]uint16 {
    8:   0x0000,
    16:  0x0020,
    32:  0x0040,
    64:  0x0060,
    128: 0x0080,
    250: 0x00A0,
    475: 0x00C0,
    860: 0x00E0,
}

const CONFIG_DATA_RATE_128 uint16 = 0x0080;
const CONFIG_COMP_DISABLE  uint16 = 0x0003
const POINTER_CONFIG       byte   = 0x01;
const POINTER_CONVERSION   byte   = 0x00;


type adci2c struct {
    isOpen bool;
    pin byte;
    dev *i2c.Device;
    logger *log.Logger;
}


func NewAdc(pin byte) *adci2c {
    adc := new(adci2c);
    adc.isOpen = false;
    adc.pin = pin;
    adc.dev = nil;
    adc.logger = log.New(os.Stderr, "I2C ", log.Ldate | log.Ltime);
    return adc;
}


func (adc *adci2c) GetVal() (int16, bool) {
    if(!adc.isOpen) {
        return 0, false;
    }
    config := makeAdcConfig(adc.pin);
    bytes := make([]byte, 2)
    binary.BigEndian.PutUint16(bytes, config);
    err := adc.dev.WriteReg(POINTER_CONFIG, bytes);
    if(err != nil) {
        adc.logger.Println("Couldn't write reg with err: " + err.Error());
        return 0, false;
    }
    time.Sleep(80 * time.Millisecond);
    result := make([]byte, 2);
    err = adc.dev.ReadReg(POINTER_CONVERSION, result);
    if(err != nil) {
        adc.logger.Println("Couldn't read reg with err: " + err.Error());
        return 0, false;
    }
    resultValue := binary.BigEndian.Uint16(result);
    resultIntValue := int16(resultValue);
    return resultIntValue, true;
}

func makeAdcConfig (pin byte) uint16 {
    var config uint16 = CONFIG_OS_SINGLE;
    var mux byte = 0x04;
    config |= (uint16((pin + mux) & 0x07) << CONFIG_MUX_OFFSET);
    config |= CONFIG_GAIN[1];
    config |= CONFIG_MODE_SINGLE;
    config |= CONFIG_DATA_RATE[128];
    config |= CONFIG_COMP_DISABLE;
    return config;
}


func (adc *adci2c) Open() bool {
    d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x48);
    if(err != nil) {
        adc.logger.Print("Couldn't open device with err: " + err.Error());
        return false;
    }
    adc.dev = d;
    adc.isOpen = true;
    return true;
}

func (adc *adci2c) Close() bool {
    err := adc.dev.Close();
    if(err != nil) {
        adc.logger.Println("Closing device failed with err: " + err.Error());
        return false;
    }
    return true;
}

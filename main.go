package main

import (
    "fmt"
    "github.com/ggavriil/lab-climate/adc"
    "github.com/ggavriil/lab-climate/gpio"
    "github.com/ggavriil/lab-climate/dht"
    "time"
)

func main() {
    fmt.Println("Hello!");
    adc:= adc.NewAdc(0);
    e:= adc.Open();
    if(!e) {
        fmt.Println("Error");
    }
    for {
        v, eq := adc.GetVal();
        t, h, et := dht.DhtRead();
        if(eq && et == 0) {
            fmt.Printf("Temperature: %f, Humidity: %f, Air Quality: %d\n",
                        t, h, v);
        }
        time.Sleep(time.Second * 2);
    }
    e = adc.Close();
    if(!e) {
        fmt.Println("Error");
    }
    fmt.Println("Hello");
    gpio.InitGpio();
    gpio.InputPin(17);
    gpio.OutputPin(17);
    for {
        gpio.SetPin(17);
        time.Sleep(time.Second);
        gpio.ClearPin(17);
        time.Sleep(time.Second);
    }
}

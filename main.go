package main

import (
    "fmt"
    "tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
    fmt.Println("BLE Scanner for Nice!Nano")
    
    // Включаем BLE адаптер
    err := adapter.Enable()
    if err != nil {
        fmt.Println("Error enabling BLE:", err)
        return
    }

    // Запускаем сканирование
    fmt.Println("Scanning for BLE devices...")
    err = adapter.Scan(scanHandler)
    if err != nil {
        fmt.Println("Error starting scan:", err)
        return
    }

    // Бесконечный цикл для поддержания работы
    for {
        select {}
    }
}

func scanHandler(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
    deviceInfo := fmt.Sprintf("Device found: [%s]", result.Address.String())
    
    if result.RSSI != 0 {
        deviceInfo += fmt.Sprintf(" RSSI: %d dBm", result.RSSI)
    }
    
    if name := result.LocalName(); name != "" {
        deviceInfo += fmt.Sprintf(" Name: %s", name)
    }
    
    fmt.Println(deviceInfo)
}
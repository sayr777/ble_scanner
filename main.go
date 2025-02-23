package main

import (
    "encoding/hex"
    "fmt"
    "strings"
    "sync"
    "tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

type DeviceInfo struct {
    Address     string
    RSSI        int16
    LocalName   string
    ManufacturerData []bluetooth.ManufacturerDataElement
    Payload     []byte
}

var (
    seenDevices = make(map[string]DeviceInfo)
    mu          sync.Mutex
)

func main() {
    fmt.Println("BLE Scanner for Nice!Nano")

    err := adapter.Enable()
    if err != nil {
        fmt.Println("Error enabling BLE:", err)
        return
    }

    fmt.Println("Scanning for BLE devices...")
    err = adapter.Scan(scanHandler)
    if err != nil {
        fmt.Println("Error starting scan:", err)
        return
    }

    select {}
}

func scanHandler(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
    mu.Lock()
    defer mu.Unlock()

    address := result.Address.String()
    if _, exists := seenDevices[address]; exists {
        return
    }

    device := DeviceInfo{
        Address:     address,
        RSSI:        result.RSSI,
        LocalName:   result.LocalName(),
        ManufacturerData: result.ManufacturerData(),
        Payload:     result.AdvertisementPayload.Bytes(),
    }

    seenDevices[address] = device
    printDeviceInfo(device)
}

func printDeviceInfo(device DeviceInfo) {
    var infoParts []string

    infoParts = append(infoParts, fmt.Sprintf("Address: %s", device.Address))

    if device.RSSI != 0 {
        infoParts = append(infoParts, fmt.Sprintf("RSSI: %d dBm", device.RSSI))
    }

    if device.LocalName != "" {
        infoParts = append(infoParts, fmt.Sprintf("Name: %s", device.LocalName))
    }

    if len(device.ManufacturerData) > 0 {
        for _, mfg := range device.ManufacturerData {
            infoParts = append(infoParts, fmt.Sprintf("Mfg[%04X]: %X", 
                mfg.CompanyID, mfg.Data))
        }
    }

    if len(device.Payload) > 0 {
        infoParts = append(infoParts, 
            fmt.Sprintf("Payload: %s", hex.EncodeToString(device.Payload)))
    }

    fmt.Println("Device found:", strings.Join(infoParts, ", "))
}
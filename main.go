package main

import (
    "machine"
    "time"
    "tinygo.org/x/drivers/ble"
)

func main() {
    // Инициализация UART для вывода логов
    uart := machine.UART1
    uart.Configure(machine.UARTConfig{
        BaudRate: 115200,
    })

    // Инициализация BLE
    bleRadio := ble.NewBLE()
    bleRadio.Configure()

    // Вывод сообщения о начале сканирования
    println("Starting BLE scan...")

    // Начало сканирования
    bleRadio.Scan(func(addr ble.Address, rssi int8, advData []byte) {
        // Вывод информации о найденном устройстве
        println("Device found:")
        println("  Address:", addr.String())
        println("  RSSI:", rssi)
        println("  Advertisement data:", string(advData))
    })

    // Запуск бесконечного цикла для продолжения работы программы
    for {
        time.Sleep(time.Second)
    }
}

// Функция для вывода текста через UART
func println(text ...interface{}) {
    output := machine.UART1
    for _, t := range text {
        output.WriteString(t.(string))
    }
    output.WriteString("\r\n")
}
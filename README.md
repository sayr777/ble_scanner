# BLE Device Scanner for nRF52840 using TinyGo

This program is designed to scan for Bluetooth Low Energy (BLE) devices using the `nRF52840` microcontroller on a `nice!nano` board. Below is an overview of the tasks, capabilities, and limitations of this program.


go get tinygo.org/x/bluetooth




---

## Tasks

1. **Initialize UART for Logging**:
   - Configures the UART interface (`machine.UART1`) to output log messages at a baud rate of 115200.
   - This allows debugging information to be sent to a serial terminal or another device connected via UART.

2. **Set Up BLE Radio**:
   - Initializes the BLE radio using the `tinygo.org/x/drivers/ble` package.
   - Configures the BLE module to prepare it for scanning.

3. **Start Scanning for BLE Devices**:
   - Begins scanning for nearby BLE devices using the `Scan` method.
   - For each detected device, the program logs the following details:
     - Device address (`addr.String()`).
     - Received Signal Strength Indicator (RSSI) value (`rssi`).
     - Advertisement data (`advData`), which may include device name, service UUIDs, or other metadata.

4. **Run Continuously**:
   - Enters an infinite loop to keep the program running and continuously scanning for new devices.

5. **Custom `println` Function**:
   - Implements a helper function to simplify writing text to the UART interface with automatic line breaks.

---

## Capabilities

- **Device Discovery**:
  - The program can detect BLE devices within range and display their addresses, signal strength, and advertisement data.
  
- **Real-Time Monitoring**:
  - It provides real-time feedback by logging discovered devices as they are detected.

- **Cross-Platform Compatibility**:
  - Written in Go with TinyGo, this program can be compiled and deployed on various microcontroller platforms that support TinyGo.

- **Low Resource Usage**:
  - Designed to run efficiently on resource-constrained devices like the nRF52840.

---

## Limitations

1. **No Zigbee Support**:
   - The program only supports BLE and does not natively support Zigbee. To work with Zigbee devices, additional hardware (e.g., a Zigbee-to-BLE bridge) or software integration would be required.

2. **Limited Data Parsing**:
   - The program does not parse the advertisement data (`advData`) into meaningful fields such as device names or service UUIDs. This requires additional logic to interpret the raw byte array.

3. **Infinite Loop**:
   - The program runs indefinitely, which may not be suitable for battery-powered devices without proper power management.

4. **Dependence on External Libraries**:
   - Relies on the `tinygo.org/x/drivers/ble` library, which must be compatible with the target hardware and firmware version.

5. **UART Output Only**:
   - Logs are sent exclusively via UART. If no serial connection is established, the output will not be visible.

6. **Scanning Range**:
   - The range of detection depends on the BLE antenna and environmental factors, which may limit its effectiveness in certain scenarios.

---

## Notes

- This program is intended for educational purposes or as a starting point for building more advanced BLE applications.
- To extend functionality, you could add features such as filtering specific device types, saving scanned data to storage, or integrating with other communication protocols.

If you need to integrate Zigbee functionality, consider using a dedicated Zigbee module or exploring libraries that provide Zigbee stack implementations.
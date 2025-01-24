
# **SSRF Tester**

The **SSRF Tester** is a lightweight and efficient tool designed to identify potential Server-Side Request Forgery (SSRF) vulnerabilities in web applications. It appends predefined SSRF payloads to specified URL parameters and checks for responses that indicate potential SSRF vulnerabilities.

---

## **Features**
- **Predefined Payloads**:
  - Includes common SSRF payloads targeting internal services like `localhost`, `127.0.0.1`, `169.254.169.254`, and `metadata.google.internal`.
- **Customizable Parameters**:
  - Allows you to specify the URL parameter to test for SSRF vulnerabilities.
- **Report Generation**:
  - Saves detected SSRF vulnerabilities and their status codes to a specified output file.
- **Concurrent Testing**:
  - Supports multithreaded testing for faster results.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/SSRFTester.git
   cd SSRFTester
   ```

2. **Build the Tool**:
   Compile the tool into a binary:
   ```bash
   go build -o ssrf-tester main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./ssrf-tester [flags]
   ```

---

## **Usage**
```bash
go run main.go [flags]
```

### **Available Flags**
| Flag             | Description                                        | Default               |
|-------------------|----------------------------------------------------|-----------------------|
| `-url`           | Base URL to test (e.g., `http://example.com`)      | *None* (required)     |
| `-param`         | Query parameter to append payloads                 | *None* (required)     |
| `-o`             | Path to save the results                           | `ssrf_results.txt`    |
| `-t`             | Number of concurrent threads                       | `10`                  |
| `-v`             | Enable verbose mode for detailed output            | Disabled              |

---

## **Examples**
### **Basic Usage**
```bash
go run main.go -url http://localhost:3000 -param url
```
Tests the `url` parameter of the given base URL using predefined SSRF payloads.

### **Save Results to a File**
```bash
go run main.go -url http://localhost:3000 -param url -o custom_results.txt
```
Saves the results to `custom_results.txt`.

### **Increase Threads for Faster Testing**
```bash
go run main.go -url http://localhost:3000 -param url -t 20
```
Runs the test with 20 concurrent threads for faster execution.

---

## **Sample Output**
The following is an example of the content saved in `ssrf_results.txt`:
```
[POTENTIAL SSRF] URL: http://localhost:3000?url=http%3A%2F%2Flocalhost%2F | Status Code: 200
[POTENTIAL SSRF] URL: http://localhost:3000?url=http%3A%2F%2F127.0.0.1%2F | Status Code: 200
[POTENTIAL SSRF] URL: http://localhost:3000?url=http%3A%2F%2F169.254.169.254%2F | Status Code: 200
[POTENTIAL SSRF] URL: http://localhost:3000?url=http%3A%2F%2Fmetadata.google.internal%2F | Status Code: 200
[POTENTIAL SSRF] URL: http://localhost:3000?url=http%3A%2F%2F%5B%3A%3A1%5D%2F | Status Code: 200
```

---

## **Planned Enhancements**
1. **Dynamic Payload Generation**:
   - Support for generating payloads based on target application behavior.
2. **Response Analysis**:
   - Advanced response inspection to detect SSRF beyond status codes.
3. **Integration with External Services**:
   - DNS logging or webhook payloads for out-of-band SSRF detection.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for more details.

---

## **Acknowledgments**
- Inspired by OWASP Juice Shop for testing.
- Built with Go to showcase development and penetration testing skills.

---

## **Performance Tips**
- Use higher thread counts (`-t`) for faster results.
- Save results to a file for easier review and integration into reports.
- Test against applications that handle external requests for better SSRF detection.

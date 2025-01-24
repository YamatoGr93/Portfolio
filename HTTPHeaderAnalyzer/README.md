
# **HTTP Header Analyzer**

The **HTTP Header Analyzer** is a powerful tool designed to assess the security posture of web applications by analyzing HTTP response headers. Built with Go, it helps penetration testers and bug bounty hunters quickly identify missing or misconfigured security headers and provides actionable recommendations.

---

## **Features**
- **Critical Header Analysis**:
  - Automatically checks for common security headers, including:
    - `Content-Security-Policy`
    - `Strict-Transport-Security`
    - `X-Content-Type-Options`
    - `X-Frame-Options`
    - `Referrer-Policy`
- **Detailed Recommendations**:
  - Provides remediation advice for missing or misconfigured headers.
- **Customizable Header Checks**:
  - Supports user-defined headers via a custom file.
- **Report Generation**:
  - Saves analysis results, including recommendations, to an output file.
- **Verbose Mode**:
  - Optionally enables detailed logs for debugging.
- **Lightweight and Fast**:
  - Built with Go for high performance and simplicity.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/HTTPHeaderAnalyzer.git
   cd HTTPHeaderAnalyzer
   ```

2. **Build the Tool**:
   Compile the tool to a binary:
   ```bash
   go build -o http-header-analyzer main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./http-header-analyzer <URL> [flags]
   ```

---

## **Usage**
```bash
go run main.go <URL> [flags]
```
Replace `<URL>` with the target URL (e.g., `http://localhost:3000/#/` for Juice Shop).

### **Available Flags**
| Flag             | Description                                             | Default               |
|-------------------|---------------------------------------------------------|-----------------------|
| `-o`             | Path to the output file                                 | `headers_report.txt`  |
| `-h`             | Path to a custom headers file                           | *None*                |
| `-v`             | Enable verbose mode for detailed logs                   | *Disabled*            |

---

## **Examples**
### **Basic Usage**
```bash
go run main.go http://localhost:3000/#/
```
Analyzes the default list of headers for the target URL.

### **Save Results to a File**
```bash
go run main.go http://localhost:3000/#/ -o my_report.txt
```
Saves the results to `my_report.txt`.

### **Check Custom Headers**
Create a file, e.g., `custom_headers.txt`:
```text
X-Permitted-Cross-Domain-Policies
Permissions-Policy
```

Run the tool with the custom headers file:
```bash
go run main.go http://localhost:3000/#/ -h custom_headers.txt
```

### **Enable Verbose Mode**
```bash
go run main.go http://localhost:3000/#/ -v
```
Displays detailed logs during header analysis.

---

## **Default Headers Checked**
The tool analyzes the following security headers by default:
1. **Content-Security-Policy**: Mitigates XSS by controlling resource loading.
2. **Strict-Transport-Security**: Enforces HTTPS to prevent protocol downgrade attacks.
3. **X-Content-Type-Options**: Prevents MIME type sniffing.
4. **X-Frame-Options**: Protects against clickjacking.
5. **Referrer-Policy**: Controls referrer information shared with other sites.

---

## **Performance Tips**
- Use the tool with `localhost` or hosted applications for testing.
- Save results to an output file for reference in reports or audits.
- Add custom headers relevant to your use case for a more tailored analysis.

---

## **Planned Enhancements**
Future updates may include:
- Validating header values (e.g., `max-age` for `Strict-Transport-Security`).
- Adding output formats like JSON or Markdown.
- Enhancing concurrency for analyzing multiple URLs in parallel.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for details.

---

## **Acknowledgments**
- Inspired by Juice Shop for testing.
- Built with Go to showcase both development and penetration testing skills.


# **Open Redirect Finder**

The **Open Redirect Finder** is a tool designed to identify open redirect vulnerabilities by appending predefined payloads to specific URL parameters. It is built to help penetration testers and bug bounty hunters detect this common security issue.

---

## **Features**
- **Predefined Redirect Parameters**:
  - Tests parameters like `redirect`, `url`, `next`, `target`, and others.
- **Payload Testing**:
  - Utilizes a set of predefined payloads such as `https://evil.com` and `http://localhost`.
- **Custom Wordlists**:
  - Allows testing custom redirect parameters from a user-defined wordlist.
- **Concurrent Testing**:
  - Supports multithreading for faster testing with a configurable number of threads.
- **Results Reporting**:
  - Saves discovered vulnerabilities to an output file.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/OpenRedirectFinder.git
   cd OpenRedirectFinder
   ```

2. **Build the Tool**:
   Compile the tool to a binary:
   ```bash
   go build -o open-redirect-finder main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./open-redirect-finder <URL> [flags]
   ```

---

## **Usage**
```bash
go run main.go <URL> [flags]
```
Replace `<URL>` with the target URL (e.g., `http://localhost:3000` for Juice Shop).

### **Available Flags**
| Flag             | Description                                   | Default             |
|-------------------|-----------------------------------------------|---------------------|
| `-o`             | Path to the output file                      | `redirects.txt`     |
| `-w`             | Path to a custom wordlist                    | *None*              |
| `-t`             | Number of concurrent threads                 | `10`                |
| `-v`             | Verbose mode to print detailed testing logs  | *Off*               |

---

## **Examples**
### **Basic Usage**
```bash
go run main.go http://localhost:3000
```

### **Save Results to a File**
```bash
go run main.go http://localhost:3000 -o results.txt
```

### **Use a Custom Wordlist**
```bash
go run main.go http://localhost:3000 -w wordlists/custom.txt
```

### **Increase Threads for Faster Testing**
```bash
go run main.go http://localhost:3000 -t 20
```

---

## **Limitations**
- **No Results Found**: 
  - The tool did not find any open redirect vulnerabilities in Juice Shop during testing. This is likely because Juice Shop does not have open redirect vulnerabilities by default.
  - This does not indicate a failure of the tool but rather the absence of such vulnerabilities in the tested application.
  
- **Context-Specific Testing**: 
  - The tool relies on common redirect parameters (`redirect`, `url`, etc.) and may miss custom or less common implementations of open redirects.
  - Payloads are predefined and may not cover all possible cases.

---

## **Planned Enhancements**
1. **Dynamic Payload Generation**:
   - Create payloads dynamically based on target application behavior.
2. **Custom Response Analysis**:
   - Add features to analyze response headers and bodies for additional clues.
3. **Integration with External APIs**:
   - Include DNS logging or webhook-based payloads to detect out-of-band redirects.
4. **Better Debugging and Reporting**:
   - Improve verbose output to provide more details on why a particular parameter failed or passed the test.

---

## **Future Goals**
- This tool was created to demonstrate Go programming skills and an understanding of open redirect vulnerabilities. While it works as intended, its effectiveness depends on the target application having such vulnerabilities.
- Future updates will aim to increase the tool's flexibility and expand its detection capabilities for real-world bug bounty and penetration testing scenarios.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for details.

---

## **Acknowledgments**
- Inspired by Juice Shop for testing and learning.
- Built with Go to showcase both development and penetration testing skills.
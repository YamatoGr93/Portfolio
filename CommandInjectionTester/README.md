
# **Command Injection Tester**

The **Command Injection Tester** is a lightweight tool built with Go to help penetration testers and bug bounty hunters identify potential command injection vulnerabilities in web applications. It works by injecting commonly used commands into specified parameters and observing the application's response.

---

## **Features**
- **Custom Wordlist** (`-w`): Use your own wordlist of commands for testing.
- **Save Results to File** (`-o`): Save discovered vulnerabilities to a file.
- **Multithreading** (`-t`): Perform concurrent requests for faster testing.
- **Verbose Mode** (`-v`): Enable detailed logs for debugging.
- **Default Configurations**: Includes a default wordlist (`wordlists/common.txt`) and output file (`results.txt`).

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/CommandInjectionTester.git
   cd CommandInjectionTester
   ```

2. **Build the Tool**:
   Compile the tool to a binary:
   ```bash
   go build -o command-injection-tester main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./command-injection-tester <URL> [flags]
   ```

---

## **Usage**
```bash
go run main.go <URL> [flags]
```
Replace `<URL>` with the target URL (e.g., `http://localhost:3000` for Juice Shop).

### **Available Flags**
| Flag             | Description                                   | Default               |
|-------------------|-----------------------------------------------|-----------------------|
| `-w`             | Path to the wordlist file                    | `wordlists/common.txt`|
| `-o`             | Path to the output file                      | `results.txt`         |
| `-t`             | Number of concurrent threads                 | `10`                  |
| `-v`             | Enable verbose mode for detailed logs        | *Off*                 |

---

## **Examples**
### **Basic Usage**
```bash
go run main.go http://localhost:3000
```
Scans the target URL using the default wordlist.

### **Save Results to a File**
```bash
go run main.go http://localhost:3000 -o vulnerabilities.txt
```
Saves the results to `vulnerabilities.txt`.

### **Specify a Custom Wordlist**
```bash
go run main.go http://localhost:3000 -w /path/to/wordlist.txt
```
Uses a custom wordlist for testing.

### **Increase Thread Count**
```bash
go run main.go http://localhost:3000 -t 20
```
Runs the scan with 20 concurrent threads.

---

## **Default Wordlist**
The tool includes a default wordlist at `wordlists/common.txt` containing commonly used commands. You can replace or update this wordlist with a custom file as needed.

To download a larger wordlist:
```bash
curl -o wordlists/advanced.txt https://raw.githubusercontent.com/danielmiessler/SecLists/master/Fuzzing/fuzzing-commands.txt
```

---

## **Performance Tips**
- Use a high thread count (`-t`) for faster scans, especially with larger wordlists.
- Enable verbose mode (`-v`) to get detailed feedback on which commands are being tested.
- Save results to an output file (`-o`) for easier analysis and reporting.

---

## **Planned Enhancements**
Future updates may include:
1. Advanced response analysis to detect subtle injection vulnerabilities.
2. Custom payload generation for specific command-line environments.
3. Integration with DNS logging for out-of-band detection.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for details.

---

## **Acknowledgments**
- Inspired by Juice Shop for testing and learning.
- Built with Go to demonstrate development and penetration testing skills.

---

## **Performance Tips**
- Test with larger wordlists to increase the likelihood of finding vulnerabilities.
- Save your results to a file for better organization in penetration testing reports.

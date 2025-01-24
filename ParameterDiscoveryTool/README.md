
# **ParameterDiscoveryTool**

**ParameterDiscoveryTool** is a lightweight Go-based tool designed for discovering hidden parameters in web applications. It uses a wordlist to test common parameter names against a target URL to identify potentially unused or hidden parameters.

---

## **Features**
- **Multithreaded Scanning**:
  - Conducts concurrent scans for faster parameter discovery.
- **Custom Wordlist Support**:
  - Allows users to supply their own wordlist of parameter names.
- **Verbose Mode**:
  - Enables detailed output for debugging and visibility.
- **Results Output**:
  - Saves discovered parameters to a file for easy reference.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/ParameterDiscoveryTool.git
   cd ParameterDiscoveryTool
   ```

2. **Build the Tool**:
   Compile the tool to a binary:
   ```bash
   go build -o parameter-discovery-tool main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./parameter-discovery-tool <target_url> [flags]
   ```

---

## **Usage**

### **Command Syntax**
```bash
go run main.go <target_url> -w <wordlist_file> -o <output_file> -t <threads> -v
```

### **Example**
```bash
go run main.go http://localhost:3000 -w ./wordlists/parameters.txt -o results.txt -t 10 -v
```

### **Options**
| Option             | Description                                         | Default               |
|---------------------|-----------------------------------------------------|-----------------------|
| `<target_url>`      | The target URL to scan for parameters               | *Required*            |
| `-w <wordlist>`     | Path to the wordlist file containing parameter names| *Required*            |
| `-o <output_file>`  | Path to save the results                            | `results.txt`         |
| `-t <threads>`      | Number of concurrent threads                        | `10`                  |
| `-v`                | Enable verbose mode for detailed output             | Disabled              |

---

## **Debugging Known Issues**
1. **Ensure the Wordlist File Exists**:
   Verify the wordlist file is present and accessible:
   ```bash
   ls -l <path_to_wordlist>
   ```
   Example:
   ```bash
   ls -l ./wordlists/parameters.txt
   ```

2. **Verify the Wordlist Path**:
   Ensure the correct file path is passed using the `-w` flag.

3. **Enable Verbose Mode for Debugging**:
   Use verbose mode to diagnose potential issues during the scan:
   ```bash
   go run main.go <target_url> -w <wordlist_file> -o <output_file> -t <threads> -v
   ```

---

## **Planned Enhancements**
Future updates will aim to improve the tool's usability and reliability:
1. **Enhanced Error Handling**:
   - Improve diagnostics for invalid inputs or missing files.
2. **Flag Validation**:
   - Ensure all required flags are provided before execution.
3. **Optimized File Reading**:
   - Handle larger wordlists and ensure seamless loading.
4. **Dynamic Parameter Discovery**:
   - Add features to detect and adapt to dynamic parameter names.

---

## **Contributions**
Contributions are welcome! If you have ideas for enhancements or find bugs, feel free to open an issue or submit a pull request.

---

## **License**
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## **Acknowledgments**
- Inspired by Juice Shop for testing and learning.
- Built with Go to showcase both development and penetration testing skills.

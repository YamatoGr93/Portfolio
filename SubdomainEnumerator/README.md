
# **Subdomain Enumerator**

A fast and efficient DNS-based subdomain enumeration tool built with Go. This tool is designed to help penetration testers and bug bounty hunters identify valid subdomains of a target domain using DNS brute-forcing techniques.

---

## **Features**
- **Custom Wordlist** (`-w`): Use your own wordlist for subdomain brute-forcing.
- **Save Results to File** (`-o`): Save discovered subdomains to a file for later use.
- **Silent Mode** (`-s`): Suppress console output for cleaner operation.
- **Parallelized DNS Lookups** (`-t`): Perform subdomain resolution using multiple threads for faster execution.
- **Default Configurations**: Includes a default wordlist (`wordlists/common.txt`) and output file (`results.txt`).

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/SubdomainEnumerator.git
   cd SubdomainEnumerator
   ```

2. **Build the Tool**:
   Compile the tool to a binary:
   ```bash
   go build -o subdomain-enumerator main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./subdomain-enumerator <domain> [flags]
   ```

---

## **Usage**

### **Command Syntax**
```bash
go run main.go <domain> [flags]
```
Replace `<domain>` with your target domain (e.g., `example.com`).

### **Available Flags**
| Flag            | Description                                     | Default                |
|------------------|-------------------------------------------------|------------------------|
| `-w`            | Path to the wordlist file                       | `wordlists/common.txt` |
| `-o`            | Path to the output file                         | `results.txt`          |
| `-s`            | Enable silent mode (no console output)          | `false`                |
| `-t`            | Number of threads for parallel DNS lookups      | `10`                   |

---

## **Examples**

### **Basic Usage**
```bash
go run main.go example.com
```

### **Specify a Custom Wordlist**
```bash
go run main.go example.com -w wordlists/custom.txt
```

### **Save Results to a File**
```bash
go run main.go example.com -o output.txt
```

### **Enable Silent Mode**
```bash
go run main.go example.com -s
```

### **Increase Thread Count**
```bash
go run main.go example.com -t 20
```

---

## **Default Wordlist**
The tool includes a default wordlist at `wordlists/common.txt` containing common subdomain names. You can replace or update this wordlist with a custom file as needed.

To download a larger wordlist:
```bash
curl -o wordlists/large.txt https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/DNS/subdomains-top1million-5000.txt
```

---

## **Performance Tips**
- Use a high thread count (`-t`) for faster scans, especially on larger wordlists.
- Use silent mode (`-s`) to reduce clutter during long-running scans.
- Save results to an output file (`-o`) to avoid losing discovered subdomains.

---

## **Planned Enhancements**
Future updates may include:
1. **CNAME and A Record Extraction**:
   - Extract and display CNAME and A records for identified subdomains.
2. **Integration with DNS APIs**:
   - Support DNS enumeration using APIs like SecurityTrails and VirusTotal.
3. **Wildcard Subdomain Detection**:
   - Identify and exclude wildcard subdomains from results.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for details.

---

## **Acknowledgments**
- Inspired by [SecLists](https://github.com/danielmiessler/SecLists) for wordlist resources.
- Built using Go for performance and concurrency.

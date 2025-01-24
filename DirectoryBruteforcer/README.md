
# **Directory Bruteforcer**

The **Directory Bruteforcer** is a high-performance tool designed for penetration testers and bug bounty hunters to discover hidden directories and files on web applications. Built with Go, it performs multithreaded directory brute-forcing using custom wordlists and supports saving results to an output file.

---

## **Features**
- **Custom Wordlist** (`-w`): Use your own wordlist for directory brute-forcing.
- **Save Results to File** (`-o`): Save discovered directories and files to a file.
- **Filter by Status Code** (`-s`): Filter results by HTTP status codes (e.g., `200`, `403`, `404`).
- **Multithreaded Scanning** (`-t`): Specify the number of threads for parallel requests.
- **Fast and Efficient**: Built with Go for high performance and optimized resource usage.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/DirectoryBruteforcer.git
   cd DirectoryBruteforcer
   ```

2. **Build the Tool**:
   Compile the tool to a binary:
   ```bash
   go build -o directory-bruteforcer main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./directory-bruteforcer <URL> [flags]
   ```

---

## **Usage**
```bash
go run main.go <URL> [flags]
```
Replace `<URL>` with the target URL (e.g., `http://localhost:3000` for Juice Shop).

### **Available Flags**
| Flag             | Description                                         | Default               |
|-------------------|-----------------------------------------------------|-----------------------|
| `-w`             | Path to the wordlist file                           | `wordlists/common.txt`|
| `-o`             | Path to the output file                             | `results.txt`         |
| `-s`             | Filter results by HTTP status code                  | *None*                |
| `-t`             | Number of threads for concurrent scanning           | `10`                  |

---

## **Examples**
### **Basic Usage**
```bash
go run main.go http://localhost:3000 -w wordlists/common.txt
```
Scans the target URL using the default wordlist.

### **Save Results to a File**
```bash
go run main.go http://localhost:3000 -w wordlists/common.txt -o output.txt
```
Saves the results to `output.txt`.

### **Filter by HTTP Status Code**
```bash
go run main.go http://localhost:3000 -w wordlists/common.txt -s 200
```
Displays only results with a `200 OK` status code.

### **Increase Thread Count**
```bash
go run main.go http://localhost:3000 -w wordlists/common.txt -t 20
```
Runs the scan with 20 concurrent threads.

---

## **Default Wordlist**
The tool includes a default wordlist at `wordlists/common.txt` containing common directory names. You can replace or update this wordlist with a custom file as needed.

To download a larger wordlist:
```bash
curl -o wordlists/large.txt https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/common.txt
```

---

## **Performance Tips**
- Use a higher thread count (`-t`) for faster scans.
- Filter by HTTP status codes (`-s`) to reduce noise in the results.
- Save results to an output file (`-o`) for easy reference.

---

## **Planned Enhancements**
Future updates may include:
- Recursive directory scanning.
- Support for POST requests.
- Output formats like JSON or Markdown for integration with other tools.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for details.

---

## **Acknowledgments**
- Inspired by Juice Shop for testing.
- Built with Go to showcase both development and penetration testing skills.

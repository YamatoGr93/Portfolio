
# **CSRF Token Extractor**

The **CSRF Token Extractor** is a lightweight Go-based tool designed for extracting CSRF tokens from web pages. This tool is useful for penetration testers and bug bounty hunters to identify CSRF token implementations and analyze their presence in web applications.

---

## **Features**
- **Automatic Token Extraction**:
  - Identifies CSRF tokens embedded in forms, meta tags, or hidden inputs.
- **Customizable Output**:
  - Saves extracted tokens to a specified file.
- **Efficient and Lightweight**:
  - Built with Go for speed and simplicity.

---

## **Installation**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/YamatoGr93/CSRFTokenExtractor.git
   cd CSRFTokenExtractor
   ```

2. **Build the Tool**:
   Compile the tool into a binary:
   ```bash
   go build -o csrf-token-extractor main.go
   ```

3. **Run the Tool**:
   Use the compiled binary:
   ```bash
   ./csrf-token-extractor [flags]
   ```

---

## **Usage**
```bash
go run main.go -url <target_url> [flags]
```
Replace `<target_url>` with the URL of the page to analyze.

### **Available Flags**
| Flag             | Description                                      | Default                |
|-------------------|--------------------------------------------------|------------------------|
| `-url`           | Target URL to extract CSRF tokens from           | *Required*            |
| `-o`             | File to save extracted CSRF tokens               | `csrf_tokens.txt`     |

---

## **Examples**
### **Extract Tokens from a Page**
```bash
go run main.go -url http://localhost:3000 -o csrf_tokens.txt
```
This command extracts CSRF tokens from the target URL and saves them to `csrf_tokens.txt`.

---

## **Default Behavior**
- If no CSRF tokens are found, the tool outputs: `No CSRF tokens found.`
- The extracted tokens are saved to the specified output file for future reference.

---

## **Debugging and Testing**
1. To debug the tool, print the response body by adding the following line to `main.go`:
   ```go
   fmt.Println(body)
   ```
2. Test the tool with a mock HTML page that includes CSRF tokens:
   ```html
   <input type="hidden" name="csrf-token" value="mock-csrf-token-value">
   ```

---

## **Planned Enhancements**
- **Enhanced Token Detection**:
  - Improve regex patterns to support additional CSRF token implementations.
- **Multi-Page Crawling**:
  - Add support for extracting tokens across multiple pages.
- **Output Formats**:
  - Include options for saving results in JSON or CSV format.

---

## **Limitations**
- The tool depends on common patterns to identify CSRF tokens and may not detect all implementations.
- If a web application doesn't use CSRF tokens, the tool will correctly report that no tokens are found.

---

## **License**
This project is licensed under the MIT License. See `LICENSE` for more details.

---

## **Acknowledgments**
- Inspired by common CSRF token extraction methods.
- Built with Go to demonstrate skills in both penetration testing and software development.

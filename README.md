# **YamatoGr93 Portfolio**

Welcome to my Go Developer portfolio! This repository showcases a diverse collection of projects I have developed using the Go programming language. These projects reflect my expertise in cybersecurity, software development, and automation. Each project is designed to solve specific problems, ranging from penetration testing tools to utility applications, and demonstrates my ability to create efficient, modular, and practical solutions.

---

## **Table of Contents**

1. [Overview](#overview)
2. [Projects Included](#projects-included)
3. [Highlights](#highlights)
4. [Installation](#installation)
5. [Usage Examples](#usage-examples)
6. [License](#license)

---

## **Overview**

This portfolio demonstrates my ability to:

- Design and implement custom tools and applications using Go.
- Solve complex problems in cybersecurity, automation, and software development.
- Develop modular and efficient code with multithreading and performance optimization.
- Automate repetitive tasks and improve workflows.
- Generate structured reports for professional use.

Each project includes detailed documentation in its respective subdirectory, covering usage instructions, features, and examples.

---

## **Projects Included**

### **Cybersecurity Tools**
1. **Command Injection Tester**
   - Identifies potential command injection vulnerabilities by testing crafted payloads.
   - Features: Custom wordlists, result filtering, automated testing.

2. **Directory Bruteforcer**
   - Finds hidden directories and files through brute-forcing.
   - Features: Multithreaded scanning, status code filtering, custom wordlist support.

3. **HTTP Header Analyzer**
   - Analyzes HTTP response headers for missing security configurations.
   - Features: Checks for CSP, HSTS, etc., with actionable recommendations.

4. **Open Redirect Finder**
   - Detects open redirect vulnerabilities by testing common redirect parameters.
   - Features: Custom wordlists and payloads.

5. **Parameter Discovery Tool**
   - Discovers hidden parameters in web applications.
   - Features: Brute-forces parameters using a wordlist and identifies unused ones.

6. **SSRF Tester**
   - Tests for Server-Side Request Forgery vulnerabilities using common payloads.

7. **Subdomain Enumerator**
   - Discovers valid subdomains of a target domain via DNS brute-forcing.
   - Features: Multithreading, custom wordlist support.

8. **XSS Payload Generator**
   - Generates custom and encoded XSS payloads for testing purposes.

9. **CSRF Token Extractor**
   - Extracts CSRF tokens from web pages using regex-based extraction.

10. **Vulnerability Report Generator**
    - Converts JSON-based findings into structured Markdown vulnerability reports.

---

### **Utility Applications**
1. **Automated Backup System**
   - Automates file backups with scheduling options for data safety.

2. **File System Organizer**
   - Organizes files in directories based on extensions or custom rules.

3. **Password Generator**
   - Creates secure passwords with customizable length and complexity.

4. **Text Search Tool**
   - Searches for specific text patterns across files in a directory.

5. **URL Shortener CLI**
   - A command-line tool to create short URLs for long links.

6. **Task Manager CLI**
   - A simple CLI-based task manager to track daily tasks efficiently.

7. **System Resource Monitor**
   - Monitors system resources such as CPU usage, memory usage, etc., in real-time.

8. **Data Parser**
   - Parses and extracts data from various formats (e.g., JSON, XML).

9. **Simple Chatbot**
   - A basic chatbot application that responds to user inputs interactively.

---

## **Highlights**

- **Total Projects Developed**: 19
- **Technologies Used**: Go programming language.
- **Key Skills Demonstrated**:
  - Cybersecurity expertise (vulnerability analysis and exploitation).
  - Efficient multithreading and performance optimization.
  - Automation of repetitive tasks.
  - Clean coding practices with modular design.
  - Professional report generation for penetration testing or other purposes.

---

## **Installation**

1. Clone the repository:
  git clone https://github.com/YamatoGr93/Portfolio.git
  cd Portfolio


2. Navigate to any project's directory and build it:
  cd ProjectName
  go build -o project-binary main.go

3. Run the project using the appropriate flags (refer to individual project documentation).

---

## **Usage Examples**

### Example 1: SSRF Tester
  go run main.go -url http://localhost:3000 -param url -o ssrf_results.txt


### Example 2: File System Organizer
  go run main.go --source /path/to/source --destination /path/to/destination 


### Example 3: Password Generator
   go run main.go --length 16 --special-chars true --output password.txt 



---

## **License**

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## **Acknowledgments**

- Inspired by real-world challenges in cybersecurity and software development.
- Special thanks to [OWASP Juice Shop](https://owasp.org/www-project-juice-shop/) for providing a safe environment for testing security tools.
- Built entirely with Go to showcase my development skills and problem-solving abilities.

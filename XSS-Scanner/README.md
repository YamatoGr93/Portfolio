📌 XSS Scanner

🔍 Overview

This tool scans web applications for Reflected XSS vulnerabilities by injecting malicious payloads and checking if they execute. It was designed to test OWASP Juice Shop and similar applications.

💡 Features

Automated XSS Testing → Sends payloads to web forms & parameters.

Custom Payloads → Uses a set of pre-defined XSS attack strings.

Feedback Page Detection → Verifies stored XSS by retrieving feedback submissions.

Logging & Error Handling → Captures failures and request errors for analysis.

Planned Headless Browser Support → Future update to check for actual execution using Chromedp.

⚙️ Installation

1️⃣ Clone the Repository

git clone https://github.com/YamatoGr93/Portfolio.git && cd Portfolio/XSS-Scanner

2️⃣ Build the Docker Image

docker build --no-cache -t xss-scanner .

3️⃣ Run the Scanner

docker run --rm xss-scanner http://172.17.0.1:3000/api/Feedbacks/

🛠️ How It Works

1️⃣ Sends XSS payloads via a POST request to a vulnerable endpoint (/api/Feedbacks/).
2️⃣ Waits a few seconds for stored XSS to appear.
3️⃣ Requests the feedback display page (/#/contact) to check if the payload is reflected.
4️⃣ Reports findings:

✅ No XSS found

🔥 XSS detected at <URL>

⚠️ Known Issues & Error Handling

❌ Issue: No XSS Detected in Automated Scans

Reason: The tool only checks raw HTML responses, but Juice Shop renders feedback via JavaScript.

✅ Planned Fix: Use a headless browser (Chromedp) to check for JavaScript execution.

❌ Issue: Some Requests Fail

Reason: Network delays or API rate-limits may cause failed requests.

✅ Planned Fix: Implement request retries (e.g., retry up to 3 times before marking a failure).

❌ Issue: No Detailed Debugging Logs

Reason: The tool does not log full request/response data for debugging.

✅ Planned Fix:

Add an option to log failed responses.

Save detailed request payloads for analysis.

❌ Issue: Payloads Might Be Filtered or Escaped

Reason: Juice Shop might be filtering some XSS payloads before reflecting them.

✅ Planned Fix:

Test different payload variations (e.g., bypassing filters with event handlers like onfocus, onmouseover).

Implement dynamic payload encoding (e.g., Base64 obfuscation).

📌 Future Plans & Improvements

🚀 Headless Browser Execution → Detect XSS beyond reflection, using Chromedp to verify JavaScript execution.
📊 Detailed Logging & Report Generation → Store results in a structured format (JSON, Markdown).
🔄 Automated Retesting → Run scans periodically & compare new findings.

📜 License

This tool is released under the MIT License. Free to use, modify, and improve!

👨‍💻 Author

Panagiotis Lampros🔗 GitHub🔗 LinkedIn📧 Contact: yamatoGr93@proton.me


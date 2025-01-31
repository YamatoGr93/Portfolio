# Automated Backup System

Automated Backup System is a simple Go application that periodically backs up files from a specified directory into a zip archive.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/AutomatedBackupSystem.git
    cd AutomatedBackupSystem
    ```

2. Build the application:

    ```sh
    go build -o autobackup main.go
    ```

### Usage

1. Ensure you have a directory named `files` in the same directory as the `main.go` file. This directory should contain the files you want to back up.

2. Run the application:

    ```sh
    ./autobackup
    ```

The application will create a zip archive of the `files` directory every 24 hours. The archive will be named `backup-<timestamp>.zip`.

### Example

Given the following directory structure:

AutomatedBackupSystem/ ├── main.go ├── autobackup (built executable) └── files/ ├── file1.txt ├── file2.txt └── file3.txt

When you run the application, it will create a zip archive named `backup-<timestamp>.zip` containing `file1.txt`, `file2.txt`, and `file3.txt`.

### License

This project is licensed under the MIT License - see the LICENSE file for details.
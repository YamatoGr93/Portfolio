# System Resource Monitor

System Resource Monitor is a simple Go application that monitors and prints the CPU, memory, and disk usage of the system.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Linux operating system (as the code uses `/proc` filesystem and `syscall` package)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/SystemResourceMonitor.git
    cd SystemResourceMonitor
    ```

2. Build the application:

    ```sh
    go build -o sysresmon main.go
    ```

### Usage

1. Run the application:

    ```sh
    ./sysresmon
    ```

The application will print the CPU, memory, and disk usage of the system.

### Example

When you run the application, it will output something like:

```sh
CPU Usage: 15.23%
Memory Usage: 45.67%
Disk Usage: 78.90%
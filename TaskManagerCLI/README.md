# Task Manager CLI

Task Manager CLI is a simple Go application that allows you to manage a list of tasks through the command line.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/TaskManagerCLI.git
    cd TaskManagerCLI
    ```

2. Build the application:

    ```sh
    go build -o taskmanager main.go
    ```

### Usage

1. Run the application:

    ```sh
    ./taskmanager
    ```

2. Use the following commands to manage your tasks:
    - `add`: Add a new task. You will be prompted to enter the task description.
    - `list`: List all tasks.
    - `exit`: Exit the application.

### Example

```sh
Enter command (add/list/exit): add
Task: Buy groceries
Enter command (add/list/exit): add
Task: Walk the dog
Enter command (add/list/exit): list
1. Buy groceries
2. Walk the dog
Enter command (add/list/exit): exit
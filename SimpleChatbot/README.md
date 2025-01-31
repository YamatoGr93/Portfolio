# Simple Chatbot

Simple Chatbot is a basic Go application that interacts with the user through the command line. It responds to predefined inputs with predefined responses.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/SimpleChatbot.git
    cd SimpleChatbot
    ```

2. Build the application:

    ```sh
    go build -o chatbot main.go
    ```

### Usage

1. Run the application:

    ```sh
    ./chatbot
    ```

2. Interact with the chatbot by typing messages. The chatbot understands the following commands:
    - `hello`: Responds with "Hi there!"
    - `bye`: Responds with "Goodbye!"
    - `exit`: Exits the application

For any other input, the chatbot will respond with "I don't understand."

### Example

```sh
You: hello
Bot: Hi there!
You: bye
Bot: Goodbye!
You: exit
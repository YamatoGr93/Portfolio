# Password Generator

Password Generator is a simple Go application that generates a random password of a specified length using a mix of alphanumeric characters and special symbols.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/PasswordGenerator.git
    cd PasswordGenerator
    ```

2. Build the application:

    ```sh
    go build -o passwordgen main.go
    ```

### Usage

1. Run the application:

    ```sh
    ./passwordgen
    ```

The application will generate and print a random password of length 12.

### Example

When you run the application, it will output a random password like:

```sh
aB3!dE5@fG7#hI

Code Explanation
The generatePassword function generates a random password of the specified length using characters from the set abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*(). It uses the crypto/rand package to generate cryptographically secure random numbers.

License
This project is licensed under the MIT License - see the LICENSE file for details.

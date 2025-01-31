# DataParser

DataParser is a simple Go application that parses CSV and JSON files and prints their contents to the console.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/DataParser.git
    cd DataParser
    ```

2. Build the application:

    ```sh
    go build -o dataparser main.go
    ```

### Usage

1. Place your [data.csv](http://_vscodecontentref_/1) and [data.json](http://_vscodecontentref_/2) files in the same directory as the [main.go](http://_vscodecontentref_/3) file.

2. Run the application:

    ```sh
    ./dataparser
    ```

### Example

Given the following [data.csv](http://_vscodecontentref_/4):

```csv
name,age
Alice,30
Bob,25

And the following: 

{
    "name": "Alice",
    "age": 30
}

The application will output:

[[name age] [Alice 30] [Bob 25]]
map[age:30 name:Alice]

License
This project is licensed under the MIT License - see the LICENSE file for details.
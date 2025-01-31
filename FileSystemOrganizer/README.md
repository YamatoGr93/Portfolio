# FileSystemOrganizer

FileSystemOrganizer is a simple Go program that organizes files in the current directory into subdirectories based on their file extensions.

## Usage

1. Ensure you have Go installed on your system.
2. Clone this repository or copy the `main.go` file to your local machine.
3. Navigate to the directory containing `main.go`.
4. Run the program using the following command:

    ```sh
    go run main.go
    ```

## How it works

- The program reads all files in the current directory.
- For each file, it determines the file extension.
- It creates a subdirectory named after the file extension (e.g., `.txt` files go into a `txt` directory).
- It moves the file into the corresponding subdirectory.
- If a file has no extension, it is moved to a `no_extension` directory.

## Error Handling

- If the program fails to read the directory, it logs a fatal error and exits.
- If the program fails to create a subdirectory, it logs the error and continues with the next file.
- If the program fails to move a file, it logs the error and continues with the next file.
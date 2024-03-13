# hexreplacer

A simple command-line tool to find and replace hex strings within files.

## Description

This tool allows you to quickly replace specific hex strings with new ones within a given file. It performs the following tasks:

1. Reads the input file.
2. Decodes the provided old and new hex strings.
3. Checks for the existence of the old hex string in the file.
4. Creates a backup of the original file with a ".bak" extension.
5. Performs the replacement in memory.
6. Writes the modified content back to the original file.
7. Prints a success message and the path to the backup file.

## Installation

**Go installation:** Ensure you have Go installed on your system. Refer to the official Go documentation for installation instructions: [https://go.dev/doc/install](https://go.dev/doc/install).

**Installation from Source:**

1. Clone the repository:

```bash
git clone https://github.com/dimasmaulana/hexreplacer.git
```

1. Navigate to the cloned directory:

```bash
cd hexreplacer
```

1. Execute the following command to install:

```bash
go install
```

## Usage

```bash
hexreplacer <file> <oldHexString> <newHexString>
```

**Example:**

```bash
hexreplacer myfile.txt 48656C6C6F 576F726C64
```

## Backups

The tool automatically creates a backup of the original file with a ".bak" extension before making any modifications.

## Error Handling

The tool handles various errors, such as invalid hex strings, file access issues, and more. It prints informative error messages to guide the user.

## License

This tool is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Pull requests are welcome! Feel free to contribute to the project by submitting bug fixes, enhancements, or new features.

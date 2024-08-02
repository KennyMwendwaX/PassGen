# PassGen

A simple, secure command-line password generator written in Go using the Cobra library.

## Features

- Generate passwords of customizable length (8-128 characters)
- Option to include letters, numbers, and special characters
- Secure random generation using crypto/rand
- Calculates and displays password entropy (in verbose mode)
- Easy-to-use command-line interface

## Installation

1. Ensure you have Go installed on your system (version 1.11 or later recommended).
2. Clone this repository:

```bash
git clone https://github.com/KennyMwendwaX/PassGen.git
```

3. Navigate to the project directory:

```bash
cd PassGen
```

4. Build the project:

```bash
 cd PassGen
```

## Usage

After building the project, you can run the password generator using:

```bash
./PassGen [flags]
```

### Flags

- `-l, --length int`: Length of the password (default 16)
- `-t, --letters`: Include letters (default true)
- `-n, --numbers`: Include numbers (default true)
- `-s, --specials`: Include special characters
- `-v, --verbose`: Verbose output (displays password entropy)

### Examples

1. Generate a password with default settings (16 characters, letters and numbers):

```bash
./PassGen
```

2. Generate a 20-character password including special characters:

```bash
./PassGen -l 20 -s
```

3. Generate a password and display its entropy:

```bash
./PassGen -v
```

## Security

This password generator uses Go's `crypto/rand` package to ensure cryptographically secure random number generation. The generated passwords are suitable for most security purposes.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

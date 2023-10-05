---

# BeSafe by PortySec

BeSafe is a comprehensive security tool written in Go, designed to assist with various security tasks. Whether you're looking to hash data, scan ports, or validate password strength, BeSafe has got you covered.

## Features

### 1. **PassGuard**

- **Description**: Assess the strength of a password.
- **Usage**:
  ```
  porty passguard <password>
  ```
  - **Example**:
  ```
  porty passguard 1234@abcd
  ```
  - **Output**:
  ```
  Password Is Strong Enough
  ```
  or
  ```
  Password must Contain [A-Z,a-z,0-9,chars] and have a length of at least 12
  ```

### 2. **Scanner**

- **Description**: Check connectivity to a host on a specific port.
- **Usage**:
  ```
  porty scanner <host> <port>
  ```
  - **Example**:
  ```
  porty scanner google.com 80
  ```
  - **Output**:
  ```
  Connection has been Successfully established
  ```

### 3. **HashMe**

- **Description**: Hash text or file content using various algorithms with a default of `sha256`.
- **Usage**:
  ```
  porty hashme [options] <input>
  ```
  - **Options**:
    - `-a, --algorithm <algorithm>`: Specify the hashing algorithm. Available options are `sha256`, `sha512`, and `sha1`. If not specified, `sha256` is used by default.
    - `-f, --file <filename>`: Hash the content of a file instead of plain text.
  - **Examples**:
    - Hashing a text with the default algorithm (`sha256`):
      ```
      porty hashme helloworld
      ```
    - Hashing a text using `sha512`:
      ```
      porty hashme -a sha512 helloworld
      ```
    - Hashing the content of a file using the default algorithm (`sha256`):
      ```
      porty hashme -f myfile.txt
      ```
    - Hashing the content of a file using `sha1`:
      ```
      porty hashme -a sha1 -f myfile.txt
      ```

Note: If neither `-a` nor `-f` is specified, the input is treated as plain text and hashed using the default `sha256` algorithm.

## Installation

1. Clone the repository:
```
git clone https://github.com/PortySec/BeSafe.git
```
2. Navigate to the project directory:
```
cd BeSafe
```
3. Build the project:
```
go build
```

## Contribution

Feel free to contribute to this project by opening issues or submitting pull requests. All contributions are welcomed!

## License

This project is licensed under the MIT License.
```

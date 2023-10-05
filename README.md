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
- **Description**: Hash text or file content.
- **Usage**: 
  ```
  porty hashme -a <algorithm> [options]
  ```
  - **Options**:
    - `-a, --algorithm <algorithm>`: Specify the hashing algorithm (e.g., sha256, sha512, sha1).
    - `-f, --file <filename>`: Hash the content of a file.
  - **Examples**: 
  ```
  porty hashme -a sha256 helloworld
  porty hashme -a sha256 -f myfile.txt
  ```

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

---

You can copy the content between the backticks and paste it into your README.md file on GitHub.

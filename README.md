# Credit Card Management Utility

This project is a comprehensive tool for managing, validating, and generating credit card numbers. It leverages the Luhn algorithm for validation and provides various functionalities to work with credit card data.

## Features
- **Validation**: Check the correctness of credit card numbers using the Luhn algorithm.
- **Generation**: Generate valid credit card numbers based on a given pattern.
- **Information**: Identify the brand and issuer of a credit card.
- **Issuance**: Generate unique credit card numbers for specific brands and issuers.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/eraga0423/ValidCreditcard.git
2. Ensure you have Go (1.18 or later) installed.

## Compilation
Compile the project with:
```bash
go build -o creditcard
```

## Usage
### Commands
- `validate`: Validate credit card numbers.
- `generate`: Generate valid credit card numbers.
- `information`: Retrieve information about a credit card (brand and issuer).
- `issue`: Generate a unique credit card number for a specific brand and issuer.

### Examples
1. **Validate Credit Card Numbers**
   ```bash
   echo "4400431234567890" | ./creditcard validate --stdin
   ```

2. **Generate Credit Card Numbers**
   ```bash
   ./creditcard generate --pick 440043****
   ```

3. **Get Card Information**
   ```bash
   ./creditcard information --brands=brands.txt --issuers=issuers.txt 4400431234567890
   ```

4. **Issue a Credit Card**
   ```bash
   ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer=Kaspi Gold
   ```

## Project Structure
- `main.go`: Entry point for command processing.
- `generator.go`: Handles credit card number generation.
- `input.go`: Processes input and reads data from files.
- `output.go`: Formats and outputs results.
- `luhna.go`: Implements the Luhn algorithm for validation.
- `brands.txt`: Stores mappings of credit card brands and prefixes.
- `issuers.txt`: Stores mappings of issuers and prefixes.

## Data Files
- **brands.txt**: Contains brand data in the format `BRAND:PREFIX`.
- **issuers.txt**: Contains issuer data in the format `ISSUER:PREFIX`.

## License
This project is open-source and available for use and modification.


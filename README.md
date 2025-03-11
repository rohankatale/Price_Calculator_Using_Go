# Price_Calculator_using_GO: Tax Included Price Calculator

Welcome to the Price_Calculator_using_GO repository. This project is a simple tax-included price calculator built with Go. It demonstrates several key concepts in Go programming and provides a modular project structure.

## Overview

- **Purpose:**  
  Calculate tax-included prices based on dynamic tax rates and input prices. The project reads prices from an input file, processes them, and writes the results to different JSON files based on the tax rate.

- **Key Concepts:**
  - **Interfaces:**  
    The project uses an `IOManager` interface to abstract file I/O. This helps in creating alternative IO mechanisms (e.g., command-line or file-based).
  - **Structs and Methods:**  
    The primary logic for processing prices is encapsulated in a struct (`TaxIncludedPriceJob`) with methods like `Process()`, demonstrating object-like behavior.
  - **Error Handling:**  
    Uses standard Go error handling practices across the modules.
  - **Modular Design:**  
    The code is divided into separate packages such as `prices`, `filemanager`, `conversion`, and `iomanager`, promoting separation of concerns.
  - **JSON Handling:**  
    Showcases encoding/decoding of data with JSON, useful for configuration and result files.

## Folder Structure

- **/Prices**: Contains the logic for processing price data.
- **/filemanager**: Handles reading from and writing to JSON files.
- **/iomanager**: Defines the interface used to abstract IO operations.
- **/Conversion**: Provides conversion functions (e.g., converting strings to floats).
- **/main.go**: Orchestrates the reading of price inputs, processing with various tax rates, and saving the results.
- **Additional Files:**  
  Other files and folders cover basic tutorials and advanced topics in Go, like pointers, structs, recursion, slices, maps, and more.

## How to Run

1. Make sure you have Go installed (minimum version 1.24.0).
2. Open your terminal, navigate to the project directory:
   ```
   cd /Desktop/go/Project
   ```
3. Build and run the project:
   ```
   go run main.go
   ```
4. Check the generated result JSON files (e.g., `result_10.json`, `result_07.json`) for output.

## Conclusion

This project is designed to help learners explore fundamental Go programming concepts such as interfaces, modular design, JSON data handling, and error management. Contributions and suggestions for improvements are welcome!

Enjoy your learning journey with Go!

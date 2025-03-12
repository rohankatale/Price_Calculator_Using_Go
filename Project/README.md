# Price_Calculator_using_GO: Tax Included Price Calculator

Welcome to the Price_Calculator_using_GO repository. This project is a simple tax-included price calculator built with Go. It demonstrates several key concepts in Go programming and provides a modular project structure.

## Overview

- **Purpose:**  
  Calculate tax-included prices based on dynamic tax rates and input prices. The project reads prices from an input file, processes them concurrently using goroutines and channels, and writes the results to different JSON files based on the tax rate.

- **Key Concepts:**
  - **Interfaces:**  
    The project uses an `IOManager` interface to abstract file I/O, allowing for various implementations (e.g., command-line or file-based).
  - **Structs and Methods:**  
    The core logic is encapsulated in a struct (`TaxIncludedPriceJob`) with methods like `Process()`, which demonstrates object-like behavior.
  - **Error Handling:**  
    Standard Go error handling practices are used across modules.
  - **Modular Design:**  
    The code is organized into separate packages (`prices`, `filemanager`, `conversion`, `iomanager`), promoting clear separation of concerns.
  - **JSON Handling:**  
    The project uses JSON encoding/decoding for configuration and result files.
  - **Concurrency:**  
    Using goroutines and channels, the project processes multiple tax rates simultaneously to improve performance and efficiency.

## Folder Structure

- **/Prices**: Contains the logic for processing price data.
- **/filemanager**: Manages reading from and writing to JSON files.
- **/iomanager**: Defines the interface for I/O operations.
- **/Conversion**: Provides utility functions such as converting strings to floats.
- **/main.go**: Orchestrates input reading, concurrent processing using multiple tax rates, and result saving.
- **Additional Files:**  
  Other files and directories illustrate various Go language concepts, including pointers, structs, recursion, slices, maps, concurrency, and more.

## How to Run

1. Install Go (minimum version 1.24.0).
2. Open your terminal and navigate to the project directory:
   ```
   cd /Desktop/go/Project
   ```
3. Build and run the project:
   ```
   go run main.go
   ```
4. The resulting JSON files (e.g., `result_10.json`, `result_07.json`) will be generated for each tax rate.

## Conclusion

This project is designed to help learners explore fundamental Go programming concepts such as interfaces, modular design, JSON data handling, and error management, with a special focus on concurrent processing. Contributions and suggestions are welcome!

Enjoy your learning journey with Go!

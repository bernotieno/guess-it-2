# Guess-It-2

**Guess-It-2** is a Go application designed to predict future data points using statistical analysis and linear regression. It reads a series of numerical inputs, performs calculations based on a sliding window of recent values, and provides a predicted range with a confidence interval.

## Table of Contents

- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [Author](#author)

## Installation

To use this project, you need to have Go installed. Follow the steps below to get started:

1. Clone the repository:
   ```sh
   git clone https://learn.zone01kisumu.ke/git/bernaotieno/guess-it-2.git
   cd guess-it-2
   ```


## Project Structure
 * `student/calculation/linearreg.go`: Calculates the slope and intercept of the data provided
 * `student/calculation/mean.go`: Calculates the mean of a slice of float64 numbers.
 * `student/calculation/range.go`: Calculates the range based on the linear regresion
 * `student/calculation/stddeviation.go`: Calculates the standard deviation from variance.
 * `student/calculation/variance.go`: Calculates the variance of a slice of float64 numbers.

 * `student/main.go`: The main entry point for the application.

## Usage

1. Change the permission of the script file to executable:

```sh
chmod +x ./student/script.sh
```
2. To run the project, use the provided bash script:

```sh
./student/script.sh
```
3. Enter numbers one by one and press Enter. The program will calculate and print the range after each entry. To exit, simply press `crl`+`c`.

## Testing
**To test the project, follow these steps:**

 1. Download the test [zip file](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip) containing the tester.
 2. Extract the zip file in the root directory of the project.
 3. Move the student/ folder into the extracted directory.
 4. Follow the instructions provided in the tester's README file.

## Contributing
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## Author
 * [Benard Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)

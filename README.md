# golang_webdriver_version

This is a simple go program to demosntrate how to use the selenium webdriver, named "agouti" in golang.

## Requirements

- Go 1.20 or later
- Chrome or Firefox installed

## Installation

```bash
go get github.com/sclevine/agouti
```

## Usage

You need create a ".env" file with the following variables:

```bash
tee -a .env <<EOF
# .env
WEB_USERNAME=your_username
WEB_PASSWORD=your_password
EOF
```

Because the ".env" file contains sensitive information, it is ignored by Git.

To run the program, execute the following command:

```bash
go run main.go
```

or build the program and execute the binary:

```bash
go build -o webDriver
./webDriver
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://pt.wikipedia.org/wiki/Licen%C3%A7a_MIT) file for details.

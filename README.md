# golang_webdriver_version

This is a simple go program to demosntrate how to use the selenium webdriver, named "agouti" in golang.

## Requirements

- Go 1.20 or later
- Chromedriver 2.33 or later

## Installation

You can install the "chromedriver" with the following command:

For Linux:

```bash
sudo apt-get install chromium-chromedriver
```

For Windows:

```bash
choco install chromedriver
```

You don't use the "choco" package manager, you can download the "chromedriver" from the following link:

[https://sites.google.com/a/chromium.org/chromedriver/downloads](https://sites.google.com/a/chromium.org/chromedriver/downloads)

You can install the "agouti" package with the following command:
NOTE: If you build the program the agouti package will be included in the binary file.

```bash
go get github.com/sclevine/agouti
```

## Usage

You need create a ".env" file with the following variables:

```bash
tee -a .env <<EOF
# .env
WEB_USERN=your_username
WEB_PASSWORD=your_password
WEB_URL=https://your_url.com
EOF
```

Because the ".env" file contains sensitive information, it is ignored by Git.

You can run the program with:

```bash
go run main.go
```

But I recommend to build the program with the following command:

For Linux:

```bash
go build -o webDriver .
```

For Windows:

```bash
go build -o webDriver.exe .
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://pt.wikipedia.org/wiki/Licen%C3%A7a_MIT) file for details.

# Go Programming Environment Setup Guide

Follow these steps to set up your environment for Go programming on Linux, macOS, or Windows.

---

## 1. Install Go

### Linux

```sh
sudo apt update
sudo apt install golang-go
```

Or download the latest version from the [official Go downloads page](https://go.dev/dl/) and follow the installation instructions.

### macOS

- Using Homebrew:
  ```sh
  brew install go
  ```
- Or download from [Go downloads page](https://go.dev/dl/).

### Windows

- Download the installer from [Go downloads page](https://go.dev/dl/).
- Run the installer and follow the prompts.

---

## 2. Verify Installation

```sh
go version
```
You should see the installed Go version.

---

## 3. Set Up Your Workspace

- Create a directory for your Go projects:
  ```sh
  mkdir ~/go
  ```
- Set the `GOPATH` environment variable (optional for Go 1.11+ with modules, but useful for legacy projects):
  ```sh
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOPATH/bin
  ```
  Add these lines to your `~/.bashrc`, `~/.zshrc`, or equivalent shell config.

---

## 4. Install Visual Studio Code (Recommended)

- Download and install [VS Code](https://code.visualstudio.com/).
- Install the **Go extension** from the Extensions Marketplace for code completion, linting, and debugging.

---

## 5. Test Your Setup

Create a file named `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:

```sh
go run hello.go
```

You should see:  
`Hello, Go!`

---

## 6. Explore Further

- Try running exercises in the `content/` directory.
- Use `go mod init <modulename>` to enable Go modules for new projects.

---

You are now ready to start coding in Go!
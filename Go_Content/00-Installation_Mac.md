# Installing Go on a Mac

# 1️⃣ Choose the installation method

**Option A — Official macOS pkg installer**
**Option B — Homebrew** (recommended if you already use Homebrew)

---

# 2️⃣ Install Go using the official pkg installer

1. Open your browser and go to:
   [https://golang.org/dl](https://golang.org/dl)

2. Download the **macOS package** (`.pkg`) for your architecture:

   * Intel / AMD64 → `go1.xx.x.darwin-amd64.pkg`
   * Apple Silicon (M1/M2) → `go1.xx.x.darwin-arm64.pkg`

3. Open the downloaded `.pkg` file and follow the installer prompts.

   * Keep the default install path (`/usr/local/go`).
   * The installer automatically sets `/usr/local/go/bin` in your PATH.

4. Close the installer.

---

# 3️⃣ Install Go using Homebrew (alternative method)

If you have **Homebrew**, this is simpler:

```bash
brew update
brew install go
```

To upgrade Go later:

```bash
brew upgrade go
```

---

# 4️⃣ Verify the installation

Open **Terminal** and run:

```bash
go version
```

Example output:

```
go version go1.22 darwin/amd64
```

Then check Go environment variables:

```bash
go env
```

Key values:

* `GOROOT` → Go installation path (usually `/usr/local/go`)
* `GOPATH` → default workspace (`$HOME/go`)
* `GOMODCACHE` → module cache location

---

# 5️⃣ Set up your workspace

Go modules are default now, but having a workspace helps:

```bash
mkdir -p $HOME/go
mkdir -p $HOME/projects
```

* `$HOME/go` → optional GOPATH workspace
* `$HOME/projects` → your Go projects

Optional: set `GOPATH` in your shell:

```bash
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
source ~/.zshrc
```

*(If you use bash, replace `.zshrc` with `.bash_profile`.)*

---

# 6️⃣ Test your Go installation

1. Create a project folder:

```bash
cd ~/projects
mkdir hello-go
cd hello-go
```

2. Create a file called `main.go`:

```bash
nano main.go
```

Paste:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go on macOS!")
}
```

Save (`Ctrl+O`) and exit (`Ctrl+X`).

3. Run your program:

```bash
go run main.go
```

You should see:

```
Hello, Go on macOS!
```

4. Build an executable:

```bash
go build -o hello
./hello
```

---

# 7️⃣ Initialize Go modules (recommended)

Inside your project folder:

```bash
go mod init github.com/yourusername/hello-go
go mod tidy
```

* Creates `go.mod` for module-aware projects.
* Handles dependencies cleanly.

---

# 8️⃣ Install VS Code & Go extension (recommended)

1. Install **Visual Studio Code**: [https://code.visualstudio.com/](https://code.visualstudio.com/)
2. Open your project folder in VS Code
3. Install **Go extension** (by the Go team)
4. Accept prompts to install extra Go tools (gopls, dlv, gofmt, etc.)

---

# 9️⃣ Common troubleshooting

| Problem                 | Solution                                                          |
| ----------------------- | ----------------------------------------------------------------- |
| `go: command not found` | Make sure `/usr/local/go/bin` is in PATH; restart Terminal        |
| Wrong architecture      | Uninstall Go and reinstall the correct `amd64` or `arm64` version |
| Permission issues       | Don’t install or run projects in `/System` directories            |
| `go get` blocked        | Check network/proxy settings; set GOPROXY if needed               |

---

# ✅ Quick checklist

1. Install Go: `.pkg` installer or `brew install go`
2. Verify: `go version`
3. Set up workspace: `~/go` and `~/projects`
4. Create test program: `main.go`
5. Run program: `go run main.go`
6. Initialize module: `go mod init <module>`
7. Install VS Code + Go extension

---

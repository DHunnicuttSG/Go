# Installing Go on Windows

# 1. Choose the right Go installer

1. Determine your Windows architecture:

   * 64-bit Intel/AMD → **amd64**
   * ARM64 (some Surface devices) → **arm64**

2. Recommended: use the official MSI installer for simplicity.

---

# 2. Download & install (MSI method — easiest)

1. Open your browser and go to the official Go downloads page (golang.org/dl).

   * Download the Windows MSI that matches your architecture (e.g. `go1.xx.x.windows-amd64.msi`).

2. Run the downloaded `.msi` file:

   * Accept the license.
   * Keep the default installation location (usually `C:\Program Files\Go`).
   * Let the installer add Go to your `PATH` (it does this automatically).

3. Finish the installer and close the installer window.

---

# 3. Alternative: Install with winget (command line)

If you prefer a one-line install and have Windows Package Manager:

Open **PowerShell** (as your normal user — admin not required for `winget install` if allowed):

```powershell
winget install -e --id Go.Go
```

This installs the latest Go and usually handles PATH for you.

---

# 4. Verify the installation

Open a new **PowerShell** or **Command Prompt** (important: must be *new* so environment updates are picked up) and run:

```powershell
go version
```

Expected output example:

```
go version go1.22 windows/amd64
```

If you see version info, Go is installed and on your PATH.

---

# 5. Set up your workspace (recommended)

Modern Go uses **modules**, so you don’t have to use GOPATH, but setting a personal GOPATH is still useful.

1. Create a directory for your projects (example):

```powershell
mkdir $env:USERPROFILE\go
mkdir $env:USERPROFILE\projects
```

2. (Optional) Set GOPATH to `%USERPROFILE%\go`. The MSI often leaves GOPATH unset so the default is `%USERPROFILE%\go`. To explicitly set it:

* Open **Start → Edit environment variables for your account**
* Under *User variables* click **New** (or **Edit** if it exists):

  * Name: `GOPATH`
  * Value: `C:\Users\<YourUserName>\go` (or `%USERPROFILE%\go`)
* Make sure `C:\Program Files\Go\bin` is in your `Path` (installer usually did this).

After editing env vars, open a new terminal.

---

# 6. Confirm `go env` values

Run:

```powershell
go env
```

Check these important fields:

* `GOROOT` — where Go is installed (e.g. `C:\Program Files\Go`)
* `GOPATH` — your workspace (e.g. `C:\Users\You\go`) — may be empty and default used
* `GOMODCACHE`, `GOMOD` etc.

---

# 7. Write and run your first program

1. Create a project folder and a file:

```powershell
cd $env:USERPROFILE\projects
mkdir hello-go
cd hello-go
notepad main.go
```

2. Paste this into `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go on Windows!")
}
```

Save and close Notepad.

3. Run it:

```powershell
go run main.go
```

You should see:

```
Hello, Go on Windows!
```

4. Build an executable:

```powershell
go build -o hello.exe
.\hello.exe
```

---

# 8. Initialize a module (recommended for new projects)

Inside your project directory:

```powershell
go mod init github.com/yourusername/hello-go
go mod tidy
```

This creates `go.mod` and lets you import external packages cleanly.

---

# 9. Install VS Code & Go extension (recommended editor)

1. Install Visual Studio Code if you don’t have it.
2. Open VS Code → Extensions → install **Go** (by the Go team).
3. Open your project folder in VS Code. The extension will prompt to install additional tools (gopls, gofmt, delve, etc.). Accept those prompts.

---

# 10. Common troubleshooting

* `go version` not found:

  * Close and reopen terminal after install, or reboot.
  * Ensure `C:\Program Files\Go\bin` is in your PATH.

* Permission errors building or installing packages:

  * Don’t run your project from `C:\Program Files`. Use your user folders.
  * Avoid spaces or special characters in project path if tools misbehave.

* If using an ARM machine and installed the wrong binary:

  * Uninstall and reinstall the correct `arm64` MSI.

* Proxy / firewall blocks `go get`:

  * Set `GOPROXY` or configure your network/proxy to allow access to `proxy.golang.org`.

---

# 11. Best practices & tips

* Use **Go modules** (`go.mod`) for every project — it simplifies dependency management.
* Keep your code in `%USERPROFILE%\projects` or Git repos (e.g., `C:\Users\You\projects\...`).
* Use `go fmt` or the VS Code Go extension formatting on save.
* Use `go test` to run unit tests.
* Use `go env -w GOMODCACHE=...` if you want to change module cache location.

---

# Quick checklist (copy/paste)

1. Download MSI for your architecture and run it (or `winget install Go.Go`).
2. Open new terminal and run `go version`.
3. Create project folder; `main.go` with Hello world; `go run main.go`.
4. `go mod init <module>` and `go mod tidy`.
5. Install VS Code + Go extension.

---

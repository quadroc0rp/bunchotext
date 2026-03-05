<p align="center">
  <img src="./docs/icon.svg" alt="bunchotext logo" height="200">
</p>

<p align="center">
  <a href="https://go.dev/dl/"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT"></a>
  <a href="https://github.com/quadrocorp/bunchotext/graphs/contributors"><img src="https://img.shields.io/github/contributors/quadrocorp/bunchotext?color=success&logo=github" alt="Contributors"></a>
  <a href="#-usage"><img src="https://img.shields.io/badge/Type-CLI-blue?logo=terminal" alt="Type: CLI"></a>
  <a href="https://goreportcard.com/report/github.com/quadrocorp/bunchotext"><img src="https://goreportcard.com/badge/github.com/quadrocorp/bunchotext" alt="Go Report Card"></a>
</p>

---

**bunchotext** is a lightweight, high-performance CLI tool written in Go designed to concatenate source code files into a single text document. 

It is specifically optimized for preparing codebases for **Large Language Models (LLMs)**. By bundling your project files with clear separators, you can easily paste entire modules or directories into AI chat contexts for refactoring, debugging, or analysis without losing file structure context.

---

## ✨ Features

- **🚀 Fast & Lightweight:** Built in Go for quick traversal of large directories.
- **🎯 Smart Filtering:** Supports predefined file type presets (Go, Python, TypeScript, JavaScript).
- **🚫 Noise Reduction:** Automatically ignores common non-essential directories (`.git`, `node_modules`, `vendor`, etc.).
- **📝 Clear Formatting:** Each file is separated by a header indicating the file path, making it easy for humans and AI to distinguish between files.
- **🔧 Extensible:** Easy to add new file extensions or modify ignore rules in the source code.

---

## 📦 Installation

### Option 1: Pre-built Binaries (Recommended)

Download the latest pre-compiled binary for your platform from the [Releases page](https://github.com/quadrocorp/bunchotext/releases  ).

#### **Linux (amd64)**
```bash
wget https://github.com/quadrocorp/bunchotext/releases/latest/download/bunchotext-linux-amd64  
chmod +x bunchotext-linux-amd64
sudo mv bunchotext-linux-amd64 /usr/local/bin/bunchotext
```

#### **macOS**
```bash
# Intel Macs
wget https://github.com/quadrocorp/bunchotext/releases/latest/download/bunchotext-darwin-amd64  
chmod +x bunchotext-darwin-amd64
sudo mv bunchotext-darwin-amd64 /usr/local/bin/bunchotext

# Apple Silicon (M1/M2/M3)
wget https://github.com/quadrocorp/bunchotext/releases/latest/download/bunchotext-darwin-arm64  
chmod +x bunchotext-darwin-arm64
sudo mv bunchotext-darwin-arm64 /usr/local/bin/bunchotext
```

#### **Windows**

1. Download `bunchotext-windows-amd64.exe` from the [Releases page](https://github.com/quadrocorp/bunchotext/releases  )
2. Rename to `bunchotext.exe` (optional)
3. Add the folder to your system `PATH`, or run directly from the download location

---

### Option 2: Build from Source (Requires Go)

**Prerequisites:**
- [Go](https://go.dev/dl/  ) (version 1.21 or later)
- [Make](https://www.gnu.org/software/make/  ) (optional, but recommended)

#### Using Makefile (Recommended)
```bash
# Clone the repository
git clone https://github.com/quadrocorp/bunchotext.git  
cd bunchotext

# Build for your current platform
make build

# Install to ~/.local/bin
make install
```

#### Manual Build (Without Make)
```bash
# Clone the repository
git clone https://github.com/quadrocorp/bunchotext.git  
cd bunchotext

# Build for current platform
go build -o bunchotext ./cmd/bunchotext

# Move to PATH (Linux/macOS)
sudo mv bunchotext /usr/local/bin/

# Or run directly from project directory
./bunchotext
```

#### Cross-Platform Builds
```bash
# Build for all platforms (creates release artifacts)
make release

# Build for specific platform
make build-linux
make build-mac  
make build-windows
```

---

### Option 3: Go Install

```bash
go install github.com/quadrocorp/bunchotext/cmd/bunchotext@latest
```

*Note: This requires the module to be published and accessible via Go proxy.*

---
## 🚀 Usage

### Command Structure

```bash
bunchotext [command] [flags]
```

### Available Commands

| Command  | Description                                                     | Example                      |
| -------- | --------------------------------------------------------------- | ---------------------------- |
| _(none)_ | **Basic mode**: Bundle files of a specific type (requires `-t`) | `bunchotext -t go`           |
| `auto`   | **Auto mode**: Detect dominant file type and bundle             | `bunchotext auto`            |
| `all`    | **All mode**: Bundle all text files with optional filters       | `bunchotext all --gitignore` |
| `help`   | Show help for any command                                       | `bunchotext help auto`       |

## Global Flags (work with all commands)

| Flag       | Short | Default      | Description                              |
| ---------- | ----- | ------------ | ---------------------------------------- |
| `--dir`    | `-d`  | `.`          | Root directory to scan for files         |
| `--output` | `-o`  | `output.txt` | Destination file path for bundled output |

### Mode-Specific Flags

#### Basic Mode (root command)

| Flag     | Short | Required | Description                                      |
| -------- | ----- | -------- | ------------------------------------------------ |
| `--type` | `-t`  | ✅ Yes    | File type preset: `go`, `py`, `js`, `ts`, `json` |

#### Auto Mode (`auto` subcommand)

_No additional flags required_ — automatically detects the most common file type.

#### All Mode (`all` subcommand)

| Flag                | Default | Description                                                 |
| ------------------- | ------- | ----------------------------------------------------------- |
| `--gitignore`       | `false` | Respect `.gitignore` rules from the root directory          |
| `--use-ignore-dirs` | `true`  | Exclude standard directories (`.git`, `node_modules`, etc.) |
| `--no-ignore-dirs`  | `false` | Include ALL directories, overriding `--use-ignore-dirs`     |

### Quick Examples

```bash
# 🔹 Basic: Bundle Go files from current directory
./bunchotext -t go -o go_context.txt

# 🔹 Basic: Bundle Python files from specific folder
./bunchotext -d ./src -t py -o python_bundle.txt

# 🔹 Auto: Detect and bundle the dominant file type
./bunchotext auto -o detected.txt

# 🔹 All: Bundle everything, respecting .gitignore
./bunchotext all --gitignore -o full_backup.txt

# 🔹 All: Bundle everything including node_modules and .git
./bunchotext all --no-ignore-dirs -o everything.txt

# 🔹 Help: View available commands and flags
./bunchotext --help
./bunchotext auto --help
./bunchotext all --help
```

### Output Format

All modes use consistent, clean formatting with clear file separators:

```text
=========================
# internal/core/config.go
=========================

package core
...
============================
# internal/core/processor.go
============================

package core
...
```

---

## 🔤 Supported File Types

Out of the box, **bunchotext** supports the following presets:

| Preset | Extensions                                       |
| ------ | ------------------------------------------------ |
| `go`   | `.go`, `.mod`, `.sum`                            |
| `py`   | `.py`, `.pyw`, `.ipynb`                          |
| `js`   | `.js`, `.jsx`, `.mjs`, `.cjs`, `.json`, `.jsonc` |
| `ts`   | `.ts`, `.tsx`, `.d.ts`, `.json`, `.jsonc`        |
| `json` | `.json`, `.jsonc`                                |

> 💡 Want to add support for another language? See [Contributing](#contributing).

---

## 🛠️ Development

### Project Structure

```
bunchotext/
├── cmd/
│   └── bunchotext/
│       └── main.go          # CLI entry point
├── internal/
│   ├── cmd/                 # Cobra
    │   ├── all.go
    │   ├── root.go
    │   └── auto.go
│   └── core/
│       ├── config.go        # File patterns & ignore rules
│       └── processor.go     # Core processing logic
├── go.mod                   # Go module definition
├── Makefile                 # Build automation
├── LICENSE                  # MIT License
└── README.md                # This file
```

### Makefile Targets

```bash
make build          # Build for current platform
make build-linux    # Build for Linux (amd64)
make build-mac      # Build for macOS (amd64 + arm64)
make build-windows  # Build for Windows (amd64)
make release        # Build all platforms for release
make install        # Build and install to ~/.local/bin
make uninstall      # Remove installed binary
make clean          # Remove build artifacts
make help           # Show available targets
```

### Running Tests

```bash
go test ./...
```

### Creating a Release

```bash
# 1. Clean previous builds
make clean

# 2. Build all platforms
make release

# 3. Verify outputs
ls -lh bunchotext-*

# 4. Create GitHub release:
#    - Go to Releases → Draft a new release
#    - Tag version (e.g., v1.0.0)
#    - Upload binaries:
#      • bunchotext-linux-amd64
#      • bunchotext-darwin-amd64
#      • bunchotext-darwin-arm64  
#      • bunchotext-windows-amd64.exe
#    - Write changelog and publish
```

---
## 🗺️ Roadmap

| Feature                        | Status      | Version | Description                                                   |
| ------------------------------ | ----------- | ------- | ------------------------------------------------------------- |
| **Core Foundation**            |             |         |                                                               |
| Basic mode with type filtering | ✅ Completed | v1.0.0  | Bundle files by preset type (`go`, `py`, `js`, `ts`)          |
| Smart directory ignoring       | ✅ Completed | v1.0.0  | Skip `.git`, `node_modules`, `vendor`, etc. by default        |
| Clear output formatting        | ✅ Completed | v1.0.0  | File headers with separators for easy parsing                 |
| Multi-platform builds          | ✅ Completed | v1.0.0  | Pre-built binaries for Linux, macOS, Windows                  |
| **CLI Enhancements**           |             |         |                                                               |
| Cobra framework migration      | ✅ Completed | v1.1.0  | Refactored to industry-standard CLI framework                 |
| Auto mode (type detection)     | ✅ Completed | v1.1.0  | Automatically detect and bundle dominant file type            |
| All mode (unfiltered bundle)   | ✅ Completed | v1.1.0  | Bundle all text files regardless of extension                 |
| Gitignore support              | ✅ Completed | v1.1.0  | Optional `.gitignore` pattern matching via `--gitignore`      |
| Binary file detection          | ✅ Completed | v1.1.0  | Skip non-text files to avoid corrupt output                   |
| `--no-ignore-dirs` flag        | ✅ Completed | v1.1.0  | Override default ignores for complete backups                 |
| VS Code tasks integration      | ✅ Completed | v1.1.0  | Pre-configured build/debug tasks for contributors             |
| License compliance docs        | ✅ Completed | v1.1.0  | `THIRD-PARTY-NOTICES.md` for dependency transparency          |
| **Planned Features**           |             |         |                                                               |
| JSON configuration file        | 📋 Planned  | v1.2.0  | Define custom presets and ignore rules via `.bunchotext.json` |
| Extended language presets      | 📋 Planned  | v1.2.0  | Add `rs`, `cpp`, `rb`, `php`, `java`, `cs` support            |
| Nested `.gitignore` support    | 📋 Planned  | v1.2.0  | Respect `.gitignore` files in subdirectories                  |
| Homebrew formula               | 📋 Planned  | v1.2.0  | `brew install bunchotext` for macOS users                     |
| Chocolatey package             | 📋 Planned  | v1.2.0  | `choco install bunchotext` for Windows users                  |
| Custom header/footer templates | 📋 Planned  | v1.3.0  | User-defined formatting for bundled output                    |
| Pre-configured AI templates    | 📋 Planned  | v1.3.0  | Auto-add system prompts for Claude, GPT, Gemini, etc.         |
| Progress bar / verbose output  | 📋 Planned  | v1.3.0  | Visual feedback during large directory scans                  |
| Streaming mode for large repos | 📋 Planned  | v2.0.0  | Process files incrementally to reduce memory usage            |

**Legend**: ✅ Completed | 🔄 In Progress | 📋 Planned

> 💡 Have a feature request? [Open an issue](https://github.com/quadrocorp/bunchotext/issues)

---
<a id="contributing"></a>
## 🤝 Contributing

Contributions are welcome! Since file type associations are stored in simple maps within the code, adding support for a new language is straightforward.

### Adding a New Language Preset
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/add-rust-support`
3. Update `internal/core/config.go`:
   ```go
   var FilePatterns = map[string][]string{
       // ... existing presets ...
       "rs": {".rs", ".toml"},  // Add your preset here
   }
   ```
4. Update the README to document the new preset
5. Submit a Pull Request

### General Contribution Guidelines
- Keep changes focused and well-documented
- Update the README if adding user-facing features

### Reporting Issues
Found a bug or have a feature request? [Open an issue](https://github.com/quadrocorp/bunchotext/issues) with:
- A clear description of the problem or idea
- Steps to reproduce (for bugs)
- Your OS and Go version (if relevant)

---

## 📜 License

This project is licensed under the MIT License.

```text
MIT License

Copyright (c) 2026 Tsupko Nikita "quadrocorp" Romanovich

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

> **Made with ❤️ for developers and AI enthusiasts**  
> [Report an issue](https://github.com/quadrocorp/bunchotext/issues  ) • [View source](https://github.com/quadrocorp/bunchotext  ) • [Releases](https://github.com/quadrocorp/bunchotext/releases  )

# harutmonitor

## 📦 Installation

### 🐧 Linux Installation

#### 1. Clone the repository

```bash
git clone https://github.com/BadalyanHarutyun/harutmonitor.git
cd harutmonitor
```

#### 2. Build the binary

```bash
make build
```

#### 3. Move the binary to a system path

```bash
sudo mv harutmonitor /usr/local/bin/
```

#### 4. Verify installation

```bash
harutmonitor --help
```

---

### 🪟 Windows Installation

#### 1. Clone the repository

```bash
git clone https://github.com/BadalyanHarutyun/harutmonitor.git
cd harutmonitor
```

#### 2. Build the binary

```bash
make build
```

#### 3. Move the executable to system path

```bash
move harutmonitor.exe C:\Windows\System32\
```

#### 4. Verify installation

```bash
harutmonitor.exe --help
```

---

## ⚙️ Makefile Commands

| Command      | Description                                     |
| ------------ | ----------------------------------------------- |
| `make build` | Build the Go binary (`harutmonitor`)            |
| `make run`   | Run the application directly (`go run main.go`) |
| `make clean` | Remove the compiled binary                      |
| `make tidy`  | Format code and tidy dependencies               |

---

## 🚀 Usage

After installing `harutmonitor`, you can use it like any other command-line tool.

### Example

```bash
harutmonitor sleep 5 && echo Hello
```

Or with a log file:

```bash
harutmonitor --monitor-log=mylogs.log sleep 5 && echo Hello
```

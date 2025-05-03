# harutmonitor
## 📦 Installation

### Option 1: Using `make` (recommended)
Make sure you have `make` and `wget` installed. Then simply run:
you can move you binary where you want /usr/local/bin/ or /usr/bin/ or similar place
```bash
make install
```
```bash
wget -O harutmonitor https://raw.githubusercontent.com/BadalyanHarutyun/harutmonitor/main/binaries/linux/harutmonitor
chmod +x harutmonitor
sudo mv harutmonitor /usr/local/bin/
```

## Windows dowload and installation
```bash
curl -LO https://raw.githubusercontent.com/BadalyanHarutyun/harutmonitor/main/binaries/windows/harutmonitor.exe
move harutmonitor.exe C:\Windows\System32\
```
## Usage

After installing `harutmonitor`, you can use it like any other command-line tool.

### Example

```bash
harutmonitor sleep 5 && echo Harut
harutmonitor --monitor-log=mylogs.log sleep 5 && echo Harut

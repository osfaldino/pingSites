# This project is designed for parallel pinging of selected servers
## Installation & Configuration

```bash
go build -o pingSites main.go # compile the binary
sudo cp pingSites /usr/local/bin/ # move the program to bin
sudo cp sites.conf /usr/local/bin/ # move the config to bin
```

###How to edit the configuration file?
```bash
sudo nano /usr/local/bin/sites.conf # edit the config file
```

### Each URL should be on a new line.


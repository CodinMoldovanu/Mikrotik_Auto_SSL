# Mikrotik_Auto_SSL

This is a simple Go application to create and validate SSL Certificates for any domain using LetsEncrypt/Certbot, and works if you own/run a MikroTik device for your home routing needs.

*You need to have `certbot` installed and in your PATH for this to work.*

### Usage
1. Pull this repo
2. `go build`
3. `./mikrotik_auto_ssl` - and follow the prompts. 


### To-Do
???PROFIT

You can hold your sensible information in an `.env` file in the same place as the binary and configure your IP, PORT (for the API), USERNAME, PASSWORD in there:
```
IP=192.168.1.1
PORT=8728
USERNAME=SUPERMAN
PASSWORD=LOISISQT
```


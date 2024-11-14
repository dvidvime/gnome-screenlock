# gnome-screenlock 
Simple tool to control gnome screensaver via DBus using HTTP requests.
May be useful for example to manage remote kiosk displays.

## Build
* Clone the repository and `cd` into it's directory
* Run `go build .`
* Change permissions if needed `chmod +x ./screenlock`
* Run in background `./screenlock &`

## Usage
http:// _your-ip_ :57650/_command_

### Available commands:

**/on** - enables screensaver

**/off** - disables screensaver

**/status** - returns screensaver status (true if screensaver is running now)

## Running as a service with automatic restart
Put this into `~/.config/systemd/user/screenlock.service`:

```[Unit]  
Description=Screen control Service  
After=graphical.target  

[Service]  
ExecStart=_path to folder with binary_/screenlock
Restart=always
RestartSec=5  

[Install]  
WantedBy=default.target
```
Then reload services and enable `systemctl --user daemon-reload && systemctl --user enable screenlock.service`

Check status `systemctl --user status screenlock.service`
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
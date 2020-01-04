# OctoPrint root actions

Suid-able binary to provide OctoPrint with limited root functionality.

Actions are hardcoded:

- Reboot
- Restart `octoprint.service`
- Restart `klippy.service` (http://klipper3d.org/)
- Restart `ustreamer.service` (https://github.com/pikvm/ustreamer)

Made it into a binary so it can be made setuid and run as root easily
without adding painful rules to `/etc/sudoers.d`

## Usage

```bash
$ go build
$ sudo chown root:root octoprint_root_actions
$ sudo chmod +s octoprint_root_actions
```

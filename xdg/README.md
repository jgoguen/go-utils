# go-utils/xdg

This provides a library of XDG functions suitable for use on Windows, macOS,
Linux, or BSD. The XDG base directory specification explicitly only applies to
Linux; this library assumes BSD systems are also capable of following the XDG
specification and attempts to choose sensible analogs on Windows and macOS.

Many other XDG libraries either only support Linux or attempt to directly
impose the Linux locations on macOS and/or Linux. This library attempts instead
to extrapolate the Linux XDG specifications to other platforms by using the
corresponding platform-specific directory (or the closest match). For example,
rather than defining the default value of `XDG_CONFIG_HOME` to be
`$HOME/.config` on all platforms the following defaults are chosen:

- Windows: `%USERPROFILE%\AppData\Roaming`
  (e.g. `C:\Users\jgoguen\AppData\Roaming`)
- macOS: `$HOME/Library/Preferences` (e.g. `/Users/jgoguen/Library/Preferences`)
- Other (Linux, BSD, etc.): `$HOME/.config` (e.g. `/home/jgoguen/.config`)

This library is based on the [XDG Base Directory Specification version 0.7](https://specifications.freedesktop.org/basedir-spec/basedir-spec-0.7.html)

## XDG Base Directory Specification

The XDG Base Directory Specification defines a standard location for storing
user-specific configuration and data files, as well as system-wide configuration
and data file locations. Only Linux is expected to follow this standard, but
corresponding directories exist on both Windows and macOS and other non-Linux
UNIX systems (e.g. BSD) can easily follow the XDG specification.

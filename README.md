# Backup Tool

A simple Go package for creating periodic backups of a specified folder.

## Usage

To use this package, compile the code and run the resulting executable with the following flags:

* `-f` or `--folder`: the path to the folder to be backed up
* `-o` or `--out-folder`: the path to the folder where backups will be stored
* `-t` or `--delta`: the time interval (in minutes) between backups (default: 2 minutes)

Example:
```bash
go build main.go
./main -f /path/to/folder -o /path/to/backups -t 30
```
This will create a backup of the specified folder every 30 minutes, storing the backups in the specified output folder.

## How it works

1. The program checks for the existence of a backup file with the current timestamp. If it exists, the program exits.
2. The program creates a new backup file with the current timestamp using the `tar` command.
3. The backup file is stored in the specified output folder.

## Notes

* The program uses the `tar` command to create compressed backups.
* The program uses the `time` package to calculate the current timestamp and determine the backup interval.
* The program uses the `flag` package to parse command-line flags.
* The program uses the `log` package to handle errors.

## Requirements

* Go 1.14 or later
* `tar` command installed on the system

## License

This package is licensed under the MIT License. See LICENSE for details.

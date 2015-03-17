# Gupper: Go Uptime Monitor
The goal of Gupper is to provide a non-persistent status monitoring api for situations where keeping uptime history is not required. Gupper aims to be simple to configure and run, with few requirements and a straightforward config that provides sane defaults.

The use case Gupper aims to satisfy is in non-commercial use cases. For those who run home services and would like a monitor, but don't have customers who need an entire status history.

## Program Structure
The main Gupper binary simply runs an array of commands for each system. It records the results and reports the status on the /status endpoint.

I have done it this way in hopes of making Gupper more extensible - it can handle the status of any system as long as there is a command that will exit with an appropriate status code or can print json to stdout.

## Configuration
The config file currently consists of an array of "systems". Each system can have multiple commands to monitor its status.

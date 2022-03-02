---
sidebar_position: 1
---

# Installing Veles

:::info Software-Recommendations

Whilst Veles should work on Linux, Windows and macOS, **we strongly recommend you use Linux** as this is what
Veles was developed and tested on! (Veles will run on arm-based minicomputers)

:::

## Docker

:::tip TODO

This section will come soon.

:::

## Bare-Metal

:::info Pre-Flight-Installations

Veles uses *MongoDB* as a database backend. Please [install MongoDB Community Server](https://www.mongodb.com/try/download/community) first.

:::

### Using the binary release

Veles provides ready-made binaries for the major OSes and architectures.

 1. Go to the [latest release on GitHub](https://github.com/Unkn0wnCat/matrix-veles/releases/latest)
 2. Navigate down to "*Assets*" and find the correct file for your OS and architecture
 3. Download the file (Linux/macOS: .tar.gz, Windows: .zip)
    1. (Optional) Check the md5 sum of your downloaded file against the provided md5 sum
 4. Unpack the file (Your OS should come with utilities to do this)
 5. Navigate to the unpacked directory in your Terminal
 6. Run `./matrix-veles generateConfig` to generate a basic config<br/>(Linux: You may need to allow execution of the file using `chmod +x ./matrix-veles`)
 7. Edit the configuration in `./config.yaml` to reflect your setup
 8. Start Matrix-Veles using `./matrix-veles run`

You now have a fully functioning install of Veles! 🎉 Access the web interface at http://127.0.0.1:8123!

### Building from Source

:::info

Experience with GoLang is beneficial for this!

:::

To build from source make sure you have the [latest version of GoLang](https://go.dev/dl/) installed.

1. Open a terminal and execute `go install github.com/Unkn0wnCat/matrix-veles@latest`
2. After a few minutes the build should be complete
3. Run `matrix-veles generateConfig` in the directory you want your configuration to reside in
4. Edit the configuration in `./config.yaml` to reflect your setup
5. Start Matrix-Veles using `matrix-veles run` in the same directory as your config

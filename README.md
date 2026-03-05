# Home Environment Monitoring System

![MIT License](https://img.shields.io/badge/License-MIT-green.svg) ![Workflow Status](https://github.com/Leomotors/home-env/actions/workflows/release.yml/badge.svg)

This repository contains a home environment monitoring system that utilizes ESP32 and AHT20 sensors to collect temperature and humidity data. The gathered information is accessible via a web interface and can be integrated with Prometheus for advanced monitoring.

## Features

- Real-time temperature and humidity monitoring
- Web interface for ~~my friend to spy me~~ easy data access
- Prometheus integration for advanced analytics

## Acknowledgments

- Special thanks to [WasinUddy](https://github.com/WasinUddy/Homelab-Environments-Monitor) for hardware recommendations and inspiration (basically ป้ายยา)
- Project Assistance: ChatGPT + GitHub Copilot

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Stack

### On Premise (Current)

![](./images/stackv3.webp)

(Logo from https://github.com/SAWARATSUKI/ServiceLogos)

### Previous Stack

<details>
<summary>Version 1</summary>

![](./images/stack.webp)

</details>

<details>
<summary>Version 2 Software</summary>

![](./images/stackv2vlogo.webp)

</details>

### Version 2 Hardware

![](./images/board.webp)

Note: Hardware V2 remain unchanged in Software V3, it is designed to be backward compatible with existing hardware code (I'm lazy to upload Arduino Code 💀).

### Grafana Dashboard

(Old One, to be updated when I finished building Grafana Dashboard, later)

![](./images/grafana.webp)

## Discord Alert Feature

Alert when ESP32 is not sending data

### Version 1

![](./images/discord-alert.webp)

### Version 2 Canary

![](./images/discord-alert-v2.webp)

### Version 2

![](./images/discord-alert-v2-2.webp)

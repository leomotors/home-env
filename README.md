# Home Environment Monitoring System

![MIT License](https://img.shields.io/badge/License-MIT-green.svg) ![Workflow Status](https://github.com/Leomotors/home-env/actions/workflows/release.yml/badge.svg)

This repository contains a home environment monitoring system that utilizes ESP32 and AHT20 sensors to collect temperature and humidity data. The gathered information is accessible via a web interface and can be integrated with Prometheus for advanced monitoring.

## Features

- Real-time temperature and humidity monitoring
- Web interface for ~~my friend to spy me~~ easy data access
- Prometheus integration for advanced analytics

## Acknowledgments

- Special thanks to [WasinUddy](https://github.com/WasinUddy/Homelab-Environments-Monitor) for hardware recommendations and inspiration (basically ป้ายยา)
- Project Assistance: ChatGPT

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Stack

### On Premise (Current)

![](./images/stackv2.webp)

### Previous Stack

![](./images/stack.webp)

### Version 2 Hardware

![](./images/board.webp)

### Grafana Dashboard

![](./images/grafana.webp)

## Discord Alert Feature

Alert when ESP32 is not sending data

### Version 1

![](./images/discord-alert.webp)

### Version 2 Canary

![](./images/discord-alert-v2.webp)

### Version 2

![](./images/discord-alert-v2-2.webp)

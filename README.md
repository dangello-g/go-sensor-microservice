# Go Sensor Microservice

A lightweight Go microservice that simulates humidity and light intensity data and streams it via WebSockets.

## Overview

This microservice generates real-time simulated IoT sensor data and streams it via WebSockets. It provides continuous data streams that can be consumed by frontend applications, data processing pipelines, or other services. While it currently runs as a single-instance service for small-scale IoT simulations and prototyping, it can be extended with scalability and fault tolerance features to support larger deployments.

## Features

- Simulated Sensor Data: Generates real-time humidity and light intensity data for testing and development purposes.
- WebSocket Streaming: Streams sensor data to connected clients via WebSockets, enabling real-time data consumption.

## Planned Features

- InfluxDB Integration: Stores simulated sensor data in InfluxDB, a high-performance time series database, facilitating efficient data storage and retrieval.
- Kafka Integration: Implement a Kafka producer to publish sensor data, enabling event-driven architectures and decoupled data pipelines.
- Improved Fault Tolerance: Implement automatic reconnection, data buffering, and retry mechanisms for a more resilient system.
- Modular Architecture: Refactor the codebase for easy extension, allowing additional sensor types and processing modules.
- Threshold-Based Actions: Implement a WebSocket-based mechanism where the frontend sends messages when sensor readings exceed predefined thresholds. The backend will manage these events and trigger appropriate actions, such as activating devices in response to humidity or light intensity levels.

## Technology Stack

- Go 1.24
- Gorilla WebSocket package

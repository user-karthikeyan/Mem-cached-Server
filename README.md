# Custom Mem cached Server in Go
Memcached a in-memory key-value store used to reduce latency in back-end systems by storing chunks of data.


---
### Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Supported Commands](#supported-commands)
6. [Architecture and Design](#architecture-and-design)
7. [Examples](#examples)

### Introduction
This project is a custom Memcached server implemented in Go, replicating essential caching functionalities for fast, reliable data storage and retrieval. Designed with simplicity and concurrency control in mind, it supports core Memcached commands along with several advanced operations.

### Features
- **Core Commands**: Includes `set`, `get`, and `delete` commands for standard caching functionality.
- **Advanced Commands**: Supports `cas`, `gets`, `incr`, `decr`, `append`, `prepend`, `replace`, and `add` for more granular control over cache data.
- **Check-And-Set (CAS) Support**: Allows conditional updates to cache entries.
- **Expiration Support**: Items can be set to expire automatically after a specified time.
- **Concurrency Control**: Thread-safe operations are enabled with read-write locking mechanisms.

### Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/user-karthikeyan/Mem-cached-Server
   ```
2. Navigate into the project directory:
   ```bash
   cd MEM-CACHED-SERVER
   ```
3. Build the project:
   ```bash
   go build .
   ```

### Usage
1. Start the server:
   ```bash
   ./memcached-server
   ```
2. Use any Memcached client or a custom script to interact with the server by sending supported commands.

### Supported Commands
- **set**: Store data under a specified key with optional flags, expiration time, and byte count.
- **cas**: Store data only if the supplied CAS value matches the current CAS value of the existing item.
- **get**: Retrieve the data associated with the specified key.
- **gets**: Retrieve data and its CAS value for the specified key.
- **delete**: Delete the data associated with the specified key.
- **incr**: Increment the value of an existing numeric item by a specified amount.
- **decr**: Decrement the value of an existing numeric item by a specified amount.
- **append**: Append data to the end of an existing item.
- **prepend**: Prepend data to the beginning of an existing item.
- **replace**: Replace the data for an existing item if the key already exists.
- **add**: Add data under a specified key only if the key does not already exist.

### Architecture and Design
The server is structured with modular, concurrent components:
- **Request Handling**: Accepts incoming client connections and processes commands.
- **Command Handlers**: Each Memcached command is managed by dedicated handlers to keep the code modular and maintainable.
- **Concurrency Control**: Uses mutexes to ensure that all operations on shared data are thread-safe.
- **In-Memory Storage**: Implements an efficient, in-memory store for fast access to cached data.

### Examples
1. **Set a Value**
   ```bash
   set key1 0 900 5
   value
   ```
2. **Get a Value**
   ```bash
   get key1
   ```
3. **Delete a Value**
   ```bash
   delete key1
   ```
4. **Conditional Set with CAS**
   ```bash
   cas key1 0 900 5 12345
   new_value
   ```


# Real-Time Chat Application: Tech Stack & Features

This document provides a comprehensive overview of the technology stack and features implemented in the Real-Time Chat application.

---

## 🛠️ Technology Stack

The application uses a modern, lightweight, and highly performant stack to handle real-time, bi-directional communication:

### 1. Backend (Go)
* **Language:** Go (1.20+)
* **Networking & WebSockets:** `github.com/gorilla/websocket`
  * Upgrades standard HTTP connections to persistent TCP sockets.
  * Handles concurrent I/O operations through separate Goroutines.
* **Concurrency Model:** Goroutines and Channels
  * Uses a centralized `Hub` struct serving as an orchestrator/broker.
  * Handled via a thread-safe `select` loop to process client registrations, unregistrations, and broadcasts without mutex bottlenecks.
* **Data Format:** JSON (`encoding/json`)

### 2. Frontend (React + TypeScript)
* **Library:** React 18+
* **Language:** TypeScript (for type safety and strict message contracts)
* **Build Tool:** Vite (high-performance bundler)
* **State Management:** React Hooks (`useState`, `useEffect`, `useRef` for tracking the raw socket connection)

### 3. Protocol
* **WebSocket ($ws://$)**: A full-duplex, persistent connection between the client and the server, bypassing the overhead of traditional HTTP polling.

---

## 🚀 Key Features

### 📡 1. Real-Time Bidirectional Messaging
* Instant message delivery with negligible latency.
* Broadcast pipeline: Any message sent by one user is immediately relayed to all other active connections.

### 👤 2. Dynamic Identity (Nickname Setup)
* Usernames are not hardcoded. On page load, a clean join screen prompts users for a nickname.
* The nickname is passed securely to the backend on connection handshake via URL query parameters (`ws://.../ws?username=Alice`).

### 🕒 3. Server-Generated Trusted Timestamps
* Timestamps are generated securely on the Go server (`time.RFC3339` standard) when the message is processed. This prevents clients from spoofing message order or times.

### 🎯 4. Smart UI Features
* **Auto-Scroll:** The chat container automatically scrolls to the bottom when new messages arrive.
* **System Type vs. Chat Type:** Message schemas differentiate between actual user chats and system logs.

### 🧹 5. Lifecycle & Connection Management
* Clean socket disposal: The React client closes the socket on component unmount to prevent resource leaks.
* Clean Go tear-down: The Go hub automatically deregisters clients when they close their browser tab, terminating their respective reading and writing Goroutines.

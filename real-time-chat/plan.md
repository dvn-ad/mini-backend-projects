# Plan: Real-Time Chat Application

## 1. Overview

A bidirectional, real-time messaging platform. The system uses a centralized Go server to orchestrate connections and a React (TS) frontend to provide a responsive user interface.

## 2. Tech Stack

* **Backend:** Golang (1.20+)
* **Networking:** `gorilla/websocket` (Socket handling)
* **ID Generation:** `google/uuid` (Unique client identification)


* **Frontend:** React 18+ (TypeScript)
* **State Management:** React Hooks (`useState`, `useEffect`, `useRef`)
* **Build Tool:** Vite (Optimized for TS development)


* **Protocol:** WebSockets ($ws://$)

---

## 3. Architecture & Logic

### A. The Go Backend (Orchestrator)

The backend acts as a "Hub" or "Manager" that maintains a registry of active memory addresses (Clients).

1. **The Client Manager (Hub):**
* Maintain a map of active connections.
* Use **Channels** to handle registration, unregistration, and incoming broadcast messages.
* Implement a `select` loop to listen for channel activity.


2. **The Client Instance:**
* Each connection needs its own `Read` and `Write` goroutines.
* **Read Pump:** Continuously listens for messages coming *from* the browser.
* **Write Pump:** Continuously sends messages *to* the browser from the hub's broadcast channel.


3. **The WebSocket Upgrader:**
* Convert standard HTTP requests into Persistent WebSocket connections.
* Handle Cross-Origin Resource Sharing (CORS) to allow the React dev server to connect.



### B. The React Frontend (Client)

The frontend manages the lifecycle of the socket connection within the React component tree.

1. **Custom Hook / Service:**
* Encapsulate the `WebSocket` browser API logic.
* Manage connection states: `Connecting`, `Open`, `Closed`.
* Parse incoming JSON messages into TypeScript interfaces.


2. **The UI Layer:**
* **Message List:** A scrollable container that renders an array of message objects.
* **Input Area:** A controlled form to capture user text and trigger the socket `send` method.
* **Auto-Scroll:** Logic to keep the view at the bottom when new messages arrive.


3. **TypeScript Interfaces:**
* Define a strict `Message` interface (Sender, Content, Timestamp) to ensure consistency between the Go JSON tags and the React props.



---

## 4. Development Roadmap

### Phase 1: Go Foundation

* Define `Client` and `Hub` structs.
* Implement the `Hub` run loop (handling `register`, `unregister`, and `broadcast`).
* Setup the WebSocket upgrader and the `/ws` endpoint.

### Phase 2: React Setup

* Initialize Vite + React + TS.
* Create a `useSocket` custom hook or a Context Provider to manage the `WebSocket` instance.
* Implement basic event listeners (`onmessage`, `onopen`, `onclose`).

### Phase 3: Message Flow

* Connect the `WritePump` in Go to send a "Welcome" message.
* Connect the React input to the `socket.send()` method.
* Test broadcasting: Open two browser tabs and verify messages appear in both.

### Phase 4: UI/UX Refinement

* Differentiate between "System" messages (e.g., "User Joined") and "User" messages.
* Add timestamps and "Sent by me" styling.
* Handle reconnection logic if the server drops.

---

## 5. Key Considerations

* **Concurrency:** Ensure the Go Hub uses mutexes or channels properly to avoid race conditions when multiple users join/leave simultaneously.
* **JSON Consistency:** Ensure the `json:"content"` tags in Go match the property names in your TypeScript interfaces.
* **Resource Cleanup:** Ensure the React `useEffect` cleanup function closes the socket connection to prevent memory leaks in the browser.

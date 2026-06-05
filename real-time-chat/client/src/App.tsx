import React, { useState, useEffect, useRef } from 'react';
import './App.css';

// Move interfaces outside the component
interface ChatMessage {
  type: 'chat' | 'system';
  sender: string;
  content: string;
  timestamp: string;
}

function App() {
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [input, setInput] = useState("");
  const socketRef = useRef<WebSocket | null>(null);
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const [username, setUsername]=useState("");
  const [tempUsername, setTempUsername]=useState("")
  const [isJoined,setIsJoined]=useState(false)

  useEffect(() => {
    if (chatContainerRef.current) {
      chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
    }
  }, [messages]);

  useEffect(() => {
    if(!isJoined||!username)return
    const socket = new WebSocket(`ws://localhost:8080/ws?username=${encodeURIComponent(username)}`);

    socketRef.current = socket;

    socket.onmessage = (event) => {
      try{
        const parsedMessage: ChatMessage=JSON.parse(event.data);
        setMessages((prev)=>[...prev,parsedMessage]);
      }catch(error){
        console.error("Error parsing message:", error);
      }
    };

    socket.onopen = () => console.log("Connected to Go Server");
    socket.onclose = () => console.log("Disconnected from Go Server");

    return () => {
      socket.close();
    };
  }, [isJoined,username]);

  const sendMessage = () => {
    if (input.trim() !== "" && socketRef.current?.readyState === WebSocket.OPEN) {
      socketRef.current.send(input);
      setInput(""); 
    }
  };
  const handleJoin = () => {
    if (tempUsername.trim() !== "") {
      setUsername(tempUsername.trim());
      setIsJoined(true);
    }
  };

  if (!isJoined) {
    return (
      <div className="chat-app-container">
        <div className="join-card">
          <h2 className="join-title">Enter your nickname to join</h2>
          <div className="join-form">
            <input
              className="input-styled"
              value={tempUsername}
              onChange={(e) => setTempUsername(e.target.value)}
              onKeyDown={(e) => e.key === 'Enter' && handleJoin()}
              placeholder="Nickname..."
            />
            <button className="button-styled" onClick={handleJoin}>Join</button>
          </div>
        </div>
      </div>
    );
  }
  return (
    <div className="chat-app-container">
      <div className="chat-box">
        <div className="chat-header">
          <h2>Real-Time Go Chat</h2>
          <span className="user-badge">{username}</span>
        </div>
        <div ref={chatContainerRef} className="chat-messages">
          {messages.map((msg, i) => {
            const isSystem = msg.type === 'system';
            const isSelf = msg.sender === username;
            
            return (
              <div 
                key={i} 
                className={`message-row ${isSystem ? 'system' : isSelf ? 'self' : 'other'}`}
              >
                {!isSystem && !isSelf && (
                  <span className="message-sender">{msg.sender}</span>
                )}
                <div className="message-bubble">
                  <span>{msg.content}</span>
                  {!isSystem && (
                    <div className="message-meta">
                      {new Date(msg.timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                    </div>
                  )}
                </div>
              </div>
            );
          })}
        </div>
        <div className="chat-input-area">
          <input 
            className="input-styled"
            value={input} 
            onChange={(e) => setInput(e.target.value)} 
            onKeyDown={(e) => e.key === 'Enter' && sendMessage()}
            placeholder="Type a message..."
          />
          <button className="button-styled" onClick={sendMessage}>Send</button>
        </div>
      </div>
    </div>
  );
} // <--- This closing bracket is important!

export default App;
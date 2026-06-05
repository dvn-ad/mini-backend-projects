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

  useEffect(() => {
    if (chatContainerRef.current) {
      chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
    }
  }, [messages]);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws");
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
  }, []);

  const sendMessage = () => {
    if (input.trim() !== "" && socketRef.current?.readyState === WebSocket.OPEN) {
      socketRef.current.send(input);
      setInput(""); 
    }
  };

  return (
    <div style={{ padding: '20px', maxWidth: '600px', margin: '0 auto' }}>
      <h2>Real-Time Go Chat</h2>
      <div ref={chatContainerRef}
      style={{ 
        border: '1px solid #ccc', 
        height: '300px', 
        overflowY: 'scroll', 
        marginBottom: '10px',
        background: '#f9f9f9',
        padding: '10px'
      }}>
        {messages.map((msg, i) => (
          <div key={i} style={{ 
            padding: '8px', 
            marginBottom: '5px', 
            background: '#fff', 
            borderRadius: '4px',
            boxShadow: '0 1px 2px rgba(0,0,0,0.1)'
          }}>
            <strong>{msg.sender}: </strong>
            <span>{msg.content}</span>
            <div style={{fontSize:'10px',color:'#888',marginTop:'4px'}}>
              {new Date(msg.timestamp).toLocaleTimeString()}
            </div>
          </div>
        ))}
      </div>
      <div style={{ display: 'flex', gap: '10px' }}>
        <input 
          style={{ flex: 1, padding: '10px' }}
          value={input} 
          onChange={(e) => setInput(e.target.value)} 
          onKeyDown={(e) => e.key === 'Enter' && sendMessage()}
          placeholder="Type a message..."
        />
        <button onClick={sendMessage} style={{ padding: '10px 20px' }}>Send</button>
      </div>
    </div>
  );
} // <--- This closing bracket is important!

export default App;
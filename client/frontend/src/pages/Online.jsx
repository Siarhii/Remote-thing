import React, { useState } from "react";
import { ConnectToServer } from "../../wailsjs/go/main/App";
import "./Online.css";

const DeviceConnect = () => {
  const [deviceCode, setDeviceCode] = useState("");
  const [connectionStatus, setConnectionStatus] = useState(null);

  const handleConnect = async (dc) => {
    try {
      const result = await ConnectToServer(dc);

      if (result.includes("Successfully connected")) {
        console.log("Connection Successful: ", result);
        setConnectionStatus("success");
      } else {
        console.error("Error connecting device:", result);
        setConnectionStatus("failure");
      }
    } catch (error) {
      console.error("Unexpected error:", error);
      setConnectionStatus("failure");
    }
  };
  const handleKeyPress = (event) => {
    if (event.key === "Enter") {
      handleConnect(deviceCode);
    }
  };

  return (
    <div className="online-container">
      <div className="online-card">
        <h1>Connect Device</h1>
        <p>Enter the unique code to link your device</p>

        <div className="device-code-section">
          <input
            type="text"
            value={deviceCode}
            onChange={(e) => setDeviceCode(e.target.value)}
            onKeyPress={handleKeyPress}
            placeholder="Enter device code"
            className="device-code"
          />
          <button
            onClick={() => {
              handleConnect(deviceCode);
            }}
            className="connect-button"
          >
            Connect
          </button>
        </div>

        {connectionStatus && (
          <div className={`connection-message ${connectionStatus}`}>
            {connectionStatus === "success"
              ? "Device successfully connected"
              : "Device failed to connect"}
          </div>
        )}

        <div className="connection-info">
          <div className="connection-status">
            <span
              className={`status-indicator ${
                connectionStatus === "success" ? "online" : ""
              }`}
            ></span>
            {connectionStatus === "success" ? "Connected" : "Disconnected"}
          </div>
          <p className="connection-details">
            WebSocket: {connectionStatus === "success" ? "Active" : "Inactive"}
            <br />
            Last Sync: {connectionStatus === "success" ? "Just now" : "N/A"}
          </p>
        </div>
      </div>
    </div>
  );
};

export default DeviceConnect;

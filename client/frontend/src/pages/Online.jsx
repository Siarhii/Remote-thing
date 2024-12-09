import React, { useState } from "react";
import { ConnectToServer } from "../../wailsjs/go/main/App";
import "./Online.css";

const DeviceConnect = () => {
  const [deviceCode, setDeviceCode] = useState("");
  const [connectionStatus, setConnectionStatus] = useState(null);

  const handleConnect = async () => {
    const mockConnectDevice = async (deviceCode) => {
      return new Promise((resolve, reject) => {
        // Call the Go function with the deviceCode and a callback
        ConnectToServer(deviceCode, (result) => {
          if (result.includes("Successfully connected")) {
            resolve(result); // Resolve promise with success message
          } else {
            reject(result); // Reject promise with error message
          }
        });
      });
    };

    try {
      const successMessage = await mockConnectDevice("some-device-id");
      console.log("Connection Successful: ", successMessage);
      setConnectionStatus("success"); // Update UI to show success
    } catch (error) {
      console.error("Error connecting device:", error);
      setConnectionStatus("failure"); // Update UI to show failure
    }
  };

  const handleKeyPress = (event) => {
    if (event.key === "Enter") {
      handleConnect();
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
          <button onClick={handleConnect} className="connect-button">
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

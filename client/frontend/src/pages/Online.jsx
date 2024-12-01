import React, { useState } from "react";
import "./Online.css";

const Online = () => {
  const [deviceCode, setDeviceCode] = useState("ABC123XYZ");
  const [buttonName, setButtonName] = useState("COPY");

  const copyToClipboard = () => {
    navigator.clipboard.writeText(deviceCode);
    setButtonName("Copied!");
    setTimeout(() => {
      setButtonName("Copy");
    }, 2000);
  };

  return (
    <div className="online-container">
      <div className="online-card">
        <h1>Device Connected</h1>
        <p>Copy this unique code to link your device</p>

        <div className="device-code-section">
          <code className="device-code">{deviceCode}</code>
          <button onClick={copyToClipboard} className="copy-button">
            {buttonName}
          </button>
        </div>

        <div className="connection-info">
          <div className="connection-status">
            <span className="status-indicator online"></span>
            Connected
          </div>
          <p className="connection-details">
            WebSocket: Active
            <br />
            Last Sync: Just now
          </p>
        </div>
      </div>
    </div>
  );
};

export default Online;

import React, { useState } from "react";
import "./Offline.css";
import { ActionFromClient } from "../../wailsjs/go/main/App";

const Offline = () => {
  const [timer, setTimer] = useState(0);
  const [action, setAction] = useState("");
  const [message, setMessage] = useState("");
  const [messageType, setMessageType] = useState("");
  const [countdown, setCountdown] = useState(null);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!action) {
      setMessage("Please select an action");
      setMessageType("error");
      return;
    }

    if (timer > 44000) {
      setMessage("You cannot set timer greater than a month (44000 minutes)");
      setMessageType("error");
      return;
    }

    setMessage(
      `Your device will ${
        action.charAt(0).toUpperCase() + action.slice(1)
      } in ${timer} minutes`
    );
    setMessageType("success");

    ActionFromClient(action, timer);
    setCountdown(timer * 60);
    const countdownInterval = setInterval(() => {
      setCountdown((prevCountdown) => {
        if (prevCountdown <= 1) {
          clearInterval(countdownInterval);
          return null;
        }
        return prevCountdown - 1;
      });
    }, 1000);
  };

  const formatCountdown = (totalSeconds) => {
    if (totalSeconds === null) return "";
    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;
    return `${hours.toString().padStart(2, "0")}:${minutes
      .toString()
      .padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`;
  };

  return (
    <div className="offline-container">
      <div className="offline-card">
        <h1>Device Control</h1>
        <p>Schedule an Action with a timer</p>

        <form onSubmit={handleSubmit} className="action-form">
          <div className="action-group">
            <label className="action-label">
              <input
                type="radio"
                name="action"
                value="shutdown"
                checked={action === "shutdown"}
                onChange={(e) => setAction(e.target.value)}
                className="action-radio"
              />
              <span>Shutdown</span>
            </label>

            <label className="action-label">
              <input
                type="radio"
                name="action"
                value="restart"
                checked={action === "restart"}
                onChange={(e) => setAction(e.target.value)}
                className="action-radio"
              />
              <span>Restart</span>
            </label>

            <label className="action-label">
              <input
                type="radio"
                name="action"
                value="sleep"
                checked={action === "sleep"}
                onChange={(e) => setAction(e.target.value)}
                className="action-radio"
              />
              <span>Sleep</span>
            </label>
          </div>

          <div className="timer-section">
            <label>Timer (minutes):</label>
            <input
              type="number"
              value={timer}
              onChange={(e) => setTimer(e.target.value)}
              min="0"
              className="timer-input"
            />
          </div>

          <button type="submit" className="schedule-button">
            Schedule Action
          </button>
        </form>
      </div>

      {message && (
        <div className={`message-container ${messageType}`}>
          <div>
            {message}
            {countdown !== null && (
              <span className="countdown">{formatCountdown(countdown)}</span>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

export default Offline;

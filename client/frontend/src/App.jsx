import React, { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import { Routes, Route } from "react-router-dom";
import Online from "./pages/online";
import Offline from "./pages/offline";
import { ActionFromClient } from "../wailsjs/go/main/App";
import "./App.css";

const App = () => {
  const location = useLocation();
  const [showPowerOptions, setShowPowerOptions] = useState(false);

  function togglePowerOptions() {
    setShowPowerOptions(!showPowerOptions);
  }

  return (
    <div className="app-container">
      <div className="sidebar">
        <Link to="/" className="nav-button-link">
          <button
            className={`nav-button ${
              location.pathname === "/" ? "active" : ""
            }`}
          >
            Offline
          </button>
        </Link>

        <Link to="/m" className="nav-button-link">
          <button
            className={`nav-button ${
              location.pathname === "/m" ? "active" : ""
            }`}
          >
            Online
          </button>
        </Link>

        <div className="shutdown-container">
          <button onClick={togglePowerOptions} className="shutdown-button">
            <svg
              enable-background="new 0 0 24 24"
              height="24"
              viewBox="0 0 24 24"
              width="24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="m12 24c-5.5 0-10-4.5-10-10 0-4.1 2.4-7.7 6.2-9.2.5-.2 1.1 0 1.3.5s0 1.1-.5 1.3c-3.1 1.2-5 4.1-5 7.4 0 4.4 3.6 8 8 8s8-3.6 8-8c0-3.3-1.9-6.2-4.9-7.4-.5-.2-.8-.8-.5-1.3.2-.5.8-.8 1.3-.5 3.7 1.5 6.1 5.1 6.1 9.2 0 5.5-4.5 10-10 10z" />
              <path d="m12 14c-.6 0-1-.4-1-1v-12c0-.6.4-1 1-1s1 .4 1 1v12c0 .6-.4 1-1 1z" />
            </svg>
          </button>

          {showPowerOptions && (
            <div className="power-options">
              <button
                className="power-option shutdown"
                onClick={() => {
                  ActionFromClient("shutdown", 0);
                }}
              >
                <svg
                  enable-background="new 0 0 24 24"
                  height="24"
                  viewBox="0 0 24 24"
                  width="24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path d="m12 24c-5.5 0-10-4.5-10-10 0-4.1 2.4-7.7 6.2-9.2.5-.2 1.1 0 1.3.5s0 1.1-.5 1.3c-3.1 1.2-5 4.1-5 7.4 0 4.4 3.6 8 8 8s8-3.6 8-8c0-3.3-1.9-6.2-4.9-7.4-.5-.2-.8-.8-.5-1.3.2-.5.8-.8 1.3-.5 3.7 1.5 6.1 5.1 6.1 9.2 0 5.5-4.5 10-10 10z" />
                  <path d="m12 14c-.6 0-1-.4-1-1v-12c0-.6.4-1 1-1s1 .4 1 1v12c0 .6-.4 1-1 1z" />
                </svg>
              </button>
              <button
                className="power-option restart"
                onClick={() => {
                  ActionFromClient("restart", 0);
                }}
              >
                <svg
                  width="24"
                  height="24"
                  viewBox="0 0 16 16"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    clip-rule="evenodd"
                    d="M12.75 8a4.5 4.5 0 0 1-8.61 1.834l-1.391.565A6.001 6.001 0 0 0 14.25 8 6 6 0 0 0 3.5 4.334V2.5H2v4l.75.75h3.5v-1.5H4.352A4.5 4.5 0 0 1 12.75 8z"
                  />
                </svg>
              </button>
              <button
                className="power-option sleep"
                onClick={() => {
                  ActionFromClient("sleep", 0);
                }}
              >
                <svg
                  height="24"
                  viewBox="0 0 32 32"
                  width="24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path d="m13.5025 5.4136a15.0755 15.0755 0 0 0 11.5935 18.1946 11.1134 11.1134 0 0 1 -7.9749 3.3893c-.1385 0-.2782.0051-.4178 0a11.0944 11.0944 0 0 1 -3.2008-21.5839m1.4775-2.4136a1.0024 1.0024 0 0 0 -.1746.0156 13.0959 13.0959 0 0 0 1.8246 25.9817c.1641.006.3282 0 .4909 0a13.0724 13.0724 0 0 0 10.702-5.5556 1.0094 1.0094 0 0 0 -.7833-1.5644 13.08 13.08 0 0 1 -11.1504-17.4973 1.0149 1.0149 0 0 0 -.9092-1.38z" />
                  <path d="m0 0h32v32h-32z" fill="none" />
                </svg>
              </button>
            </div>
          )}
        </div>
      </div>

      <div className="content-area">
        <Routes>
          <Route path="/" element={<Offline />} />
          <Route path="/m" element={<Online />} />
        </Routes>
      </div>
    </div>
  );
};

export default App;

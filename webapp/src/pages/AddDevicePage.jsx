import React, { useState } from "react";
import { Copy, CheckCheck, AlertCircle, Check } from "lucide-react";

const AddDevicePage = () => {
  const [DeviceName, setDeviceName] = useState("");
  const [DevicePassword, setDevicePassword] = useState("");
  const [connectionCode, setConnectionCode] = useState("");
  const [copied, setCopied] = useState(false);
  const [addStatus, setAddStatus] = useState({
    status: null,
    message: "",
  });
  const [isLoading, setIsLoading] = useState(false);

  const backendURL = import.meta.env.VITE_BACKEND_URL;

  const handleAddDevice = async () => {
    setAddStatus({ status: null, message: "" });
    setConnectionCode("");
    setIsLoading(true);

    try {
      setIsLoading(true);

      const response = await fetch(`${backendURL}/api/getdeviceID`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          DeviceName,
          DevicePassword,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        console.error("Raw error response:", errorText);
        throw new Error(errorText || "Failed to add device");
      }

      const data = await response.json();
      setConnectionCode(data.connectionCode);

      setAddStatus({
        status: "success",
        message: "Device added successfully!",
      });
    } catch (error) {
      // Handle errors
      console.error("Error adding device:", error);
      setAddStatus({
        status: "error",
        message: error.message || "Failed to add device",
      });
    } finally {
      setIsLoading(false);
    }
  };

  const handleCopyCode = () => {
    if (connectionCode) {
      navigator.clipboard.writeText(connectionCode);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  return (
    <div>
      <h1 className="text-2xl font-bold mb-6">Add New Device</h1>
      <div className="bg-gray-800 p-6 rounded-lg space-y-4">
        {/* Status Message */}
        {addStatus.status && (
          <div
            className={`
              p-3 rounded flex items-center space-x-2
              ${
                addStatus.status === "success"
                  ? "bg-green-600/20 text-green-400"
                  : "bg-red-600/20 text-red-400"
              }
            `}
          >
            {addStatus.status === "success" ? <Check /> : <AlertCircle />}
            <span>{addStatus.message}</span>
          </div>
        )}

        <div>
          <label className="block mb-2">Device Name</label>
          <input
            type="text"
            value={DeviceName}
            onChange={(e) => setDeviceName(e.target.value)}
            className="w-full bg-gray-700 p-2 rounded"
            placeholder="Enter device name"
            disabled={isLoading}
          />
        </div>

        <div>
          <label className="block mb-2">Device Password</label>
          <input
            type="password"
            value={DevicePassword}
            onChange={(e) => setDevicePassword(e.target.value)}
            className="w-full bg-gray-700 p-2 rounded"
            placeholder="Set a device password"
            disabled={isLoading}
          />
        </div>

        {connectionCode && (
          <div>
            <label className="block mb-2">Connection Code</label>
            <div className="flex items-center space-x-2">
              <input
                type="text"
                value={connectionCode}
                readOnly
                className="flex-1 bg-gray-700 p-2 rounded"
              />
              <button
                onClick={handleCopyCode}
                className="bg-blue-600 hover:bg-blue-700 p-2 rounded"
              >
                {copied ? <CheckCheck /> : <Copy />}
              </button>
            </div>
            <p className="text-sm text-gray-400 mt-2">
              Use this code on the client device to establish connection
            </p>
          </div>
        )}

        <button
          onClick={handleAddDevice}
          className={`
            w-full p-2 rounded 
            ${
              isLoading
                ? "bg-gray-600 cursor-not-allowed"
                : "bg-green-600 hover:bg-green-700"
            }
          `}
          disabled={isLoading}
        >
          {isLoading ? "Adding Device..." : "Add Device"}
        </button>
      </div>
    </div>
  );
};

export default AddDevicePage;

import React, { useState, useEffect } from "react";
import { Timer, Lock } from "lucide-react";
const backendURL = import.meta.env.VITE_BACKEND_URL;

const performDeviceCommand = async (
  deviceId,
  Command,
  password,
  scheduleTime = null
) => {
  try {
    const response = await fetch(`${backendURL}/api/sendcommand`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        deviceId,
        Command,
        password,
        scheduleTime,
      }),
    });

    if (!response.ok) {
      const errorResponse = await response.text();
      const errorMessage = errorResponse || "Unknown error occurred";
      throw new Error(errorMessage);
    }

    const result = await response.json();
    return result;
  } catch (error) {
    console.error("API call failed:", error);
    throw error; // Re-throw the error to be caught by the calling function
  }
};

const DevicesPage = () => {
  const [devices, setDevices] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchDevices = async () => {
    try {
      const response = await fetch(`${backendURL}/api/devices`);

      if (!response.ok) {
        throw new Error("Failed to fetch devices");
      }

      const data = await response.json();
      setDevices(data);
      setIsLoading(false);
    } catch (err) {
      setError("Failed to fetch devices");
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchDevices();
  }, []);

  if (isLoading) return <div className="text-white">Loading devices...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div>
      <h1 className="text-2xl font-bold mb-6">Your Devices</h1>
      <div className="space-y-4">
        {devices.map((device) => (
          <DeviceCard
            key={device.id}
            device={device}
            onDeviceUpdate={(updatedDevice) => {
              setDevices(
                devices.map((d) =>
                  d.id === updatedDevice.id ? updatedDevice : d
                )
              );
            }}
          />
        ))}
      </div>
    </div>
  );
};

const DeviceCard = ({ device, onDeviceUpdate }) => {
  const [passwordModal, setPasswordModal] = useState(false);
  const [selectedCommand, setSelectedCommand] = useState(null);
  const [scheduleTime, setScheduleTime] = useState("00:00");
  const [password, setPassword] = useState("");
  const [CommandError, setCommandError] = useState(null);

  const handleDeviceCommand = async () => {
    try {
      setCommandError(null);

      // Convert scheduleTime (hh:mm) to total minutes
      const convertToMinutes = (time) => {
        if (time === "00:00") return 0; // No scheduling
        const [hours, minutes] = time.split(":").map(Number); // Split and convert to numbers
        return `${hours * 60 + minutes}`;
      };

      const scheduleTimeInMinutes = convertToMinutes(scheduleTime);

      const result = await performDeviceCommand(
        device.id,
        selectedCommand,
        password,
        scheduleTimeInMinutes
      );

      if (result.success) {
        const updatedDevice = {
          ...device,
          scheduledCommand:
            scheduleTime !== "00:00"
              ? {
                  type: selectedCommand,
                  remainingTime: `${Math.floor(scheduleTimeInMinutes / 60)
                    .toString()
                    .padStart(2, "0")}:${(scheduleTimeInMinutes % 60)
                    .toString()
                    .padStart(2, "0")}`,
                }
              : null,
        };

        onDeviceUpdate(updatedDevice);
        setPasswordModal(false);
        alert(result.message || `${selectedCommand} successful`);
      } else {
        setCommandError(result.message || "Command failed");
      }
    } catch (error) {
      console.error("Error details:", error);

      if (error.message === "Device is offline") {
        setCommandError("Device is offline");
      } else {
        setCommandError(error.message || "Failed to perform Command");
      }
    }
  };

  const openPasswordModal = (Command) => {
    setSelectedCommand(Command);
    setPasswordModal(true);
    setScheduleTime("00:00");
    setPassword("");
    setCommandError(null);
  };

  const cancelScheduledCommand = async () => {
    try {
      const result = await performDeviceCommand(device.id, "cancel", password);

      if (result.success) {
        const updatedDevice = {
          ...device,
          scheduledCommand: null,
        };

        onDeviceUpdate(updatedDevice);
        setPasswordModal(false);
        alert("Scheduled Command cancelled");
      } else {
        setCommandError(result.message || "Failed to cancel Command");
      }
    } catch (error) {
      setCommandError("Failed to cancel Command");
    }
  };

  return (
    <div className="bg-gray-800 p-4 rounded-lg flex justify-between items-center">
      <div>
        <div className="flex items-center space-x-2">
          <h2 className="font-semibold">{device.name}</h2>
          {device.scheduledCommand && (
            <div
              className="bg-yellow-500/20 text-yellow-400 px-2 py-1 rounded-full text-xs flex items-center space-x-1"
              title={`Scheduled ${device.scheduledCommand.type}`}
            >
              <Timer size={12} />
              <span>{device.scheduledCommand.remainingTime}</span>
            </div>
          )}
        </div>
        <p
          className={`
            text-sm 
            ${device.status === "online" ? "text-green-500" : "text-red-500"}
          `}
        >
          {device.status} • {device.onlineSince}
        </p>
      </div>

      {/* Device Commands */}
      <div className="flex space-x-2">
        {device.scheduledCommand ? (
          <button
            onClick={() => openPasswordModal("cancel")}
            className="bg-red-500 hover:bg-red-600 text-white p-2 rounded"
          >
            Cancel Scheduled Command
          </button>
        ) : (
          <>
            <button
              onClick={() => openPasswordModal("Shutdown")}
              className="bg-[#2E3440] hover:bg-[#3B4252] text-white border border-[#4C566A] py-2 px-4 rounded-md transition"
            >
              Shutdown
            </button>
            <button
              onClick={() => openPasswordModal("Restart")}
              className="bg-[#2E3440] hover:bg-[#3B4252] text-white border border-[#4C566A] py-2 px-4 rounded-md transition"
            >
              Restart
            </button>
            <button
              onClick={() => openPasswordModal("Sleep")}
              className="bg-[#2E3440] hover:bg-[#3B4252] text-white border border-[#4C566A] py-2 px-4 rounded-md transition"
            >
              Sleep
            </button>
          </>
        )}
      </div>

      {/* Password Modal */}
      {passwordModal && (
        <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
          <div className="bg-gray-800 p-6 rounded-lg w-96">
            <h2 className="text-xl font-bold mb-4 capitalize">
              {selectedCommand === "cancel"
                ? "Cancel Scheduled Command"
                : `${selectedCommand}`}
            </h2>

            {/* Time input */}
            <div className="mb-4">
              <label className="block mb-2">Schedule Time (HH:MM)</label>
              <input
                type="time"
                value={scheduleTime}
                onChange={(e) => setScheduleTime(e.target.value)}
                className="w-full bg-gray-700 p-2 rounded"
              />
            </div>

            {/* Password Input */}
            <div className="mb-4">
              <label className="block mb-2 flex items-center">
                <Lock size={16} className="mr-2" />
                Device Password
              </label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full bg-gray-700 p-2 rounded"
                placeholder="Enter device password"
              />
            </div>

            {/* Error Message */}
            {CommandError && (
              <div className="text-red-500 mb-4">{CommandError}</div>
            )}

            <div className="flex space-x-2">
              <button
                onClick={() =>
                  selectedCommand === "cancel"
                    ? cancelScheduledCommand()
                    : handleDeviceCommand()
                }
                className="flex-1 bg-green-600 hover:bg-green-700 p-2 rounded"
              >
                {selectedCommand === "cancel" ? "Cancel Command" : "Confirm"}
              </button>
              <button
                onClick={() => {
                  setPasswordModal(false);
                  setCommandError(null);
                }}
                className="flex-1 bg-gray-700 hover:bg-gray-600 p-2 rounded"
              >
                Close
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default DevicesPage;

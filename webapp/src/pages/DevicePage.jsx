import React, { useState, useEffect } from "react";
import { Timer, Lock } from "lucide-react";

const performDeviceAction = async (
  deviceId,
  action,
  password,
  scheduleTime = null
) => {
  try {
    //do actiosnss
    const response = await fetch("/api/device-action", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        deviceId,
        action,
        password,
        scheduleTime,
      }),
    });

    const result = await response.json();
    return result;
  } catch (error) {
    console.error("API call failed:", error);
    throw error;
  }
};

const DevicesPage = () => {
  const [devices, setDevices] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  //get initial fetch
  useEffect(() => {
    const fetchDevices = async () => {
      try {
        const mockDevices = [
          {
            id: "1",
            name: "Home Desktop",
            status: "online",
            onlineSince: "2 mins ago",
            scheduledAction: null,
          },
          {
            id: "2",
            name: "Work Laptop",
            status: "online",
            onlineSince: "30 mins ago",
            scheduledAction: {
              type: "shutdown",
              remainingTime: "01:45:30",
            },
          },
        ];

        setDevices(mockDevices);
        setIsLoading(false);
      } catch (err) {
        setError("Failed to fetch devices");
        setIsLoading(false);
      }
    };

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
  const [selectedAction, setSelectedAction] = useState(null);
  const [scheduleTime, setScheduleTime] = useState("00:00");
  const [password, setPassword] = useState("");
  const [actionError, setActionError] = useState(null);

  const handleDeviceAction = async () => {
    try {
      setActionError(null);
      const result = await performDeviceAction(
        device.id,
        selectedAction,
        password,
        scheduleTime !== "00:00" ? scheduleTime : null
      );

      if (result.success) {
        const updatedDevice = {
          ...device,
          scheduledAction:
            scheduleTime !== "00:00"
              ? {
                  type: selectedAction,
                  remainingTime: scheduleTime,
                }
              : null,
        };

        onDeviceUpdate(updatedDevice);
        setPasswordModal(false);
        alert(result.message || `${selectedAction} successful`);
      } else {
        setActionError(result.message || "Action failed");
      }
    } catch (error) {
      setActionError("Failed to perform action");
    }
  };

  const openPasswordModal = (action) => {
    setSelectedAction(action);
    setPasswordModal(true);
    setScheduleTime("00:00");
    setPassword("");
    setActionError(null);
  };

  const cancelScheduledAction = async () => {
    try {
      const result = await performDeviceAction(device.id, "cancel", password);

      if (result.success) {
        const updatedDevice = {
          ...device,
          scheduledAction: null,
        };

        onDeviceUpdate(updatedDevice);
        setPasswordModal(false);
        alert("Scheduled action cancelled");
      } else {
        setActionError(result.message || "Failed to cancel action");
      }
    } catch (error) {
      setActionError("Failed to cancel action");
    }
  };

  return (
    <div className="bg-gray-800 p-4 rounded-lg flex justify-between items-center">
      <div>
        <div className="flex items-center space-x-2">
          <h2 className="font-semibold">{device.name}</h2>
          {device.scheduledAction && (
            <div
              className="bg-yellow-500/20 text-yellow-400 px-2 py-1 rounded-full text-xs flex items-center space-x-1"
              title={`Scheduled ${device.scheduledAction.type}`}
            >
              <Timer size={12} />
              <span>{device.scheduledAction.remainingTime}</span>
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

      {/* Device Actions */}
      <div className="flex space-x-2">
        {device.scheduledAction ? (
          <button
            onClick={() => openPasswordModal("cancel")}
            className="bg-red-500 hover:bg-red-600 text-white p-2 rounded"
          >
            Cancel Scheduled Action
          </button>
        ) : (
          <>
            <button
              onClick={() => openPasswordModal("shutdown")}
              className="bg-[#2E3440] hover:bg-[#3B4252] text-white border border-[#4C566A] py-2 px-4 rounded-md transition"
            >
              Shutdown
            </button>
            <button
              onClick={() => openPasswordModal("restart")}
              className="bg-[#2E3440] hover:bg-[#3B4252] text-white border border-[#4C566A] py-2 px-4 rounded-md transition"
            >
              Restart
            </button>
            <button
              onClick={() => openPasswordModal("sleep")}
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
              {selectedAction === "cancel"
                ? "Cancel Scheduled Action"
                : `${selectedAction}`}
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
            {actionError && (
              <div className="text-red-500 mb-4">{actionError}</div>
            )}

            <div className="flex space-x-2">
              <button
                onClick={() =>
                  selectedAction === "cancel"
                    ? cancelScheduledAction()
                    : handleDeviceAction()
                }
                className="flex-1 bg-green-600 hover:bg-green-700 p-2 rounded"
              >
                {selectedAction === "cancel" ? "Cancel Action" : "Confirm"}
              </button>
              <button
                onClick={() => {
                  setPasswordModal(false);
                  setActionError(null);
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

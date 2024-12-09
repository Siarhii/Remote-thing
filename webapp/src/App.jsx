import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import SideNavbar from "./components/SideBar";
import AddDevicePage from "./pages/AddDevicePage";
import DevicesPage from "./pages/DevicePage";

function App() {
  return (
    <Router>
      <div className="flex h-screen bg-gray-900 text-white">
        <SideNavbar />
        <div className="flex-1 bg-gray-900 p-6 overflow-y-auto">
          <Routes>
            <Route path="/" element={<DevicesPage />} />
            <Route path="/add-device" element={<AddDevicePage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;

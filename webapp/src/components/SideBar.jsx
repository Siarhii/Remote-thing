import React from "react";
import { Link, useLocation } from "react-router-dom";
import { UserPlus, Monitor } from "lucide-react";

const SideNavbar = () => {
  const location = useLocation();

  const NavIcon = ({ to, icon: Icon, label }) => {
    const isActive = location.pathname === to;

    return (
      <Link
        to={to}
        className={`
          p-3 rounded-lg transition-all duration-200 
          ${
            isActive
              ? "bg-blue-600 text-white"
              : "text-gray-400 hover:bg-gray-700 hover:text-white"
          }
        `}
      >
        <Icon />
      </Link>
    );
  };

  return (
    <div className="w-20 bg-gray-800 flex flex-col items-center py-4 space-y-4">
      <NavIcon to="/" icon={Monitor} label="Devices" />
      <NavIcon to="/add-device" icon={UserPlus} label="Add Device" />
    </div>
  );
};

export default SideNavbar;

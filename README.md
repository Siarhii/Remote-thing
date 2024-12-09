# 🚀 Remote Shutdown App

![GitHub stars](https://img.shields.io/github/stars/Siarhii/Remote-thing?style=social)  
![License](https://img.shields.io/github/license/Siarhii/Remote-thing)  
![GitHub issues](https://img.shields.io/github/issues/Siarhii/Remote-thing)  
![GitHub last commit](https://img.shields.io/github/last-commit/Siarhii/Remote-thing)

A remote shutdown and monitoring application that allows users to schedule **shutdown, sleep, and restart commands** for their homelabs. The app supports **both offline and online modes**, with ongoing development for advanced features.

---

## 📋 Features

### 🔌 Offline Tab (Fully Functional)

- **Schedule Shutdown, Sleep, and Restart**:  
  Allows users to schedule actions with a maximum limit of **1 month**.
- **Task Scheduling (Future Update)**:  
  Plan tasks such as shutting down the device after completing downloads or gaming.

### 🌐 Online Tab (Partially Functional)

- **User Authentication**:  
  Login functionality for managing connected devices.
- **Device Management**:  
  Add devices and monitor their uptime status _(In Progress)_.
- **Remote Command Scheduling**:  
  Backend supports **unique code generation**, and the client app connects to the backend. **Currently under development with known errors.**

### 🔒 Security (Planned Enhancements)

- Implement secure communication protocols for all backend and frontend interactions.
- Add resource monitoring features like **CPU usage, RAM utilization, and device temperature stats**.

### 🛡️ Background Mode (In Progress)

- Enable the app to run in the background for continuous monitoring without needing the main window to remain open.

---

## 🎨 Screenshots

### 🖥️ Offline Tab

![Offline Tab Screenshot](SS/offline.png)

### 🌐 Online Tab

![Online Tab Screenshot](SS/online.png)

### 📊 Web App Screenshots

- **Dashboard**  
  ![add Device page](SS/addDevice.png)

- **Device Management**  
  ![devices page](SS/devicesPage.png)

- **Command Scheduler**  
  ![Execute Command](SS/executeCommand.png)

---

## 🛠️ Installation

### Option 1: Precompiled Binary

Download the executable: **[remote-shutdown.exe](#)**

### Option 2: Build From Source

Follow these steps to clone and build the app locally:

```bash
# Clone the repository
git clone https://github.com/Siarhii/Remote-thing.git

# Navigate into the project folder
cd Remote-thing/client

# Build the app (requires Go and Wails)
wails build

# Run the application
./Remote-shutdown.exe


💡 Usage
Offline Mode:
Use the Offline Tab to schedule shutdown/sleep/restart tasks with a specified timer.
Online Mode (Beta):
Log into your account.
Add devices using the unique codes generated by the backend.
Manage remote commands (note: under development and may throw errors).
🚧 Roadmap
✅ Polished frontend UI for offline and online tabs.
✅ Unique code generation for backend.
⬜ Full online tab functionality (remote command scheduling, device monitoring).
⬜ Enhanced security protocols.
⬜ Add advanced monitoring stats (CPU, RAM, temperature).
👥 Contributors
Siarhii - Developer
📜 License
This project is licensed under the MIT License - see the LICENSE file for details.

🤝 Contributing
Contributions, issues, and feature requests are welcome!
Feel free to check the issues page.

🌟 Show Your Support
If you like this project, give it a ⭐️ and share it with others!
```

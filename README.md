# 🚀 Remote Shutdown App

![GitHub stars](https://img.shields.io/github/stars/Siarhii/Remote-thing?style=social)  
![License](https://img.shields.io/github/license/Siarhii/Remote-thing)  
![GitHub issues](https://img.shields.io/github/issues/Siarhii/Remote-thing)  
![GitHub last commit](https://img.shields.io/github/last-commit/Siarhii/Remote-thing)

A comprehensive remote shutdown and monitoring application that allows users to schedule **shutdown, sleep, and restart commands** for their homelabs. The app now supports **fully functional offline and online modes** with advanced device management capabilities.

---

## 📋 Features

### 🔌 Offline Tab (Fully Functional)

- **Schedule Shutdown, Sleep, and Restart**:  
  Precise scheduling of system actions with a maximum limit of **1 month**.
- **Flexible Task Planning**:  
  Easily plan and manage system tasks with intuitive controls.

### 🌐 Online Tab (Now Fully Operational)

- **Seamless User Authentication**:  
  Secure login functionality for managing multiple connected devices.
- **Advanced Device Management**:
  - Add and manage devices with unique identification
  - Real-time device status monitoring
  - Instant remote command execution

### 🔒 Security Features

- Secure communication protocols for backend and frontend interactions
- Unique code generation for device pairing
- Robust authentication mechanisms

### 🛡️ Background Mode

- Continuous monitoring capability
- Runs efficiently in the background without keeping the main window open

---

## 🎨 Screenshots

### 🖥️ Offline Tab

![Offline Tab Screenshot](SS/offline.png)

### 🌐 Online Tab (New!)

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

Download the executable: **[remote-shutdown.exe](./remote-shutdown.exe)**

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
```

### 💡 Usage

#### Offline Mode

- Use the **Offline Tab** to schedule shutdown, sleep, or restart tasks with a specified timer.

#### Online Mode

1. Log into your account
2. Add devices using unique generated codes
3. Manage and execute remote commands with ease
4. Monitor device status in real-time

---

### 🚧 Roadmap and Progress

- ✅ Working offline tab
- ✅ Polished frontend UI for offline and online tabs
- ✅ Unique code generator for client apps
- ✅ Authentication and remote command scheduling
- ✅ Online device management
- ⬜ Enhance local app storage
- ⬜ Improve auto socket reconnection
- ⬜ Add advanced monitoring stats (CPU, RAM, temperature)
- ⬜ Implement custom command scheduling for specific use cases

---

### 👥 Contributors

- **[Siarhii](https://github.com/Siarhii)** - Developer

---

### 📜 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---

### 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check the [issues page](https://github.com/Siarhii/Remote-thing/issues).

---

### 🌟 Show Your Support

If you like this project, **give it a ⭐️** and share it with others!

**Recent Updates:**

- Online mode is now fully functional
- Added comprehensive device management
- Improved authentication and remote command execution

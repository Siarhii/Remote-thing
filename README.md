# 🚀 Remote Shutdown App

![GitHub stars](https://img.shields.io/github/stars/Siarhii/Remote-thing?style=social)  
![License](https://img.shields.io/github/license/Siarhii/Remote-thing)  
![GitHub issues](https://img.shields.io/github/issues/Siarhii/Remote-thing)  
![GitHub last commit](https://img.shields.io/github/last-commit/Siarhii/Remote-thing)

A comprehensive remote shutdown and monitoring application that allows users to schedule **shutdown, sleep, and restart commands** for their homelabs. The app supports **offline and online modes** with device management capabilities.

---

## 📋 Features

### 🔌 Offline Tab (Fully Functional)

- **Schedule Shutdown, Sleep, and Restart**:  
  Precise scheduling of system actions with a maximum limit of **1 month**.
- **Flexible Task Planning**:  
  Easily plan and manage system tasks with intuitive controls.

### 🌐 Online Tab (Developing)

- **Device Password Protection**:  
  Secure device addition with password-based access
- **Device Management**:
  - Add devices with unique identifiers
  - Monitor device status
  - Execute remote commands

### 🔒 Current Security Approach

- Password protection for device addition
- Unique code generation for device pairing

### 🛡️ Background Mode

- Continuous monitoring capability
- Runs efficiently in the background without keeping the main window open

---

## 🎨 Screenshots

### 🖥️ Offline Tab

![Offline Tab Screenshot](SS/offline.png)

### 🌐 Online Tab

![Online Tab Screenshot](SS/online.png)

### 📊 Web App Screenshots

- **Add Device**  
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

1. Add devices using unique generated codes
2. Set device-specific passwords
3. Manage and execute remote commands
4. Monitor device status

---

### 🚧 Roadmap and Progress

- ✅ Working offline tab
- ✅ Polished frontend UI for offline and online tabs
- ✅ Unique code generator for device addition
- ✅ Device password protection
- ✅ Basic online device management
- ⬜ Full user authentication system
- ⬜ User profile management
- ⬜ Advanced security features
- ⬜ Add advanced monitoring stats (CPU, RAM, temperature)
- ⬜ Implement custom command scheduling for specific use cases

---

# 🚀 Remote Shutdown App

![GitHub stars](https://img.shields.io/github/stars/Siarhii/Remote-thing?style=social)  
![License](https://img.shields.io/github/license/Siarhii/Remote-thing)  
![GitHub issues](https://img.shields.io/github/issues/Siarhii/Remote-thing)  
![GitHub last commit](https://img.shields.io/github/last-commit/Siarhii/Remote-thing)

A comprehensive remote shutdown and monitoring application that allows users to schedule **shutdown, sleep, and restart commands** for their homelabs. The app supports **offline and online modes** with device management capabilities.

---

## How This Works

### Behind the Scenes

So, you're wondering how this app came to life? Let me break it down for you:

- **Desktop App**: Built with Go and Wails
- **Backend**: Pure Go with socket magic
- **Web Interface**: React

### The Online Mode Adventure

Here's how you can use it:

1. **Device Registration**
   Hop into the web app's "Add Devices" tab.

   - Create a password
   - The backend will hand you a unique connection code

2. **Connecting the Dots**

   - Open your desktop app
   - Paste that magic code you just got
   - Boom! Instant connection between your device and the backend

3. **Remote Control Time**
   - Want to shut down a computer from across the room (or the world)?
   - Jump into the web app
   - Enter your device's password
   - Execute commands like a tech wizard

### Why I Built This

Real talk - I created this app to solve a few headaches:

- **No More Network Nightmares**: Forget about port forwarding or firewalls
- **Simplicity is King**: Remote control should be EASY,you dont need to do anything,just install the exe and youre done
- **Learning Go**: Because why not challenge myself and build something cool?

### 🔐 Current Security Approach

It's not Fort Knox yet, but i'm working on it:

- Device-specific passwords
- Unique connection codes

**NOTE**: This is a temperory solution which will be changed when full authentication is implemented

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

- Online mode device management implemented
- Password protection for device addition
- Improved remote command execution

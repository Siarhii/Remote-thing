# Remote Shutdown App

A remote shutdown and monitoring application that allows users to schedule shutdown, sleep, and restart commands for thier homelabs. The app features both online and offline modes, with the offline tab fully functional and the online tab under development.

## Features

### Offline Tab (Working)

- **Schedule Shutdown, Sleep, and Restart**: Users can schedule their device to shut down, go to sleep, or restart after a timer, with a maximum limit of **1 month**.
- **Task Scheduling (to be added)**: Users can also add tasks like shutting down the device after a game or software download completes.

### Online Tab (Under Development)

- **User Authentication**: Users will be able to log into their accounts to manage their devices.
- **Device Management**: Users can add client devices to their account and monitor the uptime of those devices.
- **Remote Command Scheduling**: Once logged in, users can remotely schedule commands for their devices.

### Security (Future Implementation)

- somehow make this more secure
- add more monitoring stats like cpu,ram usage etc

### Background Mode (In Progress)

- **Keep App Running in Background**: The app will run in the background even when closed, allowing continuous monitoring and scheduling without needing to keep the window open.

## Screenshots

![App Screenshot](SS/offline.png)
![Online Tab (Under Development)](SS/online.png)

## Installation

**You can directly download the remote-shutdown.exe**

OR

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Siarhii/Remote-thing.git
   ```

2. **Navigate into the project folder**:

   ```bash
   cd Remote-thing/client
   ```

3. **Install dependencies**:

   If you have Go and wails installed, simply run:

   ```bash
   wails build
   ```

4. **Run the application**:

   After building the app, run it with:

   ```bash
   ./Remote-shutdown.exe
   ```

## Usage

- **To schedule a task** in the offline tab, select a shutdown/sleep/restart action and specify the timer.
- **In the online tab**, log into your account, add devices, and start scheduling remote commands.

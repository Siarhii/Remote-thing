# Remote Shutdown App

A remote shutdown and monitoring application that allows users to schedule shutdown, sleep, and restart commands on their devices. The app features both online and offline modes, with the offline tab fully functional and the online tab under development.

## Features

### Offline Tab (Working)
- **Schedule Shutdown, Sleep, and Restart**: Users can schedule their device to shut down, go to sleep, or restart after a timer, with a maximum limit of **1 month**.
- **Task Scheduling**: Users can also add tasks like shutting down the device after a game or software download completes.

### Online Tab (Under Development)
- **User Authentication**: Users will be able to log into their accounts to manage their devices.
- **Device Management**: Users can add client devices to their account and monitor the uptime of those devices.
- **Remote Command Scheduling**: Once logged in, users can remotely schedule commands for their devices.

### Security (Future Implementation)
- **Secure Login**: Plans to add secure user authentication to prevent unauthorized access to user accounts and devices.
- **Encrypted Communication**: All communications between the client app and the server will be secured.

### Background Mode (In Progress)
- **Keep App Running in Background**: The app will run in the background even when closed, allowing continuous monitoring and scheduling without needing to keep the window open.

## Screenshots

![App Screenshot](path/to/your/image.png)

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/your-username/remote-shutdown-app.git
    ```

2. **Navigate into the project folder**:

    ```bash
    cd remote-shutdown-app
    ```

3. **Install dependencies**:

    If you have Go installed, simply run:

    ```bash
    go build
    ```

4. **Run the application**:

    After building the app, run it with:

    ```bash
    ./remote-shutdown-app
    ```

## Usage

- **To schedule a task** in the offline tab, select a shutdown/sleep/restart action and specify the timer.
- **In the online tab**, log into your account, add devices, and start scheduling remote commands.



## Navy Six Missile Managment System

<img src="https://github.com/liel-almog/navy-six/blob/main/images/logo.webp" width="200" height="200">

- [Navy Six Missile Managment System](#navy-six-missile-managment-system)
  - [Features](#features)
    - [Store New Missiles:](#store-new-missiles)
    - [Launch Missile:](#launch-missile)
    - [Inventory Report:](#inventory-report)
    - [Clean Out Missiles:](#clean-out-missiles)
    - [Shutdown:](#shutdown)
  - [Technical Overview](#technical-overview)
  - [How to Use:](#how-to-use)
  - [Installation](#installation)


Welcome to the Missile Management System console application README. This application provides a command-line interface for managing missile inventory, including storing new missiles, launching missiles, generating inventory reports, cleaning out old missiles, and safely shutting down the system.
### Features

The Missile Management System includes the following options within its menu:

<img src="https://github.com/liel-almog/navy-six/blob/main/images/menu.png" width="200" height="200">

#### Store New Missiles: 
Add new missiles to your inventory.

<img src="https://github.com/liel-almog/navy-six/blob/main/images/store_missiles.png" width="500" height="250">

#### Launch Missile: 
Select and launch a missile from your inventory.

<img src="https://github.com/liel-almog/navy-six/blob/main/images/launch.png" width="500" height="250">

#### Inventory Report: 
Generate a detailed report of the current missile inventory.

<img src="https://github.com/liel-almog/navy-six/blob/main/images/inventory_report.png" width="300" height="300">

#### Clean Out Missiles: 
Remove old, unused, failed missiles from the inventory.

#### Shutdown: 
Safely shut down the Missile Management System application.
 
### Technical Overview

The application leverages Go (Golang), known for its efficiency and simplicity, especially in concurrent tasks and system-level programming. The menu-driven interface is intuitive, facilitating quick navigation and operation within the system.

### How to Use:

To use the Missile Management System, follow these steps:

1. Start the application in your console.
2. The main menu will be displayed with the options listed above.
3. Enter the number corresponding to the action you wish to perform.
4. Follow the on-screen instructions for each action.

### Installation

To install and run the Missile Management System, ensure you have Go installed on your system, clone the repository, and execute the main application file.

```bash
git clone <repository-url>
cd <repository-directory>
go run main.go
```
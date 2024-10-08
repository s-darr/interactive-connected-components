# Interactive Grid Visualisation of Connected Components

This project demonstrates the use of algorithms like **Union-Find**, **Breadth-First Search (BFS)** and **Depth-First Search (DFS)** to visualise the **largest connected component** in a grid, where each cell can either represent **land (1)** or **water (0)**.

## Introduction

In this project, we have an interactive grid where each cell can either be land (`1`) or water (`0`). The goal is to visualise the largest connected component dynamically as the user toggles between land and water cells. A land cell is considered connected with land cells that are adjacent (not connected diagonally).

This problem is solved using two algorithms:

- **Depth-First Search (DFS)**
- **Breadth-First Search (BFS)**

## Features

- **Interactive Grid**: Click on cells to toggle between land and water. The grid dynamically updates to reflect the largest connected component.
- **Real-time Updates**: As you modify the grid, the largest connected component is recalculated and displayed at the top of the window.

## How to Run

### Prerequisites

- **Golang**: You can download from [https://golang.org/dl/](https://golang.org/dl/).
- **Fyne**: You can download from [https://docs.fyne.io/started/]. Follow the instructions to install any prerequisites if needed.

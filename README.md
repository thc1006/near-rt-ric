# Near-RT RIC

This repository contains the components for a Near-Real-Time RAN Intelligent Controller (Near-RT RIC), a key component of the O-RAN architecture.

## Overview

The Near-RT RIC is a platform that enables near-real-time control and optimization of the Radio Access Network (RAN). It hosts third-party applications (xApps) that can monitor and control the RAN in near-real-time.

The main components in this repository are:

*   **Dashboard:** A web-based user interface for the Near-RT RIC.
*   **xAPP Dashboard:** A separate dashboard for managing and visualizing xAPPs.
*   **E2 Interface:** The interface that connects the Near-RT RIC to the E2 nodes (e.g., eNBs, gNBs).

### Architecture Diagram

```
+---------------------+       +-----------------------+
|      Dashboard      |<----->| Near-RT RIC Dashboard |
+---------------------+       +-----------------------+
           ^
           |
           v
+---------------------+
|    Near-RT RIC      |
+---------------------+
           ^
           | E2 Interface
           v
+---------------------+
|      E2 Nodes       |
| (eNBs, gNBs, etc.)  |
+---------------------+
```

## Quick-Start

To get started with the Near-RT RIC, you will need the following:

*   **Go:** Version 1.24.0 or later
*   **Node.js:** Version 16.14.2 or later
*   **Kubernetes:** Version 1.25.x

## O-RAN Alliance Specifications

The Near-RT RIC is based on the O-RAN Alliance architecture. You can find more information about the O-RAN architecture and specifications at the following links:

*   [O-RAN Alliance Website](https://www.o-ran.org/)
*   [O-RAN Architecture Specification](https://www.o-ran.org/specifications)
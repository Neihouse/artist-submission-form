# Venue Management System

## Overview

This project is aimed at developing a comprehensive venue management system tailored to manage events, artist submissions, financial tracking, and fan engagement. The system leverages a modern tech stack including Golang, Echo, HTMX, Tailwind CSS, Google Cloud Platform, and various external APIs for seamless integration and automation.

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Setup and Configuration](#setup-and-configuration)
- [Backend Development](#backend-development)
- [Frontend Development](#frontend-development)
- [Integration and Automation](#integration-and-automation)
- [Testing and Deployment](#testing-and-deployment)
- [Monitoring and Maintenance](#monitoring-and-maintenance)
- [Contributing](#contributing)
- [License](#license)

## Features

1. **Artist Submission and Inbound Management**
   - Interactive artist submission forms with Spotify integration.
   - Centralized inbound dashboard for managing submissions.

2. **Calendar and Event Management**
   - Shared digital calendar for event management.
   - Single source of truth for event-related information.

3. **Automation and AI Integration**
   - AI-powered artist recommendations.
   - Automated follow-up reminders.

4. **Offers and Settlements**
   - Templates for creating and managing offers.
   - Automated settlement updates based on ticketing information.

5. **Finance and P&L Management**
   - Finance tools for tracking income and expenses.
   - Automated Profit and Loss (P&L) statements.

6. **Asset Management**
   - Centralized storage for event-related files.
   - Controlled access and sharing capabilities.

7. **Marketing and Fan Engagement**
   - Tools for tracking marketing efforts.
   - Fan profiles creation and engagement tools.

## Tech Stack

### Backend
- **Golang**: Main programming language
- **Echo**: Web framework for building RESTful APIs
- **GORM**: ORM library for database interactions
- **SQLite/PostgreSQL**: Database for storing application data
- **Google Cloud Platform (GCP)**: Hosting and managing resources
  - **Google Cloud Workflows**
  - **Google Cloud Scheduler**
  - **Google Cloud IAM**
  - **Google Cloud Storage**
  - **Google Cloud Firestore**
  - **Google Cloud Functions**
  - **Google Pub/Sub**
- **Pulumi**: Infrastructure as code tool

### Frontend
- **HTML5/CSS3**: Markup and styling for the web interface
- **HTMX**: For dynamic updates without using JavaScript
- **Tailwind CSS**: Utility-first CSS framework

## Setup and Configuration

### Prerequisites

1. **Golang**: Ensure you have Golang installed.
2. **Node.js**: Required for frontend development.
3. **Google Cloud SDK**: For managing Google Cloud resources.
4. **Pulumi**: For infrastructure as code management.

### Clone the Repository
```bash
git clone https://github.com/Neihouse/venue-management-system.git
cd venue-management-system
```

### Environment Variables
Create a `.env` file in the root directory and add the following variables:
```
DATABASE_URL=your_database_url
SPOTIFY_API_KEY=your_spotify_api_key
GOOGLE_CLOUD_PROJECT_ID=your_project_id
GOOGLE_CLOUD_STORAGE_BUCKET=your_storage_bucket
```

## Backend Development

### Directory Structure
```
/backend
  ├── /cmd
  ├── /pkg
  ├── /internal
  ├── /api
  ├── /migrations
  └── main.go
```

### Setting Up
```bash
cd backend
go mod init github.com/Neihouse/venue-management-system
go get -u github.com/labstack/echo/v4
go get -u gorm.io/gorm
```

### Running the Server
```bash
go run main.go
```

## Frontend Development

### Directory Structure
```
/frontend
  ├── /src
  ├── /components
  ├── /pages
  ├── /styles
  └── index.html
```

### Setting Up
```bash
cd frontend
npm install
npm start
```

## Integration and Automation

### External APIs Integration
- Integrate Spotify, Google Cloud Services, and other relevant APIs.
- Set up data enrichment workflows.

### Automation Workflows
- Automate P&L calculations.
- Implement settlement processes using Google Cloud Functions.

## Testing and Deployment

### Testing
- Conduct unit testing for backend and frontend components.
- Perform integration testing.

### Deployment
- Deploy the application to Google Cloud Platform.
- Set up CI/CD pipelines.

## Monitoring and Maintenance

### Monitoring
- Implement monitoring and logging using Google Cloud services.

### Maintenance
- Regularly update dependencies.
- Perform routine maintenance tasks.
- Address user feedback and improvements.

## Contributing
Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) for more information.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

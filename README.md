# Calculator

A simple calculator project with front-end and back-end components, built for learning and demonstration purposes.

## Features

- Front-end and back-end structure
- Basic calculator operations
- Simple and clean interface
- Easy local execution 
- Docker support

## Repository Structure

├── frontend/  
├── backend/  
├── docker-compose.yml  
└── README.md  

## Setup and Run

### Prerequisites

- Docker and Docker Compose (optional)  
- Web browser  

### Clone the Repository

```bash
git clone https://github.com/luanramos15/calculator.git
cd calculator
docker compose up
```

To make requests directly to the back end, just make a POST request to "localhost:8080/" with an example body as:

```bash
{
    "expression" : "2+√4*2+1"
}
```

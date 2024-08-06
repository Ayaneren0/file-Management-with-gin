
## 🌐 Socials:
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?logo=linkedin&logoColor=white)]wwww.linkedin.com/in/ayanahmad15) 

# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white) ![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)
# 📊 GitHub Stats:
![](https://github-readme-stats.vercel.app/api?username=Ayaneren0&theme=dark&hide_border=false&include_all_commits=false&count_private=false)<br/>
![](https://github-readme-streak-stats.herokuapp.com/?user=Ayaneren0&theme=dark&hide_border=false)<br/>

# File Sharing and Management Platform

## Project Description

This project is a simple web application built with Go that allows users to upload, view, and manage files. Users can upload documents, images, and other file types, which are then displayed in a gallery format. The application includes form validation, static file serving, and template rendering.

## Features

- **Home Page**: Displays a list of uploaded files with thumbnails and basic information.
- **Upload Page**: A form for users to upload files with validation (e.g., file type and size restrictions).
- **File Details Page**: Displays detailed information about the file, including a download link.
- **Static Files**: Efficiently serves CSS, JavaScript, and uploaded files (e.g., images).
- **Template Rendering**: Uses Go templates to render pages dynamically.

## Directory Structure

project_root/ <br>
|-- main.go<br>
|-- templates/<br>
|   |-- index.html<br>
|   |-- upload.html<br>
|   |-- file.html<br>
|-- static/<br>
    |-- style.css<br>
    |-- uploads/ (empty, for uploaded files)<br>
    |-- thumbnails/ (empty, for generated thumbnails)<br>
## Getting Started

### Prerequisites

- Go (https://golang.org/doc/install)

### Running the Project

1. Clone the repository:
    ```sh
    git clone https://github.com/Ayaneren0/file-Management-with-gin.git
    cd project_root
    ```

2. Create the directories and files as outlined in the Directory Structure section.

3. Populate the HTML and CSS files with the provided code.

4. Run the application:
    ```sh
    go run main.go
    ```

5. Open your browser and visit `http://localhost:8080`.

## Usage

### Home Page

- Visit `http://localhost:8080` to see the list of uploaded files.

### Upload Page

- Visit `http://localhost:8080/upload` to upload a new file.
- Select a file to upload and click the "Upload" button.
- After uploading, the file will appear on the home page.

### File Details Page

- Click on any file link on the home page to view the file details.
- The file details page will display the image and provide a download link.

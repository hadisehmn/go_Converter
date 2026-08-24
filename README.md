
# GoConvert

GoConvert is a file conversion service built with Go.

The project provides a REST API that allows clients to upload files, convert
them to another format, and download the converted result.

GoConvert was built as a practical project to learn and apply Go, HTTP servers,
file processing, REST API design, clean code, and deployment.

## Features

- Upload files through a REST API
- Detect input file formats
- Convert files to supported formats
- Return converted files to the client
- Support multiple file formats
- HTTP API
- Error handling
- Deployed and available on a production server

## Tech Stack

- **Go** — Application development
- **HTTP** — Server and request handling
- **REST API** — API communication
- **File Processing** — File conversion and handling

## API

### Convert File

Convert an uploaded file to the requested format.

```http
POST /convert
Content-Type: multipart/form-data

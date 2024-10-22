# GO Web Application with Mobile Integration

## About
This project was initiated to explore the capabilities of the GO programming language in building secure and scalable web applications. The development will include best practices in security, efficiency, and cross-platform compatibility.

## Goals
This project aims to create a secure and efficient web application using the GO language, with future integration for iOS and Android platforms. The application includes user authentication, secure data management, and cross-platform capabilities.

## Upcoming Tasks
- [x] **Input Sterilization**: Implement measures to sanitize user input and protect against SQL injection and other malicious attacks.
- [x] **Password bcrypt Hashing**: Securely hash user passwords before storing them in the database.
- [x] **User Authentication**: Develop a robust authentication mechanism to manage user sessions and permissions.
- [x] **Database Encryption**: Ensure the database is encrypted to protect sensitive user data.
- [ ] **Mobile Development**: Extend the application to iOS and Android platforms using relevant tools.
- **Additional Tasks**:
  - [ ] Enhance user interface and experience
  - [ ] Implement error handling and logging
  - [ ] Conduct performance optimization

## Project Setup

### Clone the Repository

```bash
git clone https://github.com/yourusername/website-minigo.git
cd website-minigo
```

### Install Dependencies

```bash
go mod tidy
```

### Database Encryption

The database is encrypted using AES encryption.
Make sure to set the `DB_SECRET_KEY` environment variable in your `.env` file:

```plaintext
DB_SECRET_KEY=your_32_byte_secret_key_goes_here!
```

### Run

```bash
go run main.go
```

This will start the web application on http://localhost:8080.
Logs will be generated for database creation, encryption, decryption, and deletion processes for better visibility during runtime.

Feel free to tweak. 🤓

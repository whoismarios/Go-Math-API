# Math API Documentation

## Base URL
http://localhost:8080

## Endpoints

### 1. Sum of Two Numbers
**URL:** /sum

**Method:** POST

**Request Body:**
{
    "a": int,
    "b": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 15

### 2. Subtraction of Two Numbers
**URL:** /sub

**Method:** POST

**Request Body:**
{
    "a": int,
    "b": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 5

### 3. Multiplication of Two Numbers
**URL:** /mul

**Method:** POST

**Request Body:**
{
    "a": int,
    "b": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 50

### 4. Division of Two Numbers
**URL:** /div

**Method:** POST

**Request Body:**
{
    "a": int,
    "b": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 2
- 400 Bad Request if division by zero.

### 5. Modulus of Two Numbers
**URL:** /mod

**Method:** POST

**Request Body:**
{
    "a": int,
    "b": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 1
- 400 Bad Request if division by zero.

### 6. Power of a Number
**URL:** /pow

**Method:** POST

**Request Body:**
{
    "a": int,
    "b": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 1024

### 7. Square Root of a Number
**URL:** /sqrt

**Method:** POST

**Request Body:**
{
    "a": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 4

### 8. Sine of an Angle (in degrees)
**URL:** /sin

**Method:** POST

**Request Body:**
{
    "a": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 0

### 9. Cosine of an Angle (in degrees)
**URL:** /cos

**Method:** POST

**Request Body:**
{
    "a": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 1

### 10. Tangent of an Angle (in degrees)
**URL:** /tan

**Method:** POST

**Request Body:**
{
    "a": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 0

### 11. Natural Logarithm of a Number
**URL:** /log

**Method:** POST

**Request Body:**
{
    "a": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 2
- 400 Bad Request if the number is non-positive.

### 12. Absolute Value of a Number
**URL:** /abs

**Method:** POST

**Request Body:**
{
    "a": int
}

**Response:**
- 200 OK on success with result.
- Example: Result: 5

### 12. Health Check of the Server
**URL** /health

**Method** GET

**Response**
- 200 OK on success
- Message: `Server is up and running`
- 404 NOT FOUND when server is down
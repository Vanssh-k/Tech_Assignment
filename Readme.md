## Getting Started
1. **Clone the Repository**  
   ```bash
   git clone https://github.com/Vanssh-k/Tech_Assignment.git
   ```

2. **Navigate to the Project Directory**  
   ```bash
   cd Tech_Assignment
   ```

3. **Start the Application (Building Docker Compose)**  
   ```bash
   docker-compose up --build
   ```

## API Routes

1. **POST /signup**  
   This endpoint allows users to sign up or create a new account. The required inputs are the email and password, which should be sent as a JSON body.
   ```bash
   curl --location 'http://localhost:8080/signup' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "email" : "vanshkapoor2001@gmail.com",
       "password" : "123456789"
   }'
   ```

2. **POST /signin**  
   This endpoint is used for user login and returns a JWT access token upon successful authentication. The required inputs are the email and password, which should be sent as a JSON body.
   ```bash
   curl --location 'http://localhost:8080/signin' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "email" : "vanshkapoor2001@gmail.com",
       "password" : "123456789"
   }'
   ```

3. **GET /protected**  
   This endpoint simulates a protected route. Users must provide a valid access token in order to access this route; otherwise, access will be denied.
   ```bash
   curl --location 'http://localhost:8080/protected' \
   --header 'Authorization: Bearer <ACCESS_TOKEN>'
   ```

4. **POST /revoketoken**  
   This endpoint allows users to revoke their access token, effectively logging them out.
   ```bash
   curl --location --request POST 'http://localhost:8080/revoketoken' \
   --header 'Authorization: Bearer <ACCESS_TOKEN>'
   ```

5. **POST /refreshtoken**  
   This endpoint is used to refresh the access token, providing a new token for continued access.
   ```bash
   curl --location --request POST 'http://localhost:8080/refreshtoken' \
   --header 'Authorization: Bearer <ACCESS_TOKEN>'
   ```

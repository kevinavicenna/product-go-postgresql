# **CRUD PostgreSQL using Golang, GoFiber and GORM**

This project demonstrates a simple CRUD (Create, Read, Update, Delete) application using **Golang**, **GORM** (an ORM for Go), and **PostgreSQL**. It includes a RESTful API built with **Fiber**, a fast and lightweight web framework for Go.

---

## **How to Run the Project**

1. **Clone the Repository**  
   Clone the project to your local machine:
   ```bash
   git clone https://github.com/kevinavicenna/product-go-postgresql.git
   cd product-go-postgresql
   ```

2. **Set Up Environment Variables**  
   Rename `.env.example` to `.env` and update the PostgreSQL connection details:
   ```env
   DB_HOST=
   DB_USER=
   DB_PASSWORD=
   DB_NAME=
   DB_PORT=
   ```

3. **Install Dependencies**  
   Install the required Go packages:
   ```bash
   go mod tidy
   ```

4. **Run the Application**  
   Start the server:
   ```bash
   go run main.go
   ```

5. **Access the Application**  
   - Open your browser and navigate to `http://localhost:8080`.  
   - Alternatively, use **Postman** to interact with the API.

---

## **API Endpoints**

| Method | Endpoint              | Description                     |
|--------|-----------------------|---------------------------------|
| POST   | `/api/create_product` | Create a new product            |
| GET    | `/api/products`       | Get all products                |
| GET    | `/api/get_product/:id`| Get a product by ID             |
| PUT    | `/api/update_product/:id` | Update a product by ID      |
| DELETE | `/api/delete_product/:id` | Delete a product by ID      |

---

## **Example Requests**

#### **Create a Product**
- **URL**: `POST /api/create_product`
- **Body**:
  ```json
  {
      "name": "Smartphone",
      "category": "Electronics",
      "description": "A high-end smartphone with 5G support."
  }
  ```

#### **Get All Products**
- **URL**: `GET /api/products`

#### **Get a Product by ID**
- **URL**: `GET /api/get_product/1`

#### **Update a Product**
- **URL**: `PUT /api/update_product/1`
- **Body**:
  ```json
  {
      "name": "Chicken",
      "category": "Animals",
      "description": "A collection of chickens"
  }
  ```

#### **Delete a Product**
- **URL**: `DELETE /api/delete_product/1`



## **Result API Response in Postman**


![Postman Result](img/img.png)


## **Why Include `.env`?**

The `.env` file is included in this project for **learning purposes only**. It contains environment variables like database credentials. However, in a production environment, **never include `.env` files** in your repository for security reasons. Use environment management tools like **Docker secrets**, **AWS Secrets Manager**, or **GitHub Secrets** instead.



## **Acknowledgments**

- Special thanks to the **Golang**, **GORM**, and **Fiber** communities for their amazing tools and libraries.
- Inspired by various tutorials and open-source projects.

---

Feel free to contribute, report issues, or suggest improvements! 🚀


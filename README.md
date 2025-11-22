
# **Campay Mobile Money Service (Go)**


A lightweight Go service for initiating **Campay Mobile Money payments** (MTN/Orange Cameroon).
This project demonstrates how to:

* Send Mobile Money payment requests to Campay with the API_KEY loaded using `.env`

---

##  **Getting Started**

### **1: Set up your Campay Account**
Visit https://demo.campay.net/en/ and create an account. This should be very straight forward. Just make sure to verify your email address via the link they send to your address.

### **2: Getting an access token (PAT or API_KEY)**
- Click on the "Applications" tab on the left side of the dashboard.
- Click "Register Application"
- Give your application any name you'd like
- You can leave the rest of the fields empty
- Click on the application you just created
- Scroll down and copy the "Permanent Access token" and use it as the **API_KEY** in this project

### **3. Initialize the Go module**

```bash
go mod init github.com/repo-user/campay-mobile-money
```

### **4. Install dependencies**

```bash
go get github.com/joho/godotenv
go mod tidy
```

### **5. Run the project**

```bash
go run .
```

---

## **Project Structure**

```
campay-mobile-money/
│── go.mod
│── main.go
└── request_impl.go
```

* `main.go` → Entry point containing `func main()`
* `request_impl.go` → Contains payment logic (`MakePayment`)
* All files use `package main`, so they compile together automatically

---

## **Environment Variables**

Create a `.env` file in the root directory:

```
API_KEY=Token YOUR_CAMPAY_API_KEY
```

---

## **Running Payments**

The service constructs a `CollectRequest` and sends a POST request to:

```
https://demo.campay.net/api/collect/
```

It returns either:

* A successful `PaymentResponse`, **or**
* An `error`

---

#  **IMPORTANT NOTE — CAMEROON MTN/ORANGE FORMAT**

The Campay API **only supports MTN and Orange Cameroon numbers**.

- **Number must start with country code `237`**
- No `+237`, no parentheses, no spaces
- Format must be:

```
2376XXXXXXXX
2375XXXXXXXX
```

The code should  validates the phone number and if it does not have a 237 prefix, it prompts you to re-enter

---


## **Contributing**

Pull requests are welcome.
For major changes, please open an issue first to discuss what you would like to change.


Here is a clean, professional **README.md** including your commands, instructions, and a Cameroon-specific validation note.

---

# 📌 **Campay Mobile Money Service (Go)**

A lightweight Go service for initiating **Campay Mobile Money payments** (MTN/Orange Cameroon).
This project demonstrates how to:

* Send Mobile Money payment requests to Campay
* Load configuration using `.env`
* Validate numbers in the **Cameroon format (237XXXXXXXXX)**
* Organize a clean Go module

---

## 🚀 **Getting Started**

### **1. Initialize the Go module**

```bash
go mod init github.com/repouser/campay-mobile-money
```

### **2. Install dependencies**

```bash
go get github.com/joho/godotenv
go mod tidy
```

### **3. Run the project**

```bash
go run .
```

---

## 📁 **Project Structure**

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

## 📦 **Environment Variables**

Create a `.env` file in the root directory:

```
API_KEY=Token YOUR_CAMPAY_API_KEY
AMOUNT=5
FROM=237600000000
DESCRIPTION=Test Transaction
REFERENCE=1234567890
```

---

## 📡 **Running Payments**

The service constructs a `CollectRequest` and sends a POST request to:

```
https://demo.campay.net/api/collect/
```

It returns either:

* A successful `PaymentResponse`, **or**
* An `error`

---

# ⚠️ **IMPORTANT NOTE — CAMEROON MTN/ORANGE FORMAT**

The Campay API **only supports MTN and Orange Cameroon numbers**.

✔ **Number must start with country code `237`**
✔ No `+237`, no parentheses, no spaces
✔ Format must be:

```
2376XXXXXXXX
2375XXXXXXXX
```

The code should  validates the phone number and if it does not have a 237 prefix, it prompts you to re-enter

---


## ❤️ **Contributing**

Pull requests are welcome.
For major changes, please open an issue first to discuss what you would like to change.


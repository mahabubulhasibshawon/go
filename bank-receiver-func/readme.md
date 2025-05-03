
---

# 🏦 bank-receiver-func

A basic Go program that demonstrates how to use **method receivers** with a `BankAccount` struct to simulate banking operations such as deposit, withdrawal, and balance checking.

## 📁 Project Structure

```
bank-receiver-func/
├── main.go
└── README.md
```

## 🚀 Features

- Create a bank account with an owner name and initial balance
- Deposit money into the account
- Withdraw money from the account (with insufficient balance check)
- Check the current balance

## 🧠 Concepts Demonstrated

- Structs in Go
- Pointer vs. value receivers
- Encapsulation of behavior using method receivers

## 🧾 Example Output

```bash
Current Balance for John Doe is :  1000
Deposited:  500 New Balance:  1500
Withdrawn:  200 New Balance:  1300
Current Balance for  John Doe is :  1300
```

## 🛠️ How to Run

1. Make sure you have Go installed. You can download it from: [https://golang.org/dl/](https://golang.org/dl/)

2. Clone the repository or copy the `main.go` file into your workspace.

3. Run the program:

```bash
go run main.go
```

## 📌 Notes

- The `Deposite` and `Withdraw` methods use **pointer receivers** so changes affect the original `BankAccount`.
- The `CheckBalance` method uses a **value receiver** since it only reads data and does not modify the struct.


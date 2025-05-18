//account.go
package main
import (
	"errors"
)
type Operation interface{
	CheckBalance() float64
	Deposit(amount float64)(string,float64)
	Withdraw(amount float64)(string,float64)
}
type Account struct{
	AccountID string
	balance float64
}
func (account *Account) CheckBalance() float64 {
	return account.balance
}
func (account *Account)Deposit(amount float64)(string,error){
	account.balance+=amount
	return "Deposited successful",nil
} 
func (account *Account)Withdraw(amount float64)(string,error){
	if amount > account.balance{
		return "Failed",errors.New("Failed,Insufficient Balance")
	}
	account.balance-=amount
	return "Withdrawal Sucessful",nil
}
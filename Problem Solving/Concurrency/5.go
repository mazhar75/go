/*
আপনার কাজ হল প্রোগ্রামটি ইমপ্লিমেন্ট করা, নিশ্চিত করা যে অ্যাকাউন্টের Balance সর্বদা নির্ভুল থাকবে,
 এমনকি যখন একাধিক গো-রুটিন একই সাথে এটি অ্যাক্সেস করে। একটি ডেমো কোড দেয়া হল । sync.Mutex ব্যবহার করে এটা করে ফেলুন 
 */
package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (a *BankAccount) Deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.balance += amount
	fmt.Printf("Deposited %d, new balance is %d\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew %d, new balance is %d\n", amount, a.balance)
	} else {
		fmt.Printf("Not enough funds to withdraw %d, balance is %d\n", amount, a.balance)
	}
	
}

func main() {
	account := &BankAccount{balance: 0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				if rand.Intn(2) == 0 {
					account.Deposit(rand.Intn(100))
				} else {
					account.Withdraw(rand.Intn(100))
				}
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
    
	for range ticker.C {
		account.mutex.Lock()
		fmt.Printf("Current balance is %d\n", account.balance)
		account.mutex.Unlock()
	}
}
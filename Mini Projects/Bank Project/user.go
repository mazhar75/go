package main


type Adress struct{
	Village string
	Postcode int
	District string
	Division string
}

type User struct{
	Gmail string
	Name string
	Age int
	Sex string
	Address Adress
    AccountList []Account

}
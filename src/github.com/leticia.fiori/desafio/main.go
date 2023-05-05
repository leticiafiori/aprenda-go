package main

import (
	"hash"
	"time"
)


//POST /accounts
func createAccount() {
	// cria uma nova conta e cria tbm o login
}

//GET /accounts
func listAccount() []Account {
	// lista todas as contas existentes
	return []
}

//GET /accounts/{account_id}/balance
func getBalance(id int) float64 {
	// traz o saldo da conta solicitada
	return 0.0
}

//Para criação de conta precisa:
func checkBalance(balance float64){
	// Ver se o saldo é positivo
}

func encryptSecret(secret string) hash{
	//Transforma a senha em hash
}

//Login

//POST /login
func authenticateUser(id int, secret string) string {
	// autentifica o usuario retornando o token
	return "token"
}

//Tranfers

//GET /transfers
func listTransfer(id string) []Transfer {
	// traz todas as transferencias da conta logada
}

//POST /transfers
func makeTransfer() {
	// realiza uma tranferencia
}

//Para realizar a tranfericia precisa:
func getAccountOriginId(token string) {
	//busca o id da conta atraves do token
}

func checkAccountDestinationId(account_destination_id string){
	//verifica se conta de destino existe
}

func chackBalance(account string){
	// verifica se a conta de origem tem saldo suficiente para a tranferencia 
}

func updateBalance(accountOrigem, accountDestino string, amount float64){
	//retira o saldo de Origem e passa para a de Destino
}

func main() {

}

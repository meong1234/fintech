# Fintech Project

Problem Statement:

Suatu hari anda terfikir untuk membangun sebuah applikasi fintech yang dapat melakukan transaksi non tunai, sebagai Minimum Viable Productnya anda ingin applikasi ini memiliki feature berikut
 
 Transaksi : 
  - customer dapat melakukan pendaftaran 
  - sebagai owner anda dapat mendaftarkan merchant baru yang dapat melakukan transaksi 
  - customer dapat melakukan pengisian balance (Topup)
  - customer dapat melakukan pembayaran kepada merchant" yang telah terdaftar 

Meet the actor :
 -> customer: 
 -> merchant: 

-> transaksi 
    -> customer dapat topup
    -> customer dapat transaksi ke merchant yang telah terdaftar

Some Spec: 
```
GIVEN I am unregister person
WHEN i register as customer with (name, email, phone)
THEN e-wallet should record it as new customer and return my wallet-id

GIVEN I am owner of Ewallet
WHEN i register new merchant with (name, email)
THEN e-wallet should record it as new merchant and return merchant wallet-id

GIVEN I am customer of Ewallet
WHEN i topup some amount
THEN e-wallet my balance should increased

GIVEN I am customer of Ewallet
WHEN i pay merchant 
THEN my balance should decreased, and merchant balance should increased, and i get transactionsID
```

```
(entity) user_accounts
 - user_id
 - name
 - email
 - phonenumber
 - user_type (merchant | customer)

(entity) wallet
 - wallet_id
 - user_id
 - balance

(entity) transaksi
  - id
  - reference
  - date
  - description
  - type (topup | payment)
  
topup:
  - credit wallet

payment:
  - credited wallet (merchant walletid)
  - debited wallet (customer walletid)
```

```
service
  - register_customer | register_merchant =
    - add new user to user_accounts
    - add new wallet to wallet 

  - topup:
    - add new transaksi with type topup 
    - increase balance from wallet

  - payment:
    - add new transaksi with type payment 
    - decrease balance from debited wallet
    - incrase balance from credited wallet
```

```
API:
  -> :register_customer
    {
      "name": "",
      "email": "",
      "phonenumber": ""
    }
  -> :register_merchant
    {
      "name": "",
      "email": ""
    }
  -> :topup
    {
      "wallet_id": "",
      "amount": ""
    }
  -> :payment
    {
      "wallet_id": "",
      "amount": "",
      "merchant_id": ""
    }
 ```
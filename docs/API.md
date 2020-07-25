**Show All Supported Coins(Assets)**

\----

Returns all supported coins(Assets)

\* **URL**

 /coins

\* **Method:**

   `GET`

\*  **URL Params**

   **Optional:** 

​     `id=[string]`

\* **Success Response:**  

  \* **Code:** 200 <br />

​    **Content:**
```json
[
    {
        "id": "988f08b3b21e372",
        "blockchain": {
            "platform": "ethereum",
            "network": "ropsten"
        },
        "symbol": "eth",
        "decimals": 18,
        "meta": {}
    },
    {
        "id": "b21e3988f08b372",
        "blockchain": {
            "platform": "bitcoin",
            "network": "mainnet"
        },
        "symbol": "eth",
        "decimals": 8,
        "meta": {}
    }
]
```

\* **Sample Call:**
```
curl --request GET 'localhost:8080/coins'
```

---

**Show A Coin(Asset)**

\----

Returns data about a coin

\* **URL**

 /coin/{id}

\* **Method:**

   `GET`

\* **Success Response:**  

  \* **Code:** 200 <br />

​    **Content:**
```json
{
    "id": "988f08b3b21e372",
    "blockchain": {
        "platform": "ethereum",
        "network": "ropsten"
    },
    "symbol": "eth",
    "decimals": 18,
    "meta": {}
}
```

\* **Sample Call:**
```
curl --request GET 'localhost:8080/coins/988f08b3b21e372'
```

---
```
curl --request GET 'localhost:8080/coins'
```

---

**Show Wallets**

\----

Returns data about all wallets

\* **URL**

 /wallets/{id}

\* **Method:**

   `GET`

\* **Success Response:**  

  \* **Code:** 200 <br />

​    **Content:**
```json
{
    "id": "e372988fb3b2108",
    "scheme": "multiSig",
    "addresses": [
        {
            "address": "0x3a065000ab4183c6bf581dc1e55a605455fc6d61",
            "coin": {
                "id": "988f08b3b21e372",
                "blockchain": {
                    "platform": "ethereum",
                    "network": "ropsten"
                },
                "symbol": "eth",
                "decimals": 18,
                "meta": {}
            },
            "balance": "38089899992"
        },
        {
            "address": "5fc6d3a065000ab418dc1e55a605453c6bf58161",
            "coin":{
                "id": "988f08b3b21e372",
                "blockchain": {
                    "network": "bitcoin",
                    "platform": "mainnet"
                },
                "symbol": "btc",
                "decimals": 8,
                "meta": {}
            },
            "balance": "21239"
        }
    ]
}
```

\* **Sample Call:**
```
curl --request GET 'localhost:8080/wallets
```

---
**Create Wallet**

\----

Create a wallet

\* **URL**

 /wallets

\* **Method:**

   `POST`

\* **Data Params**

`scheme=[multiSig or normal]`

`password=[string]`

`meta=[object]`

\* **Success Response:**  

  \* **Code:** 200 <br />

​    **Content:**
```json
{
    "scheme": "multiSig",
    "password": "password",
    "meta": {}
}
```

\* **Sample Call:**
```
curl --request GET 'localhost:8080/wallets --data '{ "scheme": "multiSig", "password": "password", "meta":{}}'
```

---
**Transfer**

\----

Transfer coin

\* **URL**

 /wallets/{id}/transfer

\* **Method:**

   `POST`
\* **Data Params**

`to=[string]`

`amount=[string]`

`coin=[object]`

`meta=[object]`

\* **Success Response:**  

  \* **Code:** 200 <br />

​    **Content:**
```json
{
    "to": "0x3a065000ab4183c6bf581dc1e55a605455fc6d61",
    "amount": "3",
    "coin":{
        "id": "988f08b3b21e372",
        "blockchain": {
            "network": "bitcoin",
            "platform": "mainnet"
        },
        "symbol": "btc",
        "decimals": 8,
        "meta": {}
    },
    "meta": {}
}
```

\* **Sample Call:**
```
curl --request GET 'localhost:8080/wallets/3c6bf581dc1e55/transfer --data \
'{ 
    "to": "0x3a065000ab4183c6bf581dc1e55a605455fc6d61", 
    "coin":{
        "id": "988f08b3b21e372",
        "blockchain": {
            "network": "bitcoin",
            "platform": "mainnet"
        },
        "symbol": "btc",
        "decimals": 8
    },
    "balance": "21239",
    "meta": {}
}'
```
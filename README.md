## Crypto-Balance 🚀

I needed a refresher on YAML parsing, number formatting, and data structures in GoLang,
so now I have a handy binary to check my coin balances against CoinGecko :)

Ok, I was especially curious to see support for emoji...

Adjust your coins amount in balance.yaml

```yaml
xmr: 150.32
eth: 45.99
doge: 250000
```

Obtain a snapshot of your fortune with:
`$ go run ./balance.go` or
`$ go build -o ./bin/balance balance.go` and `$ ./bin/balance`

```bash
+------+-------------------+------------+------------+-----------+------------+
| COIN |      AMOUNT       | MARKET 💵  | BALANCE 💵 | MARKET 💶 | BALANCE 💶 |
+------+-------------------+------------+------------+-----------+------------+
| XMR  |        150.320007 |   163.9000 |  24,637.45 |  134.8700 |  20,273.66 |
| ETH  |         45.990002 | 1,107.9800 |  50,956.00 |  911.6400 |  41,926.32 |
| DOGE |    250,000.000000 |     0.0086 |   2,156.70 |    0.0071 |   1,774.54 |
+------+-------------------+------------+------------+-----------+------------+
|  🗓️   | 12 JAN 2021 12:16 |  TOTAL 💵  | 77,750.16  | TOTAL 💶  | 63,974.52  |
+------+-------------------+------------+------------+-----------+------------+
```

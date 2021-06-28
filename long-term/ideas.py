spot_fee = 0.002
future_fee = 0.002
denominator = 4
least_invest = 50 



# api
参数：
"BTC_CW"表示BTC当周合约，"BTC_NW"表示BTC次周合约，"BTC_CQ"表示BTC当季合约, "BTC_NQ"表示次季度合约"

https://api.hbdm.com/api/v1/contract_api_state 
json.get("data")
for coin in json:
    coin.get("symbol")

more detail:
 https://api.hbdm.com/api/v1/contract_open_interest
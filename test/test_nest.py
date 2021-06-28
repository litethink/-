
from pandas import DataFrame
import requests

resp = requests.get('https://api.huobi.pro/market/history/kline?period=1min&size=2000&symbol=btcusdt', 
                    proxies=dict(http='socks5://127.0.0.1:1080',
                                 https='socks5://127.0.0.1:1080'))                            
btcusdt = resp.json().get("data")


resp = requests.get('https://api.huobi.pro/market/history/kline?period=1min&size=2000&symbol=nestusdt', 
                    proxies=dict(http='socks5://127.0.0.1:1080',
                                 https='socks5://127.0.0.1:1080'))                                 
nestusdt = resp.json().get("data")


resp = requests.get('https://api.huobi.pro/market/history/kline?period=1min&size=2000&symbol=nestbtc', 
                    proxies=dict(http='socks5://127.0.0.1:1080',
                                 https='socks5://127.0.0.1:1080'))
nestbtc = resp.json().get("data")

usdtdf = DataFrame(nestusdt)
btcdf=DataFrame(btcusdt)
quotedf = DataFrame(nestbtc)
from matplotlib import pyplot as plt
plt.plot(quotedf.close.pct_change())
plt.plot((usdtdf.close/btcdf.close).pct_change())
plt.show()


#5min data 
from data  import quote_5min,btc_5min,nest_5min
from pandas import DataFrame
usdtdf = DataFrame(nest_5min.get("data"))
usdtdf
usdtdf["close"]
btcdf=DataFrame(btc_5min.get("data"))
btcdf
nestdf
quotedf = DataFrame(quote_5min.get("data"))
quotedf
from matplotlib import pyplot as plt
plt.plot(quotedf.close*btcdf.close)
plt.plot(usdtdf.close)
plt.show()


%hist -f main.py
plt.plot(quotedf.close)
plt.plot(quotedf.close)
plt.show()



from data_1  import btcusdt,nestusdt,nestbtc
from pandas import DataFrame
usdtdf = DataFrame(nestusdt.get("data"))
usdtdf["close"]
btcdf=DataFrame(btcusdt.get("data"))
quotedf = DataFrame(nestbtc.get("data"))
from matplotlib import pyplot as plt
plt.plot(quotedf.close*btcdf.close)
plt.plot(usdtdf.close)
plt.show()


%hist -f main.py
plt.plot(quotedf.close)
plt.plot(usdtdf.close/btcdf.close)
plt.show()


## fetchall
```sh
go run fetchall.go \
	"http://google.com" \
	"http://youtube.com" \
	"http://tmall.com" \
	"http://facebook.com" \
	"http://baidu.com" \
	"http://qq.com" \
	"http://sohu.com" \
	"http://login.tmall.com" \
	"http://taobao.com" \
	"http://360.cn" \
	"http://yahoo.com" \
	"http://jd.com" \
	"http://wikipedia.org" \
	"http://amazon.com" \
	"http://sina.com.cn" \
	"http://weibo.com" \
	"http://pages.tmall.com" \
	"http://zoom.us" \
	"http://live.com" \
	"http://netflix.com" \
	"http://reddit.com" \
	"http://xinhuanet.com" \
	"http://microsoft.com" \
	"http://office.com" \
	"http://okezone.com" \
	"http://vk.com" \
	"http://blogspot.com" \
	"http://csdn.net" \
	"http://instagram.com" \
	"http://alipay.com" \
	"http://yahoo.co.jp" \
	"http://twitch.tv" \
	"http://bongacams.com" \
	"http://bing.com" \
	"http://google.com.hk" \
	"http://microsoftonline.com" \
	"http://livejasmin.com" \
	"http://tribunnews.com" \
	"http://panda.tv" \
	"http://zhanqi.tv" \
	"http://stackoverflow.com" \
	"http://naver.com" \
	"http://amazon.co.jp" \
	"http://worldometers.info" \
	"http://twitter.com" \
	"http://tianya.cn" \
	"http://aliexpress.com" \
	"http://google.co.in" \
	"http://ebay.com" \
	"http://mama.cn"
```

```plain
0.49s    15792  http://google.com
0.54s    36433  http://live.com
0.63s       81  http://baidu.com
0.67s   157460  http://panda.tv
0.70s    39134  http://yahoo.co.jp
0.73s    93925  http://office.com
0.75s   301197  http://twitter.com
0.77s    14990  http://livejasmin.com
0.79s   175883  http://xinhuanet.com
0.83s    16041  http://google.com.hk
0.87s    92060  http://twitch.tv
0.87s   244730  http://qq.com
0.91s    16495  http://google.co.in
0.93s    91078  http://bing.com
0.99s   209386  http://sohu.com
1.01s    55868  http://aliexpress.com
1.04s    14990  http://bongacams.com
1.09s   181515  http://naver.com
1.12s   113926  http://stackoverflow.com
1.13s    92685  http://weibo.com
1.18s   319977  http://youtube.com
1.22s   178164  http://microsoft.com
1.26s    79156  http://360.cn
1.26s   168104  http://mama.cn
1.26s    69046  http://wikipedia.org
1.31s    23560  http://alipay.com
1.39s    20288  http://jd.com
1.45s   438354  http://csdn.net
1.45s   194918  http://ebay.com
1.48s   288565  http://reddit.com
1.54s     7804  http://tianya.cn
1.55s   232048  http://tribunnews.com
1.58s   129612  http://okezone.com
1.72s   529722  http://sina.com.cn
1.73s   418953  http://amazon.co.jp
1.74s    94478  http://blogspot.com
1.75s     6592  http://amazon.com
1.77s    13899  http://login.tmall.com
1.81s   173907  http://facebook.com
1.82s    80365  http://zoom.us
1.86s    43386  http://instagram.com
1.91s    67499  http://worldometers.info
1.93s    14665  http://pages.tmall.com
2.07s   231600  http://tmall.com
2.32s   157468  http://zhanqi.tv
2.47s   395528  http://yahoo.com
2.89s    11268  http://vk.com
3.01s   526380  http://netflix.com
3.74s   395936  http://taobao.com
Get "http://microsoftonline.com": dial tcp 40.84.199.233:80: i/o timeout
30.01s elapsed
```

## 応答が返って来ない(TCPハンドシェイクが始まらないなど)
予想: 返って来ないとなると、TCP仕様やライブラリのタイムアウト時間まで待つことになりそう。

結果: `Get "http://microsoftonline.com": dial tcp 40.84.199.233:80: i/o timeout`

I/Oタイムアウトになる。30秒がタイムアウト時間。おそらくこれかな。
https://golang.org/pkg/net/http/#RoundTripper

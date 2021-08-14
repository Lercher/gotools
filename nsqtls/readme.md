# nsqtls

nsqtls demonstrates how to use nsq.io with self-signed TLS
and a client TLS certificate.

## To start the server part

You'll need openssl to create X.509 certificate files for
a new CA, the server and the client. See Web Ressources how to do so.

* ca.key, ca.cert
* server.key, server.cert
* client.key, client.cert

Then start

```sh
nsqd --tls-cert server.cert --tls-key server.key --tls-required --tls-client-auth-policy require --tls-root-ca-file ca.cert
```

## Web Ressources

* <https://improveandrepeat.com/2020/09/how-to-create-self-signed-client-side-ssl-certificates-that-work/>
* <https://www.makethenmakeinstall.com/2014/05/ssl-client-authentication-step-by-step/>
* <https://deliciousbrains.com/ssl-certificate-authority-for-local-https-development/>

## The Program

The program outputs sth like that:

```log
C:\git\src\github.com\lercher\gotools\nsqtls>go run .
2021/08/14 21:08:09 This is C:\Users\Megaport\AppData\Local\Temp\go-build2039266101\b001\exe\nsqtls.exe (C) 2021 by Martin Lercher
2021/08/14 21:08:09 connecting to NSQD linux-pm81:4150 and subscribing to topic testtopic on channel consolePeeker#ephemeral ...
2021/08/14 21:08:09 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) connecting to nsqd
2021/08/14 21:08:09 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) upgrading to TLS
----------------- raw cert 0 ------------------------
00000000  30 82 04 16 30 82 01 fe  02 01 64 30 0d 06 09 2a  |0...0.....d0...*|
00000010  86 48 86 f7 0d 01 01 0b  05 00 30 49 31 0b 30 09  |.H........0I1.0.|
00000020  06 03 55 04 06 13 02 44  45 31 0b 30 09 06 03 55  |..U....DE1.0...U|
00000030  04 08 0c 02 42 59 31 0c  30 0a 06 03 55 04 07 0c  |....BY1.0...U...|
00000040  03 4d 55 43 31 0d 30 0b  06 03 55 04 0a 0c 04 6e  |.MUC1.0...U....n|
00000050  73 71 64 31 10 30 0e 06  03 55 04 03 0c 07 72 6f  |sqd1.0...U....ro|
00000060  6f 74 20 43 41 30 1e 17  0d 32 31 30 38 31 34 31  |ot CA0...2108141|
00000070  37 35 39 33 36 5a 17 0d  33 31 30 38 31 32 31 37  |75936Z..31081217|
00000080  35 39 33 36 5a 30 59 31  0b 30 09 06 03 55 04 06  |5936Z0Y1.0...U..|
00000090  13 02 44 45 31 0b 30 09  06 03 55 04 08 0c 02 42  |..DE1.0...U....B|
000000a0  59 31 0c 30 0a 06 03 55  04 07 0c 03 4d 55 43 31  |Y1.0...U....MUC1|
...
00000400  b7 9e 83 e9 8b 3e 71 35  02 ad 66 b8 78 69 50 61  |.....>q5..f.xiPa|
00000410  04 49 2f 38 58 2d e5 11  9d a8                    |.I/8X-....|

2021/08/14 21:08:09 connecting to NSQD linux-pm81:4150 and producing to to topic testtopic ...
2021/08/14 21:08:09 INF    2 (linux-pm81:4150) connecting to nsqd
2021/08/14 21:08:09 INF    2 (linux-pm81:4150) upgrading to TLS
----------------- raw cert 0 ------------------------
00000000  30 82 04 16 30 82 01 fe  02 01 64 30 0d 06 09 2a  |0...0.....d0...*|
...
00000410  04 49 2f 38 58 2d e5 11  9d a8                    |.I/8X-....|

2021/08/14 21:08:09 testtopic
message current time is: Aug 14 21:08:09
2021/08/14 21:08:10 testtopic
message current time is: Aug 14 21:08:10
2021/08/14 21:08:11 testtopic
message current time is: Aug 14 21:08:11
2021/08/14 21:08:12 testtopic
message current time is: Aug 14 21:08:12
2021/08/14 21:08:13 all messages produced, hit ^C to stop
2021/08/14 21:08:13 INF    2 stopping
2021/08/14 21:08:13 INF    2 exiting router
2021/08/14 21:08:13 stoping
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] stopping...
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) received CLOSE_WAIT from nsqd
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) beginning close
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) readLoop exiting
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) breaking out of writeLoop
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) writeLoop exiting
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) finished draining, cleanup exiting
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] (linux-pm81:4150) clean close complete
2021/08/14 21:08:13 WRN    1 [testtopic/consolePeeker#ephemeral] there are 0 connections left alive
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] stopping handlers
2021/08/14 21:08:13 INF    1 [testtopic/consolePeeker#ephemeral] rdyLoop exiting
2021/08/14 21:08:13 shutdown complete
```
# jitterg
Generate C2 dataset with custom jitter

This code generates a dataset with simulated jitter, replicating the behavior of C2 code in introducing delays between beacons. The jitter distribution is incorporated into the timestamp values of each log in the dataset. 


```
% ./jitterg -starttime 2023-10-20T15:04:22 -jitter 900s-10% -duration 4h
timestamp id src_ip src_port dest_ip dest_port bytes protocol
2023-10-20T15:18:35Z.161 fbRc1RULfdb8LQiL 10.203.219.196 49965 191.4.64.60 443 221 tls
2023-10-20T15:33:31Z.209 Jh3bgeRaQbQR20PO 10.203.219.196 5604 191.4.64.60 443 221 tls
2023-10-20T15:49:26Z.414 73VS6Z04XKO3PKJ4 10.203.219.196 3637 191.4.64.60 443 221 tls
2023-10-20T16:04:04Z.223 J052XKfbT77SO2Se 10.203.219.196 49849 191.4.64.60 443 221 tls
2023-10-20T16:19:33Z.325 37Oe7SKVKQcZ8dWP 10.203.219.196 19005 191.4.64.60 443 221 tls
2023-10-20T16:34:21Z.369 VUZKWh5YK2NgR0gU 10.203.219.196 26887 191.4.64.60 443 221 tls
2023-10-20T16:50:09Z.408 Yf5iZdfPSe2SO6P0 10.203.219.196 19858 191.4.64.60 443 221 tls
2023-10-20T17:04:01Z.344 30NbO6ULWh1260Yd 10.203.219.196 26840 191.4.64.60 443 221 tls
2023-10-20T17:20:19Z.149 dV24XMf1b0dP6dPJ 10.203.219.196 52108 191.4.64.60 443 221 tls
2023-10-20T17:34:32Z.232 iK494ZbeYOX2T8ZN 10.203.219.196 16201 191.4.64.60 443 221 tls
2023-10-20T17:50:32Z.340 beUe1KgQJKVNe1RS 10.203.219.196 64389 191.4.64.60 443 221 tls
2023-10-20T18:05:13Z.247 1cQRhORNM389RJRU 10.203.219.196 11294 191.4.64.60 443 221 tls
2023-10-20T18:19:07Z.468 34LWb9021if80dM7 10.203.219.196 61747 191.4.64.60 443 221 tls
2023-10-20T18:35:07Z.136 9XRehiSfLh9VXY2S 10.203.219.196 42648 191.4.64.60 443 221 tls
2023-10-20T18:51:00Z.388 2PJf62YJ5cfgdgU7 10.203.219.196 16176 191.4.64.60 443 221 tls
2023-10-20T19:04:38Z.267 6VZXcT0ehNKY02Z7 10.203.219.196 52820 191.4.64.60 443 221 tls

```


Random values for source ip and destnation ip will be used if not passed in the commmand arguments. 

```
% ./jitterg 
Usage of ./jitterg:
  -bytes int
    	number of bytes (default 221)
  -destination string
    	Destination IP address
  -duration duration
    	duration in hours (Default 1 week) (default 168h0m0s)
  -jitter string
    	Jitter: sleep-percentage -> (e.g 60s-10%)
  -port int
    	Destination port (default 443)
  -protocol string
    	Protocol (default "tls")
  -source string
    	Source IP address
  -starttime string
    	Start Time, Example: 2023-10-20T15:10:10
```

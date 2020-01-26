# HiddenLake

> Decentralized network. Version 1.0.0s.

### Characteristics:
1. F2F network. End to end encryption;
2. Package transfer in blockchain;
3. Symmetric algorithm: AES256-CBC;
4. Asymmetric algorithm: RSA2048-OAEP;
5. Hash function: HMAC(SHA256);

### Home page:
<img src="/images/HiddenLake1.png" alt="HomePage"/>

### Used libraries/frameworks:
1. gopeer: [github.com/number571/gopeer](https://github.com/number571/gopeer);
2. vuejs: [github.com/vuejs/vue](https://github.com/vuejs/vue);
3. bootstrap: [github.com/twbs/bootstrap](https://github.com/twbs/bootstrap);
4. go-sqlite3: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3);
5. websocket: [golang.org/x/net/websocket](https://golang.org/x/net/websocket);
6. jquery: [github.com/jquery/jquery](https://github.com/jquery/jquery);

### Signup page:
<img src="/images/HiddenLake2.png" alt="SignupPage"/>

### Modules:
1. Network (kernel): 
* connects/disconnects servers;
* creates blockchain transfer;
* sends Packages with a response to servers (TCP);
2. Intermediate (server): 
* sends API responses to the client (HTTPS, WSS);
* sends Packages with a requests to the network (TCP);
* saves information in a database;
3. Interface (client): 
* sends API requests to the server (HTTPS, WSS);
* single page application;
* native application routing;

### Account page:
<img src="/images/HiddenLake3.png" alt="AccountPage"/>

### Default configuration (config.cfg): 
> Configuration file is created when the application starts.
```json
{
	"host": {
		"http": {
			"ipv4": "localhost",
			"port": ":7545",
			"tls": {
				"crt": "tls/cert.crt",
				"key": "tls/cert.key"
			}
		},
		"tcp": {
			"ipv4": "localhost",
			"port": ":8080"
		}
	}
}
```

### Network page:
<img src="/images/HiddenLake4.png" alt="NetworkPage"/>

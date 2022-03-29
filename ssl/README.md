# SSL

## OpenSSL

Check if you already have OpenSSL installed:

```shell
openssl version
```

### ⚠️ If error

#### `Windows - Chocolatey`

```shell
choco install openssl
```

#### `Otherwise`

Please follow instructions [here](https://github.com/openssl/openssl) to install it.

## Run

### `Linux/MacOS`

```shell
chmod +x ssl.sh
./ssl.sh
```

### `Windows - Powershell`

```powershell
.\ssl.ps1
```

or, if you have a `SecurityError`:

```powershell
powershell -ExecutionPolicy unrestricted .\ssl.ps1
```

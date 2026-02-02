# transaction-service

A small transaction service API.

---

## ✅ Prerequisites

- **Docker** installed and running on your local machine.

## 🔧 Start the service

Run the included script which pulls the container image from GitHub Packages and starts it on port **80**:

```bash
# Make executable (if needed) and run
chmod +x ./run.sh
./run.sh
```

The script will pull the image, stop any previous container named `pismo_ts_container`, and run a new container that exposes the service on `http://localhost`.

> If you run the container on a different host or port, replace `localhost`/`80` in the examples below.

## 💡 Example API calls (use `curl`)

- Create an account:

```bash
curl -s -X POST http://localhost/accounts \
  -H "Content-Type: application/json" \
  -d '{"document_number":123456789}'
```

- Get an account (replace `1` with the `account_id` returned from create):

```bash
curl -s http://localhost/accounts/1
```

- Create a transaction:

```bash
curl -s -X POST http://localhost/transactions \
  -H "Content-Type: application/json" \
  -d '{"account_id":1,"operation_type_id":1,"amount":100.0}'
```

- Check Prometheus metrics:

```bash
curl -s http://localhost/metrics
```

---

If something fails, ensure Docker is running and confirm the container is up with `docker ps`.
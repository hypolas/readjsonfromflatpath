# Exemple

```json
{
  "hostname": "myhost",
  "services": [
    {
      "nom": "nginx",
      "status": "up"
    },
    {
      "nom": "httpd",
      "status": "down"
    }
  ]
}
```

The hostname can be find with:
- `services`

The status of httpd can be find with:
- `services__1__status`

The name of first service can be find with:
- `services__0__nom`

Can be work with environnement variable:
- `HYPOLAS_HEALTHCHECK_HTTP_JSON`
- `HYPOLAS_READ_JSON`

```bash
export HYPOLAS_HEALTHCHECK_HTTP_JSON=services__0__nom
export HYPOLAS_READ_JSON=services__0__nom
```

Entrypoint: 
```go
ReadJSONFromFlatPath(jpath string, jsonFile []byte)
```

If `jpath` is empty, module with try `HYPOLAS_HEALTHCHECK_HTTP_JSON` and `HYPOLAS_READ_JSON.`

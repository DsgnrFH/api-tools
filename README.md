# API Tools

A growing collection of utility endpoints — from ID generation to cryptography to networking.

> [!IMPORTANT]
> This project is a work in progress and may contain bugs or incomplete features.

The server runs on port `8080`.

---

## Table of Contents

- [Identity](#identity)
  - [UUID](#uuid)
  - [UUIDv7](#uuidv7)
- [Cryptography](#cryptography)
  - [SHA-256](#sha-256)
- [Network](#network)
  - [Subnetting](#subnetting)

---

## Identity

Base path: `/api/v1/id`

---

### UUID

**`GET /api/v1/id/uuid`**

Generates a cryptographically random UUIDv4. Uses Go's `crypto/rand` for randomness — not `math/rand`. The version and variant bits are set according to the RFC 4122 spec (version nibble `4`, variant bits `10xx`).

**Parameters**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `format` | string | No | Output format. `text` returns a plain string. Omit for JSON. |

**Examples**

```
GET /api/v1/id/uuid
GET /api/v1/id/uuid?format=text
```

**Responses**

Default (JSON):
```json
{"uuid": "550e8400-e29b-41d4-a716-446655440000"}
```

With `?format=text`:
```
550e8400-e29b-41d4-a716-446655440000
```

---

### UUIDv7

**`GET /api/v1/id/uuidv7`**

Generates a UUIDv7. Unlike v4, UUIDv7 embeds a Unix millisecond timestamp in the first 48 bits, making generated IDs monotonically sortable. The implementation handles clock drift and sub-millisecond collisions via a sequence counter, and uses `crypto/rand` for the remaining random bytes.

**Parameters**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `format` | string | No | Output format. `text` returns a plain string. Omit for JSON. |

**Examples**

```
GET /api/v1/id/uuidv7
GET /api/v1/id/uuidv7?format=text
```

**Responses**

Default (JSON):
```json
{
  "uuid": "018f4e3a-2b1c-7e4d-9a8b-1c2d3e4f5a6b",
  "type": "uuidv7"
}
```

With `?format=text`:
```
018f4e3a-2b1c-7e4d-9a8b-1c2d3e4f5a6b
```

---

## Cryptography

Base path: `/api/v1/crypto`

---

### SHA-256

**`GET /api/v1/crypto/sha256`**

Returns the SHA-256 hash of the provided input string. SHA-256 is a one-way cryptographic hash function producing a 256-bit (64 hex character) digest. The response echoes the original input alongside the hash.

`text` is required. Omitting it returns a `400 Bad Request`.

**Parameters**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `text` | string | Yes | The string to hash. |

**Examples**

```
GET /api/v1/crypto/sha256?text=Hello
```

**Responses**

Success:
```json
{
  "input":  "Hello",
  "sha256": "185f8db32921bd46d35d27c265e8ee1bc4c41d5d096f83bcb26d27ecac09e2ab"
}
```

---

## Network

Base path: `/api/v1/network`

---

### Subnetting

**`GET /api/v1/network/subnet`**

Calculates subnet details for a given IPv4 CIDR block. Returns the netmask, broadcast address, first and last usable host addresses, and total usable host count (i.e. 2^n - 2, excluding network and broadcast addresses).

If `cidr` is omitted, the server defaults to `192.168.1.0/24`. An invalid CIDR format returns a `400 Bad Request`.

**Parameters**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cidr` | string | No | An IPv4 CIDR block, e.g. `192.168.1.0/24`. Defaults to `192.168.1.0/24`. |

**Examples**

```
GET /api/v1/network/subnet?cidr=192.168.1.0/24
GET /api/v1/network/subnet
```

**Responses**

Success:
```json
{
  "cidr":         "192.168.1.0/24",
  "netmask":      "255.255.255.0",
  "broadcast":    "192.168.1.255",
  "first_usable": "192.168.1.1",
  "last_usable":  "192.168.1.254",
  "total_hosts":  254
}
```

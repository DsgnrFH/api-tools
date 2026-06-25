# Api tools
Anything potentially useful from calculator to tempreature converted or UUIDv7 generatoer - should be here

> [!IMPORTANT]
> This tool is WIP (Work in progress) and may contain bugs and be incomplete.

What's available:
<details>
<summary>ID related stuff</summary>
    <details>
    Endpoint: `/api/v1/id`
    Optional flags: `format`
    <summary>UUID</summary>
    Endpoint: `/api/v1/id/uuid`
    Generates a pseudo-random UUIDv4
    Example: `/api/v1/id/uuid?format="text"`
    </details>
    <details>
    <summary>UUIDv7</summary>
    Endpoint: `/api/v1/id/uuidv7`
    Generates a true random UUIDv7
    Example: `/api/v1/id/uuidv7?format="text"`
    </details>
</details>
<details>
<summary>Cryptography related stuff</summary>
    <details>
    Endpoint: `/api/v1/crypto`
    <summary>SHA256</summary>
    Endpoint: `/api/v1/crypto/sha256`
    Generates a hash for `text` provided as Get argument
    Example: `/api/v1/crypto/sha256?text="Hello"`
    </details>
</details>
<details>
<summary>Network related stuff</summary>
    <details>
    Endpoint: `/api/v1/network`
    <summary>Subnetting</summary>
    Endpoint: `/api/v1/network/subnet`
    Calculates anything needed for `cidr` provided as Get argument
    Example: `/api/v1/network/subnet?cidr=192.168.1.0/24`
    </details>
</details>


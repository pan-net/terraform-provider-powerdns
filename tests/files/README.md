This root CA is used in tests and should never ever be considered for use
in any environment besides developmet.

* Create root CA
```
mkdir -p rootCA/{certs,db,private}
touch rootCA/db/db
touch rootCA/db/db.attr

openssl req -x509 -sha256 -days 10000 -newkey rsa:3072 \
    -config root-csr.conf -keyout rootCA/private/rootCA.key \
    -out rootCA/rootCA.crt
```

* Sign "localhost" server certificate
```
mkdir -p localhost/

openssl req -new -config localhost-csr.conf -out localhost/server.csr \
        -keyout localhost/server.key

openssl ca -config rootCA.conf -days 10000 -create_serial \
    -in localhost/server.csr -out localhost/server.crt \
    -extensions leaf_ext -notext
```
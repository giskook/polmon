# polmon

## How to run
The default configuration file is monitor network 3

```azure
./polmon run
```

## RPC interface

```
curl http://127.0.0.1:8080/polmon/api/v1/fee
```

returns the current total fee costed, data is the total fee costed in wei

```azure
{
    "code": "200000",
    "desc": "",
    "data": "0"
}
```

## Database
Two tables are created in the database, one is the syncs table, the other is the statistics table.

```azure
sqlite3 ./fee.db

select * from syncs;
    
select * from statistics;
```

### syncs table
The syncs table is used to store the sync information of the block, the fields are as follows:

```azure
block_num INTEGER ,
tx_hash string,
fee INTEGER,
```

### statistics table
The statistics table is used to store the statistics information of the block, the fields are as follows:

```azure
block_num INTEGER ,
tx_hash string,
fee INTEGER,
total string,
```

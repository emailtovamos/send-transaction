# send-transaction

This main.go file sends a generic transaction to a local EVM network.
The account corresponding to the given private key is expected to have some funds beforehand which is given in the genesis file. 

So one can use the genesis file to initialise the network beforehand using the following:

- cd <bsc/go-ethereum repo>
- make geth
- mkdir tempdatadir
- ./build/bin/geth --datadir tempdatadir init tempdatadir/genesis.json
- ./build/bin/geth --datadir tempdatadir  console --http --http.corsdomain https://remix.ethereum.org --allow-insecure-unlock --http.api personal,eth,net,web3,debug --http.vhosts '*,localhost,host.docker.internal'  --http.addr "0.0.0.0" --rpc.allow-unprotected-txs --networkid 1337 --miner.etherbase 0x9fb29aac15b9a4b7f17c3385939b007540f4d791 --vmdebug
- personal.importRawKey("9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3", "123456")
- personal.unlockAccount("0x9fb29aac15b9a4b7f17c3385939b007540f4d791", "123456", 0)
- miner.start()

# Chia CLI
Golang client for communicating with [Chia](https://www.chia.net/) RPC interfaces.
### Add Node
```
chia-cli node create --url="<your_node_url>" --name="<node_name>"
```
`your_node_url` need include protocol, such as `https://1.1.1.1/` (fake address)

`node_name` only used for local display.

then you can run `chia-cli node list` to list all nodes you added.
### Import Private Cert
download private cert from node server. include folders as follows:
- full_node
- wallet
- farmer
- harvester

they are general locate at `/<home_dir>/.chia/<network_name>/config/ssl/`

```
# import full_node private cert
chia-cli node add --name="<node_name>" --cert-type="private_full_node" --crt-path="full_node/private_full_node.crt" --key-path="full_node/private_full_node.key"

# import wallet private cert
chia-cli node add --name="<node_name>" --cert-type="private_wallet" --crt-path="wallet/private_wallet.crt" --key-path="wallet/private_wallet.key"

# import farmer private cert
chia-cli node add --name="<node_name>" --cert-type="private_farmer" --crt-path="farmer/private_farmer.crt" --key-path="farmer/private_farmer.key"

# import harvester private cert
chia-cli node add --name="<node_name>" --cert-type="private_harvester" --crt-path="harvester/private_harvester.crt" --key-path="harvester/private_harvester.key"
```
run the command to import them, and you can delete all folders after import.

### using example
```
# get current information about the blockchain, including the peak, sync information, difficulty, mempool size, etc
chia-cli rpc full-node GetBlockchainState --node="<node_name>"
```
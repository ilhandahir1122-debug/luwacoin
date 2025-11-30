# Luwa Chain (PoA / Clique) — folder

Files:
- genesis.json  -> genesis for Clique PoA (chainId = 9882)
- init.sh       -> run `geth --datadir luwa-node init genesis.json`
- start.sh      -> start geth as PoA signer

## Quick start (assuming geth binary available as ./build/bin/geth)

1. Init chain:
   chmod +x luwa-chain/init.sh
   ./luwa-chain/init.sh

2. Create or import keystore account (must match a signer in genesis.extraData):
   # Create new
   ./build/bin/geth account new --datadir luwa-node

   # Or import private key
   ./build/bin/geth account import /path/to/privatekeyfile --datadir luwa-node

3. Put the keystore password in file: luwa-chain/password.txt

4. Start node:
   chmod +x luwa-chain/start.sh
   ./luwa-chain/start.sh

Notes:
- If your signer address is different, update genesis.json extraData (32 bytes vanity + concatenated signer addresses (no 0x) + 65 bytes zeros).
- For multiple signers add multiple 20-byte addresses in the extraData.
- For a private test network, using --nodiscover is recommended.

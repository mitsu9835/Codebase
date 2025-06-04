package functions

import (
  "io/ioutil"
  "log"
  "os"
  "path/filepath"
)

// CheckError aborts on any non-nil error.
func CheckError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// GethPathAndKey returns the HTTP RPC endpoint and the JSON keystore content
// for your node1 account.
func GethPathAndKey() (string, string) {
  // 1) RPC URL of your running geth node:
  gethPath := "http://127.0.0.1:8545"

  // 2) Path to the keystore JSON for node1’s funded account:
  keystoreFile := "./privatechain/node1/keystore/" +
    "UTC--2025-05-21T10-52-00.263979466Z--89126844cb02681888c4b00055045a2c8112f41f"

  content, err := ioutil.ReadFile(keystoreFile)
  CheckError(err)

  return gethPath, string(content)
}

// FileCrWr creates any parent directories and writes data to the given path.
func FileCrWr(path string, data []byte) {
  os.MkdirAll(filepath.Dir(path), 0755)
  err := ioutil.WriteFile(path, data, 0644)
  CheckError(err)
}

// Createtmp makes the local ./tmp/contractInfo folder structure.
func Createtmp() {
  os.MkdirAll("./tmp/contractInfo", 0755)
}

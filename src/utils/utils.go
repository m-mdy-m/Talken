package utils
import (
    "os"
)
func WriteFile(path string, data []byte,perm os.FileMode)error {
  return os.WriteFile(path,data,perm)
}

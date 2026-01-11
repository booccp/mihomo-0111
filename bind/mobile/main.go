package mobile

import (
    "github.com/metacubex/mihomo/constant"
    "github.com/metacubex/mihomo/hub/executor"
    "github.com/metacubex/mihomo/log"
    "os"
    "path/filepath"
)

var homeDir string

func SetHomeDir(dir string) bool {
    info, err := os.Stat(dir)
    if err != nil {
        log.Errorln("[Mihomo] SetHomeDir: %s : %+v", dir, err)
        return false
    }
    if !info.IsDir() {
        log.Errorln("[Mihomo] SetHomeDir: %s is not a directory", dir)
        return false
    }
    
    homeDir = dir
    constant.SetHomeDir(dir)
    constant.SetConfig(filepath.Join(dir, "config.yaml"))
    return true
}

func Start(configContent string) error {
    if homeDir == "" {
        return os.ErrInvalid
    }
    
    configFile := filepath.Join(homeDir, "config.yaml")
    err := os.WriteFile(configFile, []byte(configContent), 0644)
    if err != nil {
        return err
    }
    
    _, err = executor.Parse()
    return err
}

func Stop() {
    executor.Shutdown()
}

func GetVersion() string {
    return constant.Version
}

# How to use config global 
- First import the package: `import "interal/configs.go"` 
- Then get the config: `cfg := configs.Get()` and use it like `cfg.Server.Port` or `cfg.JWT.Secret` (cfg.<StructName>.<FieldName>)


# How to Use the Logger
- First, import the package: import "your_module_name/internal/logger"
- Then, call the desired function: logger.Info("Server has started") or logger.Errorf("Failed to process request %s: %v", requestID, err)
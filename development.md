# How to use config global 
- First import the package: `import "interal/configs.go"` 
- Then get the config: `cfg := configs.Get()` and use it like `cfg.Server.Port` or `cfg.JWT.Secret` (cfg.<StructName>.<FieldName>)
----------------------------------------------------------------------------------------------------------------------------------
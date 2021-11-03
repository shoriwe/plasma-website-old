# ObjectLoader

**`ObjectLoader`** are function callbacks that receive the master context of execution and the instance of the plasma
virtual machine internal controller.

They should be used to create valid plasma objects to load in the virtual machine.

```go
type ObjectLoader func (*Context, *Plasma) *Value
```


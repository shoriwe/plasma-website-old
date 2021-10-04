# Module

Modules are useful namespaces to isolate repeated names with different functionalities.

## Example

```ruby
module LogSystem
    prefix = "My prefix--"
    def log(message)
        println("[LOG] " + message.ToString())
    end
    
    def error(message)
        println("[ERROR] " + message.ToString())
    end
end

LogSystem.log(LogSystem.prefix + "This is a log")
```
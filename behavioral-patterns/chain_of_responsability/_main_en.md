# Chain of Responsabilbity pattern explained

## Chain of Responsability
Chain of Responsibility is a behavioral design pattern that allows you to pass requests through a chain of handlers. Upon receiving a request, each handler decides whether to process the request or pass it on to the next handler in the chain.[1]

## When to use it
* When you want to avoid direct coupling between the sender of a request and the handlers.
* When you need to apply multiple processing steps that may be dynamic or configurable.
* When the objects that make up the chain may be modified at runtime
* When you want to avoid an object having direct dependencies on all possible handlers.

## References
[1] - Megulho nos padrões de projetos, Alexander Shvets  
[2] - Go Design Patterns, Mario Castro Contreras
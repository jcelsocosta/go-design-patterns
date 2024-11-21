# Iterator pattern explained

## Iterator
The Iterator is a behavioral design pattern that allows you to iterate through elements of a collection without exposing its representations (list, stack, tree, etc.).[1]

## When to use it
* When you want your collection to have a complex data structure under the hood, but you want to hide its complexity from your clients (either for convenience or security reasons).
* When you want to reduce duplication of traversal code in your application.
* When you want your code to be able to traverse different data structures, or when the types of those structures are unknown beforehand

## References
[1] - Megulho nos padrões de projetos, Alexander Shvets  
[2] - Go Design Patterns, Mario Castro Contreras
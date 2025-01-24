# go-design-patterns
A demonstration of Design Patterns implementation using the Go programming language

## Creational
### 01 - Singleton
Objective of the Singleton pattern:
Example:
- Using the same connection to a database to make every query

### 02 - Builder
A Builder design pattern tries to:
- Abstract complex creations so that object creation is separated from the object
user objects
- Create an object step by step by filling its fields and creating the embedded
- Reuse the object creation algorithm between many objects

### 03 - Factory
A Factory design pattern tries to:
- Delegating the creation of new instances of structures to a different part of the program
- Working at the interface level instead of with concrete implementations
- Grouping families of objects to obtain a family object creator

### 04 - Abstract Factory
An Abstract Factory design pattern tries to:
- Provide a new layer of encapsulation for Factory methods that returns a common interface for all factories
- Group common factories into a super Factory (also called factory of factories)
# Design

## Goals

1. Ease of use
   * Tool must be really simple and intuitive to use, easier than writing
     mappings manually
2. Ease of development
   * It should be as easy as possible to read, modify and reason about the
     code, so that development could be more fun and take less effort,
     also potentially attracting more contributors
3. Performance of generated code
   * It should be more or less comparable to hand-written mappers to improve
     attractiveness of the tool. Also, good for reducing carbon footprint :) 
   
Note that performance of the code generator itself is generally 
*not* a goal - unless it becomes a usability issue, then see goal 1.

## Mapping

At the core there are two main object types - *requests* and *providers*.

*Request* is a simple name-type pair which represents need for some value
of a given name and type (actually name can be empty, indicating only
type matching is needed).

*Providers* are things that can fulfill requests. Examples are mapper
arguments and other mappers - either specified explicitly in a config or
created implicitly to do automatic type conversions. Providers can have
child providers, forming a provider tree. Also, providers can have
dependencies, which are requests that needs to be fulfilled.

For any given request resolution algorithm is following:
1. Traverse provider tree using breadth-first search to find first matching 
   provider
2. If matching provider doesn't have dependencies we're done
3. Otherwise try to recursively resolve all the dependencies using same 
   algorithm
4. If all dependencies are resolved we are done
5. Otherwise we consider this provider as non-matching and go back to step 1 
   to search for another provider.
6. If root request couldn't be resolved we are failed


## Examples

### Mapping 

```go
type Source struct {
	X int
	Y InnerSource
	YB int
}

type InnerSource struct {
	Z int
	A int
}

type Dest struct {
	X int
	Y InnerDest
	YA int
}

type InnerDest struct {
	Z int
	B int
}

func Mapper(src Source) Dest {
	panic(automap.Build())
} 
```

Root request: "_ Dest"

Main provider resolution sequence:
* "src Source"
* "srcX int"
* "srcY InnerSource"
* "srcYB int"
* "X int"
* "Y InnerSource"
* "YB int"
* "srcYZ int"
* "srcYA int"
* "YZ int"
* "YA int"

Root request fails to resolve through main sequence, however it
contains composite type, for which matching provider is generated:
* "_ Dest" <- "X int", "Y InnerDest", "YA int"

"Y InnerDest" also fails to resolve through main sequence, however
it also contains composite type, so another matching provider is generated:
* "Y InnerDest" <- "YZ int", "YB int"

"X int", "YA int", "YB int" and "YZ int" are resolved through main sequence

### Dependency injection

```go
func NewLogger() *Logger {}

func NewRepository(logger *Logger) *Repo {}

func NewService(repo *Repository, logger *Logger) *Service {}

func NewController(service *Service, logger *Logger) *Controller {}

func BuildController() *Controller {
	panic(automap.Build(
		automap.Provider(NewLogger),
		automap.Provider(NewRepository),
		automap.Provider(NewService),
		automap.Provider(NewController), 
	))
}
```

Root request: "_ *Controller"

Main provider resolution sequence:
* "_ *Logger"
* "_ *Repository" <- "logger *Logger"
* "_ *Service" <- "repo *Repository", "logger *Logger"
* "_ *Controller" <- "service *Service", "logger *Logger"

Full resolution chain happens successfully because unnamed providers
don't require name matching.
# Please, go and code generation

This project demonstrates an approach to make a Please project a valid go project when there's code generation.
Effectively, generated sources are symlinked back into the repo root allowing go tooling to pick them up.

This project relies on yet to be released features of Please. You will need to build and install Please from:
https://github.com/Tatskaari/please/tree/draft-codegen

## Generating code
You'll notice I've set a new config value in `.plzconfig`:

```
[build]
LinkGeneratedSources = true
```

This config value instructs Please to link sources back to the source tree. This enables `go build` and related tooling 
to find them. This will be done automatically as targets are built however, to generate all sources into the repo, 
there's a new subcommand:
```
$ plz codegen //some/wildcard/...
```

## Ignoring generated sources
The codegen subcommand can also be used to generate a .gitignore to ignore the generated srcs:
```
$ plz codegen --update_gitignore=apps/user_service/.gitignore
```

I've also created some helpers: `//apps/user_service:generate_gitignore` and `//apps/auth_service:generate_gitignore`, 
which will update the relevant `.gitignore` files.

## Bringing it together
Once all the code is generated, `go build`, intellij and other tooling should be able to find the relevant sources. You 
can open `apps/user_service` as a intellij project, set up module integrations, and you're away. 

For some projects, you may want to create a single go.mod in the repo root however this repo demonstrates how it might 
work for a monorepo. Where one module depends on another, a `replace` directive has been added to the go.mod withe a 
relative path to the other module. This allows us to depend on our local modules that aren't published anywhere.  

# Please, go and code generation

This project demonstrates an approach to make a Please project a valid go project when there's code generation.
Effectively, generated sources are symlinked back into the repo root allowing go tooling to pick them up. 

## Quick setup
To generated sources and link them to your source tree run:
```
$ plz generate
```

Now open up the project in your IDE and enable go modules. 

In intellij, you can open up the root folder as a project. The relevant setting is `Preferences -> Languages & 
Frameworks -> Go -> Go modules -> Enable Go modules integration`. 

Unfortunately, VSCode isn't as monorepo aware and doesn't parse the `go.mod` unless you open up that folder. If you open 
`apps/user_service`, vscode will work as intended. Maybe somebody more versed in vscode can find a way around this.   

At this point, both intellij and vscode will provide widgets next to tests and main methods. You should be able to use 
these to build, run, and debug any test or main function from within the IDE. 

Go should also work from the command line:
```
$ (cd apps/user_service && go run main.go)
Success! User service compiled and ran!

$ (cd apps/auth_service && go run main.go)
Success! Auth service compiled and ran!

$ (cd apps/user_service && go test -run '' ./dao)
ok      apps/user_service/dao   0.082s
```

## Generating code
You'll notice I've set a new config value in `.plzconfig`:

```
[build]
LinkGeneratedSources = true
```

This config value instructs Please to link sources back to the source tree. Under the hood, Please figures out what to 
link based on the `codegen` label. The outputs of any rule with this label will be linked back to the source tree. This 
enables `go build` and related tooling to find these generated sources. 

## Working with generated sources
The `generate` subcommand can be used to build any codegen targets, but it can also be used to generate a .gitignore to 
avoid the generated srcs getting checked in to source control:
```
$ plz generate --update_gitignore=apps/user_service/.gitignore
```

This will build all the generated targets under `//apps/user_service/...` and write the generated paths to that 
`.gitignore`. Please is careful not to clobber entries you may have added to this fie. Please writes its changes under a 
`DO NOT EDIT` comment, leaving all your entries above that comment alone.  

I've also created some helpers: `//apps/user_service:generate_gitignore` and `//apps/auth_service:generate_gitignore`, 
which will invoke please to update the relevant `.gitignore` files. 

## Bringing it together
Once all the code is generated, `go build`, IDEs, and go analysis/linting tools will be able to find these generated 
sources. While it's probably possible to achieve a lot of what I've talked about here through other means, this should 
facilitate a more standard developer workflow, and is considerably easier to set up. 

NB: For some projects, you may want to create a single go.mod in the repo root. For this example, I wanted to 
demonstrate how it might work for a monorepo. There's nothing stopping you from doing it the other way. To get this to 
work, I've added a replace directive where one module depends on another. This allows us to depend on our local modules 
that aren't published anywhere. 

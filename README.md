# BMO

#### What is BMO?
Use bmo to configure and run and compile complex build systems with additonal live reloading during developement. Similar to [Air](https://github.com/cosmtrek/air) but meant to be more compatible to frameworks such as Django, and Templ with live reloading

Example Scenario: Using golang, [html](https://htmx.org), [templ](https://templ.guide) while also requiring tailwind and typescript to be compiled before execution.


Feature list
- [ ] Build step at execution
    - [ ] Compile javascript using esbuild
    - [ ] Compile golang
    - [ ] Compile tailwind
    - [ ] Generate templ files
    - [ ] Add support for more general project build steps
- [ ] Automatically run golang project in a temporary directory with required asset files
- [ ] Proxy web server and inject hot reloading javascript

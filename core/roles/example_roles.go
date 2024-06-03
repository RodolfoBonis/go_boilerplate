package roles

import "github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/entities"

var ExampleRoles = entities.Roles{
	Search: "search-examples",
	Insert: "add-example",
	Detail: "example-details",
	Update: "update-example",
	Delete: "delete-example",
}

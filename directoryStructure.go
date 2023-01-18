package main

type AnsibleTemplateFile struct {
	Name    string
	Content string
	Type    string
	Path    string
}

var mainDirectoryStructure = [...]AnsibleTemplateFile{
	// Inventory file for production servers
	{
		Name:    "production",
		Path:    "/",
		Type:    "file",
		Content: "",
	},
	// Inventory file for staging/non-production servers
	{
		Name:    "staging",
		Path:    "/",
		Type:    "file",
		Content: "",
	},
	// group*.yml contains variables assigned to specific groups
	{
		Name:    "group1.yml",
		Path:    "/group_vars",
		Type:    "file",
		Content: "",
	},
	{
		Name:    "group2.yml",
		Path:    "/group_vars",
		Type:    "file",
		Content: "",
	},
	// hostname*.yml contains assigned to specific systems
	{
		Name:    "hostname1.yml",
		Path:    "/host_vars",
		Type:    "file",
		Content: "",
	},
	{
		Name:    "hostname2.yml",
		Path:    "/host_vars",
		Type:    "file",
		Content: "",
	},
	// Directory to put custom modules (optional)
	{
		Name:    "library",
		Path:    "/",
		Type:    "directory",
		Content: "",
	},
	// Directory to put custom module_utils to support modules (optional)
	{
		Name:    "module_utils",
		Path:    "/",
		Type:    "directory",
		Content: "",
	},
	// Directory to put custom filter plugins
	{
		Name:    "filter_plugins",
		Path:    "/",
		Type:    "directory",
		Content: "",
	},
	// Master playbook
	{
		Name:    "site.yml",
		Path:    "/",
		Type:    "file",
		Content: "",
	},
	// Playbook for webserver tier (example)
	{
		Name:    "webservers.yml",
		Path:    "/",
		Type:    "file",
		Content: "",
	},
	// Playbook for dbserver tier (example)
	{
		Name:    "dbservers.yml",
		Path:    "/",
		Type:    "file",
		Content: "",
	},
}

var roleDirectoryStructure = [...]AnsibleTemplateFile{
	// Tasks
	{
		Name:    "/tasks/main.yml",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Handlers file
	{
		Name:    "/handlers/main.yml",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Template file to use with the template resource
	{
		Name:    "/templates/ntp.conf.j2",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Static files that can be used with the copy resource
	{
		Name:    "files/bar.txt",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Variables associated with the role
	{
		Name:    "/vars/main.txt",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Default lower priority variables for the role
	{
		Name:    "/defaults/main.txt",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Role dependencies
	{
		Name:    "/meta/main.txt",
		Path:    "/roles",
		Type:    "file",
		Content: "",
	},
	// Custom modules for roles
	{
		Name:    "library",
		Path:    "/roles",
		Type:    "directory",
		Content: "",
	},
	// Custom module_utils for roles
	{
		Name:    "module_utils",
		Path:    "/roles",
		Type:    "directory",
		Content: "",
	},
	// Custom lookup_plugins for roles
	{
		Name:    "lookup_plugins",
		Path:    "/roles",
		Type:    "directory",
		Content: "",
	},
}

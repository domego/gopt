# gopt
Golang 项目代码自动生成工具

* Types
* Proejct
* ORM
* Controller
* Errors
* Arugments

```
Usage: gopt <command> [options]

Commands:

	gen_gin_server	generate gin server code from template
	gen_orm	generate database orm code from template
	gen_gin_api	generate gin controller code from template
	gen_js_api	generate vue-resource code from api template
	gen_api_doc	generate api doc from api template
	version	show version
	help	show help
	gen_types	generate structs code from yaml config file
	gen_errors	generate error_msg.yaml and error constants from yaml config file

Options:

	-types	types config file, default: types.yaml
	-errors	errors config file, default: errors.yaml
	-orm	database config file, default: db.yaml
	-api	gin controller config file, default: api.yaml

```

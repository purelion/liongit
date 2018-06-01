package routers

import (
	"github.com/astaxie/beego"
	"managerSystem/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/public/index", &controllers.MainController{}, "*:Index")
	beego.Router("/public/login", &controllers.MainController{}, "*:Login")
	beego.Router("/public/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/public/changepwd", &controllers.MainController{}, "*:Changepwd")

	beego.Router("/rbac/user/AddUser", &controllers.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &controllers.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/DelUser", &controllers.UserController{}, "*:DelUser")
	beego.Router("/rbac/user/index", &controllers.UserController{}, "*:Index")

	beego.Router("/rbac/node/AddAndEdit", &controllers.NodeController{}, "*:AddAndEdit")
	beego.Router("/rbac/node/DelNode", &controllers.NodeController{}, "*:DelNode")
	beego.Router("/rbac/node/index", &controllers.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &controllers.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &controllers.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &controllers.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group/index", &controllers.GroupController{}, "*:Index")

	beego.Router("/rbac/role/AddAndEdit", &controllers.RoleController{}, "*:AddAndEdit")
	beego.Router("/rbac/role/DelRole", &controllers.RoleController{}, "*:DelRole")
	beego.Router("/rbac/role/AccessToNode", &controllers.RoleController{}, "*:AccessToNode")
	beego.Router("/rbac/role/AddAccess", &controllers.RoleController{}, "*:AddAccess")
	beego.Router("/rbac/role/RoleToUserList", &controllers.RoleController{}, "*:RoleToUserList")
	beego.Router("/rbac/role/AddRoleToUser", &controllers.RoleController{}, "*:AddRoleToUser")
	beego.Router("/rbac/role/Getlist", &controllers.RoleController{}, "*:Getlist")
	beego.Router("/rbac/role/index", &controllers.RoleController{}, "*:Index")
}

package controllerfactory

import (
	controller "Go_OOP/Controller"
	"fmt"
)

type ControllerFactory struct {
	ControllerMap map[string]any
}

var controllerFactoryInstance *ControllerFactory

func (c *ControllerFactory) NewControllerFactory() {
	c.ControllerMap = make(map[string]any)

	c.ControllerMap["AuthController"] = &controller.AuthController{}
	c.ControllerMap["FileManageController"] = &controller.FileManageController{}
	c.ControllerMap["UserController"] = &controller.UserController{}
}

func (c *ControllerFactory) GetIntance() *ControllerFactory {
	if controllerFactoryInstance == nil {
		controllerFactoryInstance = &ControllerFactory{}
		controllerFactoryInstance.NewControllerFactory()
	}
	return controllerFactoryInstance
}

func (c *ControllerFactory) GetController(name string) any {
	var controller any
	var ok bool

	controller = nil

	controller, ok = c.ControllerMap[name]
	if !ok {
		fmt.Printf("Controller Not Found")
	}

	return controller
}

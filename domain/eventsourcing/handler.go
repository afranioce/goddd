package eventsourcing

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Command interface {
	Context() *gin.Context
	CommandType() string
	Body() interface{}
}

type BaseCommand struct {
	context     *gin.Context
	commandType string
	created     time.Time
	body        interface{}
}

func (b BaseCommand) Context() *gin.Context {
	return b.context
}

func (b BaseCommand) CommandType() string {
	return b.commandType
}

func (b BaseCommand) Body() interface{} {
	return b.body
}

func commandName(command interface{}) string {
	rawType := reflect.TypeOf(command)
	return rawType.String()
}

func CreateCommand(body interface{}, ctx *gin.Context) BaseCommand {
	return BaseCommand{ctx, commandName(body), time.Now(), body}
}

// CommandHandle defines the contract to handle commands
type CommandHandle func(cmd Command) error

// CommandHandlerRegister stores the handlers for commands
type CommandHandlerRegister interface {
	Add(command interface{}, handler CommandHandle)
	Get(command interface{}) (CommandHandle, error)
	// Handlers() []string
}

// CommandRegister contains a registry of command-handler style
type CommandRegister struct {
	sync.RWMutex
	registry map[string]CommandHandle
}

// NewCommandRegister creates a new CommandHandler
func NewCommandRegister() *CommandRegister {
	return &CommandRegister{
		registry: make(map[string]CommandHandle),
	}
}

// Add a new command with its handler
func (c *CommandRegister) Add(command interface{}, handler CommandHandle) {
	c.Lock()
	defer c.Unlock()

	name := commandName(command)
	c.registry[name] = handler
}

// Get the handler for a command
func (c *CommandRegister) Get(command interface{}) (CommandHandle, error) {
	name := commandName(command)

	handler, ok := c.registry[name]
	if !ok {
		return nil, fmt.Errorf("can't find %s in registry", name)
	}
	return handler, nil
}

func (d *CommandRegister) DispatchCommands(commands ...Command) {
	for _, cmd := range commands {
		h, _ := d.Get(cmd.Body())

		if ok := h(cmd); ok != nil {
			fmt.Errorf("can't execute %s", cmd.CommandType())
		}
	}
}

 +package cmd  
 3   +  
 4   +import (  
 5   +	  
 6   +	"fmt"  
 7   +	  
 8   +	"github.com/mitchellh/cli"  
 9   +	  
 10   +)  
 11   +  
 12   +type StatusCommand struct {  
 13   +	Ui   cli.Ui  
 14   +}  
 15   +  
 16   +func (c *StatusCommand) Help() string {  
 17   +	helpText := `  
 18   +Usage: m-apiserver status  
 19   +  
 20   +	Queries and Prints the current status ,  
 21   +	port details and -bind ip address of  m-apiserver  
 22   +`  
 23   +	return (helpText)  
 24   +}  
 25   +  
 26   +func (c *StatusCommand) Run(_ []string) int {  
 27   +// Initialise hard code stuff  
 28   +	Status 		:= "running"  
 29   +	NodeName	:= "m-apiserver"  
 30   +	BindAddr	:= "172.28.128.6"  
 31   +	Ports		:= "5656"  
 32   +  
 33   +  
 34   +	//format the output  
 35   +	out := fmt.Sprintf("Name          IP            Ports       Status       Data     Log \n%-14s%-14s%d\t%-14s",  
 36   +		NodeName,  
 37   +		BindAddr,  
 38   +		Ports,  
 39   +		Status)  
 40   +	c.Ui.Output(out)  
 41   +  
 42   +	return 0  
 43   +}  
 44   +  
 45   +func (c *StatusCommand) Synopsis() string {  
 46   +	return "Prints the m-apiserver Status"

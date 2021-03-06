package api

import "fmt"

//Config represents the options that primitive accepts
type Config struct {
	Input       string `json:"input,omitempty"`
	Output      string `json:"output,omitempty"`
	Number      int    `json:"number,omitempty"`
	Mode        int    `json:"mode,omitempty"`        //0=combo, 1=triangle, 2=rect, 3=ellipse, 4=circle, 5=rotatedrect, 6=beziers, 7=rotatedellipse, 8=polygon
	Rep         int    `json:"rep,omitempty"`         //add N extra shapes each iteration with reduced search (mostly good for beziers)
	Nth         int    `json:"nth,omitempty"`         //save every nth frame, %d must be in output path
	InSize      int    `json:"in_size,omitempty"`     //size to resize large input images before processing
	OutSize     int    `json:"out_size,omitempty"`    //output image size
	Alpha       int    `json:"alpha,omitempty"`       //color alpha, 0=algos choice
	Background  string `json:"background,omitempty"`  //avg, 	starting background color (hex)
	Workers     int    `json:"workers,omitempty"`     // number of parallel workers, 0 = all cores
	Verbose     string `json:"verbose,omitempty"`     //verbose: off
	VeryVerbose string `json:"veryverbose,omitempty"` //very verbose: off

}

//CreateDefault creats a default config, missing only input and output
func CreateDefault() *Config {
	return &Config{
		Input:       "",
		Output:      "",
		Number:      100,
		Mode:        0,
		Rep:         0,
		Nth:         1,
		InSize:      256,
		OutSize:     1024,
		Alpha:       0,
		Background:  "avg",
		Workers:     0,
		Verbose:     "off",
		VeryVerbose: "off",
	}
}

//Verify returns true if config is runnable
func (c *Config) Verify() bool {
	return c.Input != "" && c.Output != "" && c.Number != 0
}

//CommandForm returns the config in the form expected by the 'primitive' command
func (c *Config) CommandForm() string {
	return fmt.Sprintf(" -i %s -o %s -n %d -m %d -rep %d -nth %d -r %d -s %d -a %d -bg %s -j %d -v %s -vv %s",
		c.Input,
		c.Output,
		c.Number,
		c.Mode,
		c.Rep,
		c.Nth,
		c.InSize,
		c.OutSize,
		c.Alpha,
		c.Background,
		c.Workers,
		c.Verbose,
		c.VeryVerbose)
}

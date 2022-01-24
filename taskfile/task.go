package taskfile

// Tasks represents a group of tasks
type Tasks map[string]*Task

// Task represents a task
type Task struct {
	Task          string          `json:"task,omitempty"`
	Cmds          []*Cmd          `json:"cmds,omitempty"`
	Deps          []*Dep          `json:"deps,omitempty"`
	Label         string          `json:"label,omitempty"`
	Desc          string          `json:"desc,omitempty"`
	Summary       string          `json:"summary,omitempty"`
	Sources       []string        `json:"sources,omitempty"`
	Generates     []string        `json:"generates,omitempty"`
	Status        []string        `json:"status,omitempty"`
	Preconditions []*Precondition `json:"preconditions,omitempty"`
	Dir           string          `json:"dir,omitempty"`
	Vars          *Vars           `json:"vars,omitempty"`
	Env           *Vars           `json:"env,omitempty"`
	Silent        bool            `json:"silent,omitempty"`
	Interactive   bool            `json:"interactive,omitempty"`
	Method        string          `json:"method,omitempty"`
	Prefix        string          `json:"prefix,omitempty"`
	IgnoreError   bool            `json:"ignoreError,omitempty"`
	Run           string          `json:"run,omitempty"`
}

func (t *Task) Name() string {
	if t.Label != "" {
		return t.Label
	}
	return t.Task
}

func (t *Task) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var cmd Cmd
	if err := unmarshal(&cmd); err == nil && cmd.Cmd != "" {
		t.Cmds = append(t.Cmds, &cmd)
		return nil
	}

	var cmds []*Cmd
	if err := unmarshal(&cmds); err == nil && len(cmds) > 0 {
		t.Cmds = cmds
		return nil
	}

	var task struct {
		Cmds          []*Cmd
		Deps          []*Dep
		Label         string
		Desc          string
		Summary       string
		Sources       []string
		Generates     []string
		Status        []string
		Preconditions []*Precondition
		Dir           string
		Vars          *Vars
		Env           *Vars
		Silent        bool
		Interactive   bool
		Method        string
		Prefix        string
		IgnoreError   bool `yaml:"ignore_error"`
		Run           string
	}
	if err := unmarshal(&task); err != nil {
		return err
	}
	t.Cmds = task.Cmds
	t.Deps = task.Deps
	t.Label = task.Label
	t.Desc = task.Desc
	t.Summary = task.Summary
	t.Sources = task.Sources
	t.Generates = task.Generates
	t.Status = task.Status
	t.Preconditions = task.Preconditions
	t.Dir = task.Dir
	t.Vars = task.Vars
	t.Env = task.Env
	t.Silent = task.Silent
	t.Interactive = task.Interactive
	t.Method = task.Method
	t.Prefix = task.Prefix
	t.IgnoreError = task.IgnoreError
	t.Run = task.Run
	return nil
}

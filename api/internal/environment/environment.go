package environment

// EnvVars Environent Variables
type EnvVars struct {
	WorkerNumber int `default:"2"`
  Version      string
  QMKPath      string `default:"/qmk_firmare"`
}

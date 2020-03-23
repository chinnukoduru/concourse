package integration_test

import (
	"github.com/concourse/concourse/worker/workercmd"
	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
)

type WorkerRunnerSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *WorkerRunnerSuite) TestWorkDirIsCreated() {
	// instantiate Workercmd.Runner
	// Set the WorkDir flag to something... a tmpdir
	// Check if the workerDir is created
	// if true = passed!
	// not true = failed!

	var wrkcmd workercmd.WorkerCommand

	parser := flags.NewParser(&wrkcmd, flags.HelpFlag|flags.PassDoubleDash)
	os.Args = []string{""}

	_, err := parser.Parse()
	s.NoError(err)

	wrkcmd.WorkDir = "somedir"
	wrkcmd.Garden.UseContainerd = true

	_, err = wrkcmd.Runner([]string{})
	s.NoError(err)

	_, err = os.Stat("somedir")
	s.Equal(!os.IsNotExist(err), true)
	s.NoError(err)
}

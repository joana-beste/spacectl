package costestimator

type Params struct {
	PlanID                string
	Storage               uint64
	Stages                int
	StagesOnDemand        int
	Scaling               uint64
	BackupIntervalMinutes uint64
}

type Estimator interface {
	Estimate(params Params) (*Estimation, error)
}

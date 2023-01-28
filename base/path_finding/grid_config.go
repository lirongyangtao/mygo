package path_finding

type PathFindingConfigOptions func(cfg *PathFindingConfig)
type PathFindingConfig struct {
	weight           float64          //权重
	allowDiagonal    bool             //允许对角线行走
	DiagonalMovement DiagonalMovement //对角线行走规则
	Heuristic        Heuristic        //估算函数
	DontCrossCorners bool             //是否跨越障碍物
}

func (cfg *PathFindingConfig) check() {
	if cfg.weight == 0 {
		cfg.weight = 1
	}
	if cfg.Heuristic == nil {
		cfg.Heuristic = manhattan
	}
	if cfg.DiagonalMovement == DiagonalMovementNone {
		if !cfg.allowDiagonal {
			cfg.DiagonalMovement = DiagonalMovementNever
		} else {
			if cfg.DontCrossCorners {
				cfg.DiagonalMovement = DiagonalMovementOnlyWhenNoObstacles
			} else {
				cfg.DiagonalMovement = DiagonalMovementIfAtMostOneObstacle
			}
		}
	}
	if cfg.DiagonalMovement == DiagonalMovementNever {
		cfg.Heuristic = manhattan
	} else {
		cfg.Heuristic = octile
	}

}

func WithWeight(weight float64) PathFindingConfigOptions {
	return func(cfg *PathFindingConfig) {
		cfg.weight = weight
	}
}
func WithAllowDiagonal(allowDiagonal bool) PathFindingConfigOptions {
	return func(cfg *PathFindingConfig) {
		cfg.allowDiagonal = allowDiagonal
	}
}

func WithDiagonalMovement(DiagonalMovement DiagonalMovement) PathFindingConfigOptions {
	return func(cfg *PathFindingConfig) {
		cfg.DiagonalMovement = DiagonalMovement
	}
}
func WithHeuristic(Heuristic Heuristic) PathFindingConfigOptions {
	return func(cfg *PathFindingConfig) {
		cfg.Heuristic = Heuristic
	}
}
func WithDontCrossCorners(DontCrossCorners bool) PathFindingConfigOptions {
	return func(cfg *PathFindingConfig) {
		cfg.DontCrossCorners = DontCrossCorners
	}
}

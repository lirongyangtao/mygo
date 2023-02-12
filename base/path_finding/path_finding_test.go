package path_finding

import "testing"

//func TestPathFindingAStar(b *testing.B)   {}
//func TestPathFindingIdaStar(b *testing.B)  {}
//func TestPathFindingDijkstra(b *testing.B)    {}
//func TestPathFindingBestFirst(b *testing.B)    {}
//func TestPathFindingBreadthFirst(b *testing.B)  {}
//func TestPathFindingBiAStar(b *testing.B)        {}
//func TestPathFindingBiBreadthFirst(b *testing.B) {}
//func TestPathFindingBiBestFirst(b *testing.B)    {}
//func TestPathPathFindingBiDijkstra(b *testing.B) {}

// ===========================================================性能测试==================================
// go test -bench="^BenchmarkPathFinding" .
func BenchmarkPathFindingAStar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(AStar, 1000)
	}
}

func BenchmarkPathFindingDijkstra(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(Dijkstra, 1000)
	}
}
func BenchmarkPathFindingBestFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(BestFirst, 1000)
	}
}
func BenchmarkPathFindingBreadthFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(BreadthFirst, 1000)
	}
}
func BenchmarkPathFindingBiAStar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(BiAStar, 1000)
	}
}
func BenchmarkPathFindingBiBreadthFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(BiBreadthFirst, 1000)
	}
}
func BenchmarkPathFindingBiBestFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(BiBestFirst, 1000)
	}
}
func BenchmarkPathPathFindingBiDijkstra(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(BiDijkstra, 1000)
	}
}
func BenchmarkPathPathFindingJumpPoint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PathFinding(JumpPoint, 1000)
	}
}

//性能太差，可以自行测试
//func BenchmarkPathFindingIdaStar(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		PathFinding(IdAStar, 1000)
//	}
//}

// =======================================================地图数据==============================================

func PathFinding(cmd PathFindingType, count int) {
	for i := 0; i < count; i++ {
		m := GetMap(i / 3)
		grid, startNode, endNode := GridFromString(m)
		grid.PathFindingRoute(cmd)(startNode.X, startNode.Y, endNode.X, endNode.Y)
	}
}
func GetMap(num int) string {
	switch num {
	case 1:
		return map1
	case 2:
		return map2
	case 3:
		return map3
	}
	return map1
}

var (
	map1 = `s000000000000000000000001011111110000000000000000000000000000000
0000000000000000000111100000000011100000000000000000000000000000
0000000000000000110000000000000000110000000000000000000000000000
0000000000000011000000000000000000000000000000000000000000000000
0000000000000011011101100000000000000000000000000000000000000000
0000000000001110000000110000000000000000000000000000000000000000
0000000000010100000000001000000000000000000000000000000000000000
0000000000100100000000001000000000000000000000000000000000000000
0000000000101000000000001000000000000000000000000000000000000000
0000000000111000000000001000000000000000000000000000000000000000
0000000000110000000000011000000000000000000000000000000000000000
0000000000000000000001100000000000000000000000000000000000000000
0000000000000000000011000000000000000000000000000000000000000000
0000000000000000001100000000000000000000000000000000000000000000
0000000000000000010000000000000000000000000000000000000000000000
0000000000000001100000000000000100000000000000000000000000000000
0000000000000010000000000000000000000000000000000000000000000000
0000000000000100000000000000000000000000000000000000000000000000
0000000000111000000000000000000000000000000000000000000000000000
0000000001100000000000000000000000000000000000000000000000000000
0000000111000000000000000000000000000000000000000000000000000000
0000001100000000000000000000000000000000000000000000000000000000
0000110000000000000000000000000000000000000000000000000000000000
0001100000000000000000000000000000000000000000000000000000000000
0110000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
00000000000000000000000000000000000000000000000000e0000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000`
	map1Res = `s000000000000000000000001011111110000000000000000000000000000000
0wwwwwwwwwwwwwwwww0111100000000011100000000000000000000000000000
000000000000000011w000000000000000110000000000000000000000000000
0000000000000011000wwww00000000000000000000000000000000000000000
00000000000000110111011w0000000000000000000000000000000000000000
000000000000111000000011w000000000000000000000000000000000000000
0000000000010100000000001w00000000000000000000000000000000000000
00000000001001000000000010w0000000000000000000000000000000000000
000000000010100000000000100w000000000000000000000000000000000000
0000000000111000000000001000w00000000000000000000000000000000000
00000000001100000000000110000w0000000000000000000000000000000000
000000000000000000000110000000w000000000000000000000000000000000
0000000000000000000011000000000w00000000000000000000000000000000
00000000000000000011000000000000w0000000000000000000000000000000
000000000000000001000000000000000w000000000000000000000000000000
0000000000000001100000000000000100w00000000000000000000000000000
00000000000000100000000000000000000w0000000000000000000000000000
000000000000010000000000000000000000w000000000000000000000000000
0000000000111000000000000000000000000w00000000000000000000000000
00000000011000000000000000000000000000w0000000000000000000000000
000000011100000000000000000000000000000w000000000000000000000000
0000001100000000000000000000000000000000w00000000000000000000000
00001100000000000000000000000000000000000w0000000000000000000000
000110000000000000000000000000000000000000w000000000000000000000
0110000000000000000000000000000000000000000w00000000000000000000
00000000000000000000000000000000000000000000w0000000000000000000
000000000000000000000000000000000000000000000wwwwwe0000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000`

	map2 = `s000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000010000000000000000000000000000000000000000000000000
0000000000000010000000000000000000000000000000000000000000000000
0000000000000010000000000000000000000000000000000000000000000000
0000000000000010000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000010000000000000000000000000000000000000000000000000
0000000000100000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
00000000000000000000000000000000000000000000000000e0000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000`
	map2Result = `s000000000000000000000000000000000000000000000000000000000000000
0w00000000000000000000000000000000000000000000000000000000000000
00w0000000000000000000000000000000000000000000000000000000000000
000w000000000000000000000000000000000000000000000000000000000000
0000w00000000000000000000000000000000000000000000000000000000000
00000w0000000000000000000000000000000000000000000000000000000000
000000w000000010000000000000000000000000000000000000000000000000
0000000w00000010000000000000000000000000000000000000000000000000
00000000w0000010000000000000000000000000000000000000000000000000
000000000w000010000000000000000000000000000000000000000000000000
0000000000w00000000000000000000000000000000000000000000000000000
00000000000w0010000000000000000000000000000000000000000000000000
000000000010w000000000000000000000000000000000000000000000000000
0000000000000w00000000000000000000000000000000000000000000000000
00000000000000w0000000000000000000000000000000000000000000000000
000000000000000w000000000000000000000000000000000000000000000000
0000000000000000w00000000000000000000000000000000000000000000000
00000000000000000w0000000000000000000000000000000000000000000000
000000000000000000w000000000000000000000000000000000000000000000
0000000000000000000w00000000000000000000000000000000000000000000
00000000000000000000w0000000000000000000000000000000000000000000
000000000000000000000w000000000000000000000000000000000000000000
0000000000000000000000w00000000000000000000000000000000000000000
00000000000000000000000w0000000000000000000000000000000000000000
000000000000000000000000w000000000000000000000000000000000000000
0000000000000000000000000w00000000000000000000000000000000000000
00000000000000000000000000wwwwwwwwwwwwwwwwwwwwwwwwe0000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000`
	map3 = `0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
000s000000000001000000000000000000000000000000000000000000000000
0000000000000001000000001000000000000000000000000000000000000000
0000000000000001000000000000000e00000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000001000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000011000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000`
	map3Result = `0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
000s000000000001000000000000000000000000000000000000000000000000
0000w00000000001000000001000000000000000000000000000000000000000
00000wwwwwwwwww1wwwwwwwwwwwwwwwe00000000000000000000000000000000
000000000000000w000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000001000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000011000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000001000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000
0000000000000000000000000000000000000000000000000000000000000000`
)

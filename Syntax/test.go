package main

import (
    "bufio"
    "container/list"
    "fmt"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    var H, W int
    fmt.Fscan(in, &H, &W)

    grid := make([]string, H)
    for i := 0; i < H; i++ {
        fmt.Fscan(in, &grid[i])
    }

    const INF = 1_000_000_000
    dist := make([][]int, H)
    for i := range dist {
        dist[i] = make([]int, W)
        for j := range dist[i] {
            dist[i][j] = INF
        }
    }

    // BFS queue
    q := list.New()
    // seed with all exits
    for i := 0; i < H; i++ {
        for j := 0; j < W; j++ {
            if grid[i][j] == 'E' {
                dist[i][j] = 0
                q.PushBack([2]int{i, j})
            }
        }
    }

    dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for q.Len() > 0 {
        front := q.Remove(q.Front()).([2]int)
        i, j := front[0], front[1]
        for _, d := range dirs {
            ni, nj := i+d[0], j+d[1]
            if ni < 0 || ni >= H || nj < 0 || nj >= W {
                continue
            }
            if grid[ni][nj] == '#' {
                continue
            }
            if dist[ni][nj] > dist[i][j]+1 {
                dist[ni][nj] = dist[i][j] + 1
                q.PushBack([2]int{ni, nj})
            }
        }
    }
    out := make([][]rune, H)
    for i := 0; i < H; i++ {
        out[i] = []rune(grid[i])
    }

    for i := 0; i < H; i++ {
        for j := 0; j < W; j++ {
            if grid[i][j] != '.' {
                continue
            }
            for _, d := range dirs {
                ni, nj := i+d[0], j+d[1]
                if ni < 0 || ni >= H || nj < 0 || nj >= W {
                    continue
                }
                if grid[ni][nj] == '#' {
                    continue
                }
                if dist[ni][nj] == dist[i][j]-1 {
                    switch {
                    case ni == i-1 && nj == j:
                        out[i][j] = '^'
                    case ni == i+1 && nj == j:
                        out[i][j] = 'v'
                    case ni == i && nj == j-1:
                        out[i][j] = '<'
                    case ni == i && nj == j+1:
                        out[i][j] = '>'
                    }
                    break
                }
            }
        }
    }
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()
    for i := 0; i < H; i++ {
        w.WriteString(string(out[i]))
        w.WriteByte('\n')
    }
}

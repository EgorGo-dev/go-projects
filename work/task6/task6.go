package main

import "sort"

type Player struct {
    Name    string
    Goals   int
    Misses  int
    Assists int
    Rating  float64
}

func (p *Player) calculateRating() {
    if p.Misses == 0 {
        p.Rating = float64(p.Goals) + float64(p.Assists)/2
        return
    }
    p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
}

func NewPlayer(name string, goals, misses, assists int) Player {
    p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
    p.calculateRating()
    return p
}

func goalsSort(players []Player) []Player {
    res := make([]Player, len(players))
    copy(res, players)
    sort.Slice(res, func(i, j int) bool {
        if res[i].Goals != res[j].Goals {
            return res[i].Goals > res[j].Goals
        }
        return res[i].Name < res[j].Name
    })
    return res
}

func ratingSort(players []Player) []Player {
    res := make([]Player, len(players))
    copy(res, players)
    sort.Slice(res, func(i, j int) bool {
        if res[i].Rating != res[j].Rating {
            return res[i].Rating > res[j].Rating
        }
        return res[i].Name < res[j].Name
    })
    return res
}

func gmSort(players []Player) []Player {
    res := make([]Player, len(players))
    copy(res, players)
    sort.Slice(res, func(i, j int) bool {
        gmI := float64(res[i].Goals) / float64(res[i].Misses)
        if res[i].Misses == 0 {
            gmI = 1e9
        }
        gmJ := float64(res[j].Goals) / float64(res[j].Misses)
        if res[j].Misses == 0 {
            gmJ = 1e9
        }
        if gmI != gmJ {
            return gmI > gmJ
        }
        return res[i].Name < res[j].Name
    })
    return res
}
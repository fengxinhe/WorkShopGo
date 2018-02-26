package model

type Class struct {
    ClassID string
    ClassTitle string
    ClassSummary string
    FirstTag string
    SecondTag string
    // Content []byte
    // ImgUrl string
    // VideoUrl string
}

type Contest struct {
    ContestTitle string
    ContestSummary string
    //ImgUrl string
}

type IndexContent struct {
    Title string
    Classes []Class
    Contests []Contest
    Static string
}

func GetClass() *[]Class{
    classes := &[]Class{Class{ClassTitle: "cl1", ClassSummary: "aaaaa"},
                Class{ClassTitle: "cl2", ClassSummary: "bbbbb"},}
    return classes
}

func GetContest() *[]Contest{
    contests := &[]Contest{Contest{ContestTitle: "t1", ContestSummary: "ccccc"},
                Contest{ContestTitle: "t2", ContestSummary: "dddddd"},}
    return contests
}

func GetIndexContent() *IndexContent {
    content := &IndexContent{
        Title: "DUDU",
        Classes: []Class{Class{ClassTitle: "cl1", ClassSummary: "aaaaa"},
                    Class{ClassTitle: "cl2", ClassSummary: "bbbbb"},},
        Contests: []Contest{Contest{ContestTitle: "t1", ContestSummary: "ccccc"},
                    Contest{ContestTitle: "t2", ContestSummary: "dddddd"},},
    }
    return content
}

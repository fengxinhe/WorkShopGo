package model


type ClassTag struct{
    FirstTag string
    SecondTag string
}

type TagedClass struct {

}

func GetClassByTag(tag1 string, tag2 string) *[]Class{

    classes := &[]Class{Class{ClassTitle: "cl1", ClassSummary: "aaaaa"},
                Class{ClassTitle: "cl2", ClassSummary: "bbbbb"},}
    return classes
}

func GetClassByHeat() *[]Class{
    //class := from db
    
}

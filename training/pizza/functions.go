package main

import(
    "fmt"
)

var (
    ErrEmpieter             = fmt.Errorf("Ce slice va empieter sur un autre slice")
    ErrTaille               = fmt.Errorf("Ce slice n'a pas la bonne taille")
    ErrPasAssezIngredient   = fmt.Errorf("Pas assez d'ingredients")
    ErrTropDeCell           = fmt.Errorf("Ce slice est trop grand")
)

func (P *Pizza) to2DX(i int) (int){
    if(i > P.Rows*P.Columns){
        panic("i trop grand, pizza trop petite")
    }
    return i%P.Rows
}

func (P *Pizza) to2DY(i int) (int){
    if(i > P.Rows*P.Columns){
        panic("i trop grand, pizza trop petite")
    }
    return i/P.Rows
}

func (P *Pizza) to1D(x, y int) (int){
    if(x > P.Rows || y >P.Columns){
        panic("x ou y trop grand, pizza trop petite")
    }
    return x+y*P.Rows
}

func sliceIsValid(pizza *Pizza, slice *Slice) error{
    cell := pizza.Cell
    
    if slice.X1 < 0 || slice.X2 < 0 || slice.Y1 < 0 || slice.Y2 < 0 || slice.X1 >= pizza.Rows || slice.X2 >= pizza.Rows || slice.Y1 >= pizza.Columns || slice.Y2 >= pizza.Columns {
        return ErrTaille
    }
    
    cptT:=0
    cptM:=0
    
    for x:= slice.X1; x <= slice.X2 - slice.X1; x++ {
        for y:= slice.Y1; y <= slice.Y2 - slice.Y1; y++ {
            if(cell[pizza.to1D(x, y)].Used != nil){
                return ErrEmpieter
            }
            
            if(cell[pizza.to1D(x, y)].Ingredient == TOMATO){
                cptT++
            }else{
                cptM++
            }
        }
    }
    if(cptM < pizza.MinIngredient || cptT < pizza.MinIngredient){
        return ErrPasAssezIngredient
    }
    
    if((slice.X2 - slice.X1) * (slice.Y2 - slice.Y1) > pizza.MaxCells){
        return ErrTropDeCell
    }

    return nil
}

func placerSlice(pizza *Pizza, slice *Slice) bool{
    if(sliceIsValid(pizza, slice) != nil){
        return false
    }
    
    for x:= slice.X1; x <= slice.X2 - slice.X1; x++ {
        for y:= slice.Y1; y <= slice.Y2 - slice.Y1; y++ {
            pizza.Cell[pizza.to1D(x, y)].Used = slice
        }
    }

    return true
}
/*
func optimiserSliceXMoins(pizza *Pizza, slice *Slice) bool{
    
    
    
    
}

func optimiserSliceXPlus(pizza *Pizza, slice *Slice) bool{
    
    
    
    
}

func optimiserSliceYMoins(pizza *Pizza, slice *Slice) bool{
    
    
    
    
}

func optimiserSliceYPlus(pizza *Pizza, slice *Slice) bool{
    
    
    
    
}

//prend en paramètre un slice non valide, et essaye de le rendre valide en optimisant les slices autour
func optimiserSlice(pizza *Pizza, slice *Slice) bool{
    
    estOpti bool
    
    if slice.X1 > 0 {
        
        slice.X1 -= 1
        err := sliceIsValid(pizza, slice)
        
        if err == ErrEmpieter
        {
            estOpti = optimiserSliceXMoins(pizza, slice)
            if estOpti == true {
                return true
            }
        }
        slice.X1 += 1
        
    }
    if slice.Y1 > 0 {
        
        slice.Y1 -= 1
        err := sliceIsValid(pizza, slice)
        
        if err == ErrEmpieter
        {
            estOpti = optimiserSliceYMoins(pizza, slice)
            if estOpti == true {
                return true
            }
        }
        slice.Y1 += 1
        
    }
    if slice.X2 < (pizza.Columns - 1) {
        
        slice.X2 += 1
        err := sliceIsValid(pizza, slice)
        
        if err == ErrEmpieter
        {
            estOpti = optimiserSliceXPlus(pizza, slice)
            if estOpti == true {
                return true
            }
        }
        slice.X2 -= 1
        
    }
    if slice.Y2 < (pizza.Rows - 1) {
        
        slice.Y2 += 1
        err := sliceIsValid(pizza, slice)
        
        if err == ErrEmpieter
        {
            estOpti = optimiserSliceYPlus(pizza, slice)
            if estOpti == true {
                return true
            }
        }
        slice.Y2 -= 1
    }
    return false
}
*/

package main

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

func sliceIsValid(pizza *Pizza, slice *Slice) bool{
    cell := pizza.Cell
    
    cptT:=0
    cptM:=0
    
    for x:= slice.X1; x <= slice.X2 - slice.X1; x++ {
        for y:= slice.Y1; y <= slice.Y2 - slice.Y1; y++ {
            if(cell[pizza.to1D(x, y)].Used != nil){
                return false
            }
            
            if(cell[pizza.to1D(x, y)].Ingredient == TOMATO){
                cptT++
            }else{
                cptM++
            }
        }
    }
    if(cptM < pizza.MinIngredient || cptT < pizza.MinIngredient){
        return false
    }
    
    if((slice.X2 - slice.X1) * (slice.Y2 - slice.Y1) > pizza.MaxCells){
        return false
    }

    return true
}

func placerSlide(pizza *Pizza, slice *Slice) bool{
    if(sliceIsValid(pizza, slice) == false){
        return false
    }
    
    
    
    return true
}


package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, n int) (time int) {
    if (n <= 0) {
        n = 2
    }
    return len(layers) * n
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, j := range layers {
        if j == "noodles" {
            noodles = noodles + 50;
        } else if j == "sauce" {
            sauce = sauce + 0.2;
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string){
    ingredient := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = ingredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) (scaledQuantities []float64){
    for _, ingredient := range quantities {
            q_1 := ingredient / 2
            scaledQuantities = append(scaledQuantities, q_1 * float64(portion))
        }
    return
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.

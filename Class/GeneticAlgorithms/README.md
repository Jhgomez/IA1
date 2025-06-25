# Genetic Algorithms
In this example we use convergence, convergence is calculated whenever a neW generation is created this way we know if we have reached the limit and stop, when we stop we then choose the best member(that depends either if the problem we want to solve wants to maximize or minimize the values). Convergence is calculated using the fitness, and the fitness is calculated using the formula of we are given for example `f(x) = x2` or `f(x) = (a-b)+(2c-d)+5`, we can either do an average of the fitness values in each gen(like in this example, the fitness it their value itself, so we sum all items in the generation and divide it by the the members in the generation) or the total of the fitness values in each gen and then calculate convergence like this: `min(oldGenFitness, newGenFitness)/max(oldGenFitness, newGenFitness)`, calculating the convergence and then comparing it to a value we previously defined when calling the function can be one criteria to stop execution and choose the best candidate, or we can choose an arbitrary iteration limit and choose the best candidate in last iterations

In this algorithm we use the "hope"("esperanza" in spanish) to choose the fathers and the sons, whichever is more close to the desired number will be choosed

Once we have chosed the parents we should have defined already how the childs are going to be generated and how are they going to be selected/compete and between who are compeeting

In this example we define the next:

* Each generation will be 4 nodes, each will by any number from 0 or upto to 9999(inclusive)

* Two parents will be selected out of the four members in each generation, they are sorted from the most near to most far to the target number, and only the two most near items are selected selected as parents

* They will generate 5 childs
  * Child1: `(p1+p2)/2`

  * Child2: first identify the lowest and the highest value of the two parents, then `2*(parentWithHighestVal) - (parentWithLowestVal)`

  * Child3: `diference between parents` (note we have to get the absolute value of the diference)

  * Child4: `p1 * 1.1`, this is estrategic so we can avoid generating same values at some point(some sort of mutation)

  * Child5: `p1 * 0.9`, this is estrategic so we can avoid generating same values at some point(some sort of mutation)

* Next genration/replacement has to be generated in this order `p1, p2, c1, c2` where `c1` and `c2` can be any child that has the best fitness, in this case we use hope as the criteria, meaning the nearest two parents and nearest two childs will be choosed, and that means `p1` is closer to goal than `p2` and `h1` is closer than `h2`, the we calculate the convergence using their average and check if we reached the goal and exit or continue if goal not met yer


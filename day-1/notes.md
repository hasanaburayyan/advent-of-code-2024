# Notes


## Part One

For part one I simply just sort the lists, then iterate over them and calculate the difference between the two lists. I then sum up all the differences to get the total distance between the two lists. 

The trick here is to just ensure that the difference between each index uses absolute value before adding it to the sum.


## Part Two

I opted to use a frequency map to store the frequency of each number in the right list. I then iterate over the left list and multiply the number by the frequency of that number in the right list. I then sum up all the products to get the similarity score.


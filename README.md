# ffbp-neural-network
Creating a simple Feed-Forward-Back-Propagation Neural Network using Golang. This is in part to gain experience using Golang, and it part to investigate the advantage imposed by Golang's concurrency over Python-based AI.

## TODO

1. The first thing to do for this is to create a simple implementation of a Perceptron. This allows me first of all to relearn the basic implementation, and to practice doing so with Golang. 
2. Second, the objective is to expand this implementation so that I can have layers of perceptrons operating on each other. This completes the feed-forward portion of the algorithm.
3. Third, backpropagation needs to be implemented to complete the function of the neural net. 
4. If not done by this point, I should probably find a way to implement each perceptron/node as a struct extending an interface. This interface could then have methods for calculating linear combinations, sigmoid function values, and partial derivatives. This will (hypothetically) simplify the design of the algorithm and allow it to be extended more easily (i.e. instantiating layers of the network with various numbers of perceptrons).
5. Implementation of concurrency. I suspect that I will be changing things often enough and drastically enough that implementing concurrency as anything other than the last step would be a waste of time. Once everything else is done though, I think this would be a good way to speed up the training process.
6. Exploring use-cases. I want to actually be able to use this for something, so I should figure out what kind of problem I could solve with this. At the moment I'm thinking an image classifier would be an interesting application.
7. Publication. As a finishing touch (and a whole lot more work), it would be great to also create a web-facing application in Go from which this neural network can be seen. I'm thinking something along the lines of a single webpage with an image-upload form. The user uploads a given picture, and the neural net guesses what the picture contains (a house vs a tree or something). 

## Takeaways

I suspect that in completing this project in its entirety would be a great way to introduce myself to Go and explore a range of its functionalities. I am curious about the web development process with Go, and I wonder how well-suited it might be for AI development. This will theoretically answer some questions about each of those.

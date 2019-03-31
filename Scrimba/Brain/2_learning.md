# Scrimba > brain.js: 2. How they learn?

### I/O

Typically, nets just have their **inputs**, and do some **outputs**.

So in order to train them, for now, we just want to tell nets

1. What are we likely going to send to their **inputs**
2. and what we expect as an **output** on that inputs

### Neurons

It was proven, the best way to study is to start randomly. Each of the neurons is literally `Math.random()`

#### Activation

Popular activation nowadays is called `relu`, looking like this:

```js
const relu = value => (value < 0 ? 0 : value);
```

You may want to read [that](https://en.wikipedia.org/wiki/Activation_function) more intense.

## Layers

<p align="center">
<img src="https://cdn-images-1.medium.com/max/1600/1*-a-flCLHLCGM0-7TOcNJnQ.png" />
</p>

This is a neural net! Each circle is a **neuron**, each arrow is some maths. and the rows of neurons are called **layers**.

You maywant to know, that the completion layer is called an **output layer**, whereas every layer of studying is a **hidden layer**.

So if we would try to limitate our XOR net, giving it instead of what it used to be:

```js
const net = new brain.NeuralNetwork({ hiddenLayers: [3] });
```

Setting the hidden layer with only 1 neuron instead of 3, net would hit 20 000 errors without even solving the task.

Typically, the logic is `hiddenLayers: [x,y]` where **x** is number of neurons per layer and **y** is number of parallel layers

The deep math function net uses looks something like this:

```js
activate(inputWeights * inputs + biases);
```

You may see implication [here](https://github.com/BrainJS/brain.js/blob/9595fe1d0069939ba271b25c1e7db785edd11936/src/neural-network.js#L233)

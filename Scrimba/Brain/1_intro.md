# Srcimba > brain.js: Introduction

## What we will learn:

- Forward and backward propagation
- Layers and neurons
- Training and errors
- Feed forward networks
- Recurrent networks

And much moAR!

## What we will build:

- [XOR](/2_learning.md)
- Counter
- Basic math network
- Image recognizer
- Sentiment analyzer
- Children's book creator

Let's start!

## Let's write XOR

XOR is math object with such logics:

In: 0, 0; Out: 0

In: 0, 1; Out: 1

In: 1, 0; Out: 1

In: 1, 1; Out: 1

So let's do some js:

1. create our **training data** for network

```js
const trainingData = [
  { input: [0, 0], output: [0] },
  { input: [0, 1], output: [1] },
  { input: [1, 0], output: [1] },
  { input: [1, 1], output: [0] }
];
```

2. Before that we create our **network**:

```js
const net = new brain.NeuralNetwork({ hiddenLayers: [3] });
```

3. And in the end we **train** it:

```js
net.train(trainingData);
```

4. Now we can **run** some custom data in our `net` and ask it, what would be the output

```js
console.log(net.run([0, 0])); // Returns around ~0.058
```

We don't have `0` output, because nets are much more precise, they usually operate in the area, where you can round.

That's it. Cheers

### How we did that?

We used two methods: `train` and `run`

In `train` we used something called forward or backward propagation

#### What is it?

Imagine we have to goal the ball. The first step to that would be prediction or how we are going to throw a ball.

This prediction is **forward propagation**

When we throw a ball, we get the info about status back: it is either too far, too close or goaled. We learn by the information returned back to us. That is **backward propagation**

That is everything happening inside `net.train()`

Next we do `run`. On t hat stage we already now, how far are we from the goal. At this stage we don't need backward propagation. We are learned to throw that ball already!

### Errors

Unlike us, networks love the hard path of trials and mistakes. That's how we all learn, they are not explicit. Try this:

```js
net.train(trainingData, {
  log: error => console.log(error),
  logPeriod: 100
});
```

In my tests, it showed, `net` did **errors** around 4400 times, before understanding rules!

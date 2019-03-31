## Working with objects

How to send different shaped data into net?

If we have indexed array it may look like this and be pretty simple:

<p align="center">
<img src="/demo/1.png">
</p>

But what about Objects?

Let's try that:

Input: { red, green, blue }
Output: { light, neutral, dark}

Now let's do js out of it:

1. Set the I and O:

```js
const colors = [
  { green: 0.2, blue: 0.4 },
  { green: 0.4, blue: 0.6 },
  { red: 0.2, green: 0.8, blue: 0.8 },
  { green: 1, blue: 1 },
  { red: 0.8, green: 1, blue: 1 },
  { red: 1, green: 1, blue: 1 },
  { red: 1, green: 0.8, blue: 0.8 },
  { red: 1, green: 0.6, blue: 0.6 },
  { red: 1, green: 0.4, blue: 0.4 },
  { red: 1, green: 0.31, blue: 0.31 },
  { red: 0.8 },
  { red: 0.6, green: 0.2, blue: 0.2 }
];

const brightnesses = [
  { dark: 0.8 },
  { neutral: 0.8 },
  { light: 0.7 },
  { light: 0.8 },
  { light: 0.9 },
  { light: 1 },
  { light: 0.8 },
  { neutral: 0.7, light: 0.5 },
  { dark: 0.5, neutral: 0.5 },
  { dark: 0.6, neutral: 0.3 },
  { dark: 0.85 },
  { dark: 0.9 }
];
```

Don't be so scared, we are just assigning them to each other, so we feeding net a table like this:

| red | green | blue | dark | neutral | light |
| --- | ----- | ---- | ---- | ------- | ----- |
| 0   | 0.2   | 0.4  | 0.8  | 0       | 0     |
| 0   | 0.4   | 0.6  | 0    | 0.8     | 0     |
| 0.2 | 0.8   | 0.8  | 0    | 0       | 0.7   |
| 0   | 1     | 1    | 0    | 0       | 0.8   |
| 0.8 | 1     | 1    | 0    | 0       | 0.9   |
| 1   | 1     | 1    | 0    | 0       | 1     |
| 1   | 0.8   | 0.8  | 0    | 0       | 0.8   |
| 1   | 0.6   | 0.6  | 0    | 0.7     | 0.5   |
| 1   | 0.4   | 0.4  | 0.5  | 0.5     | 0     |
| 1   | 0.31  | 0.31 | 0.6  | 0.3     | 0     |
| 0.8 | 0     | 0    | 0.85 | 0       | 0     |
| 0.6 | 0.2   | 0.2  | 0.9  | 0       | 0     |

2. Now we parse it into `trainingData` in I/O format:

```js
const trainingData = [];

for (let i = 0; i < colors.length; i++)
  trainingData.push({ input: colors[i], output: brightnesses[i] });
```

3. and setup the network:

```js
const net = new brain.NeuralNetwork({ hiddenLayers: [3] });

const stats = net.train(trainingData);

console.log(stats);
```

That's it. Network is learned in 1500 errors.

Now we can ask it, what does it think of some color:

```js
console.log(
  net.run({
    red: 0.9
  })
);
```


Fun fact: if we want net to give us colors from params, we can just invert I and O:

```js
const invertedTrainingData = [];

for (let i = 0; i < colors.length; i++) {
    invertedTrainingData.push({
        input: brightnesses[i],
        output: colors[i]
    });
}

const invertedNet = new brain.NeuralNetwork({ hiddenLayers: [3] });

const invertedStats = invertedNet.train(invertedTrainingData);
```

and it fails, but that's not the point. The point is how they work
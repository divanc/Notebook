[< Previous abstract](React__9.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__11.md)
----------------------- | ----------------------------|-----------------------------

## 10. Lifting State Up

  Often, several components need to reflect the same changing data. We recommend lifting
  the shared state up to their closest common ancestor. Let’s see how this works in action.

Let's create a temperature calculator, starting from `BoilingVerdict` component, which says wether temperature is enough to boil water.

```js
function BoilingVerdict(props) {
  if (props.celsius >= 100) {
    return <p>The water will boil</p>;
  }
  return <p>The water will not boil</p>;
}
```

Next comes `Calculator` component:

```js
class Calculator extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      temperature: ''
    };

    this.handleChanges = this.handleChanges.bind(this);
  }

  handleChanges = event => this.setState({temperature: event.target.value});

  render() {
    const temperature = this.state.temperature
    return (
      <fieldset>
        <legend>Temperature in Celsius: </legend>
        <input value={temperature} onChange={this.handleChanges} />
        <BoilingVerdict celsius={parseFloat(temperature)}/>
      </fieldset>  
    );
  }
}
```

### Adding A Second Input

In addition to Celsius, let's add Fahrenheit, and keep them in sync

```js
//New
const scaleNames = {
  c: 'Celsius',
  f: 'Fahrenheit'
}
//
//Rename
class TemperatureInput extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      temperature: ''
    };

    this.handleChanges = this.handleChanges.bind(this);
  }

  handleChanges = event => this.setState({temperature: event.target.value});

  render() {
    const temperature = this.state.temperature
    //New
    const scale = this.props.scale;
    //
    return (
      <fieldset>
        <legend>Temperature in {scaleNames[scale]}: </legend>
        <input value={temperature} onChange={this.handleChanges} />
        <BoilingVerdict celsius={parseFloat(temperature)}/>
      </fieldset>  
    );
  }
}
```

```js
class Calculator extends React.Component {
  render() {
    return (
      <div>
        <TemperatureInput scale='c' />
        <TemperatureInput scale='f' />
      </div>
    );
  }
}
```

Having two inputs, yet these are not benefit each other with updating, let's fix it

### Writing Conversion Functions

First of all: two functions to convert temperature back and forth:

```js
function toCelsius = fahrenheit => (fahrenheit-32)*5/9;

function toFahrenheit = celsius => (celsius * 9/5) + 32;
```

And converter function

```js
function tryConvert(temperature, convert) {
  const input = parseFloat(temperature);
  if (Number.isNaN(input)) {
    return '';
  }
  const output = convert(input);
  const rounded = Math.round(output * 1000) / 1000;
  return rounded.toString();

  //tryConvert('abc', toCelsius)  >> empty string
}
```

### Lifting State Up

Currently, both `TemperatureInput` components independently keep their values in the local state. However, we want both of them be in sync. Celsius should reflect upon changes of Fahrenheit and vice-versa

  In React, sharing state is accomplished by moving it up to the closest common ancestor of the components that need it. This is called “lifting state up”. We will remove the local state from the TemperatureInput and move it into the Calculator instead.

  If the Calculator owns the shared state, it becomes the “source of truth” for the current temperature in both inputs. It can instruct them both to have values that are consistent with each other. Since the props of both TemperatureInput components are coming from the same parent Calculator component, the two inputs will always be in sync.

First, we want to replace `this.state.temperature` with `this.prop.temperature` in the `TemperatureInput`.
Lets prenend `this.state.temperature` already in `Calculator`.

```js
render() {
  // Before: const temperature = this.state.temperature;
  const temperature = this.props.temperature;
  // ...
```

We know props are *read-only*. When `temperature` was in local state, the `TemperatureInput` could just call `this.setState()` to change it, but now it has no control over it.

React decision > Controlled component. Just as usual DOM elements like `input` accept both `value` and `onChange` prop, so can the custom `TemperatureInput` accept `temperature` and `onTemperatureChange` props from parents as `Calculator`

Now, when the `TemperatureInput` wants to update its temperature, it calls `this.props.onTemperatureChange`:

```js
handleChange(event) {
  //Before: this.setState({temperature: event.target.value})
  this.props.onTemperatureChange(event.target.value);
}
```

That's our `TemperatureInput` component so far

```js
const scaleNames = {
  c: 'Celsius',
  f: 'Fahrenheit'
}

class TemperatureInput extends React.Component {
  constructor(props) {
    super(props);
    this.handleChanges = this.handleChanges.bind(this);
  }

  handleChanges = event => this.props.onTemperatureChange(event.target.value);

  render() {
    const temperature = this.props.temperature
    //New
    const scale = this.props.scale;
    //
    return (
      <fieldset>
        <legend>Temperature in {scaleNames[scale]}: </legend>
        <input value={temperature} onChange={this.handleChanges} />
        <BoilingVerdict celsius={parseFloat(temperature)}/>
      </fieldset>  
    );
  }
}
```

Talking about `Calculator`:

```js
class Calculator extends React.Component {
  constructor(props) {
    super(props);
    this.handleCelsiusChange = this.handleCelsiusChange.bind(this);
    this.handleFahrenheitChange = this.handleFahrenheitChange.bind(this);
    this.state = {temperature: '', scale: 'c'};
  }

  handleCelsiusChange = temperature => this.setState({scale: 'c',temperature})
  
  handleCelsiusChange = temperature => this.setState({scale: 'f',temperature})

  render() {
    const scale = this.state.scale;
    const temperature = this.state.temperature;
    const celsius = scale === 'f' ? tryConvert(temperature, toCelsius) : temperature;
    const fahrenheit = scale === 'c' ? tryConvert(temperature, toFahrenheit) : temperature;

    return (
      <div>
        <TemperatureInput scale='c' temperature={celsius} onTemperatureChange={this.handleCelsiusChange} />
        <TemperatureInput scale='f' temperature={fahrenheit} onTemperatureChange={this.handleFahrenheitChange} />
        
        <BoilingVerdict celsius={parseFloat(celsius)} />
      </div>
    );
  }
}
```

Let’s recap what happens when you edit an input:

* React calls the function specified as onChange on the DOM `<input>`. In our case, this is the handleChange method in the TemperatureInput component.
* The handleChange method in the TemperatureInput component calls this.props.onTemperatureChange() with the new desired value. Its props, including onTemperatureChange, were provided by its parent component, the Calculator.
* When it previously rendered, the Calculator has specified that onTemperatureChange of the Celsius TemperatureInput is the Calculator’s handleCelsiusChange method, and onTemperatureChange of the Fahrenheit TemperatureInput is the Calculator’s handleFahrenheitChange method. So either of these two Calculator methods gets called depending on which input we edited.
* Inside these methods, the Calculator component asks React to re-render itself by calling this.setState() with the new input value and the current scale of the input we just edited.
* React calls the Calculator component’s render method to learn what the UI should look like. The values of both inputs are recomputed based on the current temperature and the active scale. The temperature conversion is performed here.
* React calls the render methods of the individual TemperatureInput components with their new props specified by the Calculator. It learns what their UI should look like.
* React calls the render method of the BoilingVerdict component, passing the temperature in Celsius as its props.
* React DOM updates the DOM with the boiling verdict and to match the desired input values. The input we just edited receives its current value, and the other input is updated to the temperature after conversion.

Every update goes through the same steps so the inputs stay in sync.

### Lessons Learned

  Lifting state involves writing more “boilerplate” code than two-way binding approaches, but as a benefit, it takes less work to find and isolate bugs. Since any state “lives” in some component and that component alone can change it, the surface area for bugs is greatly reduced. Additionally, you can implement any custom logic to reject or transform user input.

  If something can be derived from either props or state, it probably shouldn’t be in the state. For example, instead of storing both celsiusValue and fahrenheitValue, we store just the last edited temperature and its scale. The value of the other input can always be calculated from them in the render() method. This lets us clear or apply rounding to the other field without losing any precision in the user input.

  When you see something wrong in the UI, you can use React Developer Tools to inspect the props and move up the tree until you find the component responsible for updating the state. 
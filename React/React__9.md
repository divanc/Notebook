
[< Previous abstract](React__8.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Next abstract >](React__10.md)
----------------------- | ----------------------------|-----------------------------

## 9. Forms

  HTML form elements work a little bit differently from other DOM elements in React, because form elements naturally keep some internal state. For example, this form in plain HTML accepts a single name:

```html
<form>
  <label>
    Name:
    <input type="text" name="name" />
  </label>
  <input type="submit" value="submit" />
</form>
```

It works in HTML in case you want user to submit and go to another page. If you want this in React — it works. However, in most cases JS functions are more convinient. Best way to do such in React called "controlled component".

### Controlled Components

HTML tags `<textarea>`,`<input>`,`<select>` usually maintain their on state depending on user input. In React we can control state of these only using `setState()`.

For example, if we want to make the previous example log the name when it is submitted, we can write the form as a controlled component:

```js
class NameForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange = event => this.setState({value: event.target.value});

  handleSubmit(event) {
    alert("A name was submitted: "+ this.state.value);
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Name:
          <input type="text" value={this.state.value} onChange={this.handleChange} />
        </label>
        <input type="submit" value="submit" />
      </form>
    );
  }
}
```

Since `value` is set on text input, the displayed value == `this.state.value`, `handleChange()` works every time user hits a key.

Within controlled component every state mutation has associated __handler function__. If we want, for example, user to input his name in all capitals we can make:

```js
handleChange(event) {
  this.setState({value: event.taget.value.toUpperCase()});
}
```

### Textarea tag

In HTML `<textarea>` defines its text as children:

```html
<textarea>
  Imma Child
</textarea>
```

In React it is but `value` instead. So it could be considered as single-line form

```js
class EssayForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: 'Please, write an essay about, why do you like me'
    };
    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
  }

  handleChange = event => this.setState({value: event.target.value});

  handleSubmit(event) {
    alert("Essay was submitted" + this.state.value);
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Essay:
          <textarea value={this.state.value} onChange={this.handleChange} />
        </label>
        <input type="submit" />
      </form>
    );
  }
}
```

### The select tag

HTML uses select like this:

```html
<select>
  <option value="grapefruit">Grapefruit</option>
  <option value="mango">Mango</option>
  <option selected value="yura">Yura</option>
</select>
```

React, instead of having `selected` attribute, uses `value` on the root `select` tag. This makes things easier for JSX.

```js
class FlavorForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: 'Yura'
    };

    this.handleSelect = this.handleSelect.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSelect = event => this.setState({value: event.target.value});

  handleSubmit(event) {
    alert(this.state.value + " was selected");
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Select:
          <select value={this.state.value} onChange={this.handleSelect}>
            <option>Mango</option>
            <option>Yura</option>
            <option>Coconut</option>
          </select>
        </label>
        <input type="submit" />
      </form>
    );
  }
}
```

#### Make your life easier, bro

You can pass an array into the value attribute, allowing you to select multiple options in a select tag:

```js
<select multiple={true} value={[Mango,Yura,Coconut]}>
```

### File Input Tag

Because its value is *read-only*, it is an **uncontrolled** component in React, we will discuss it a little bit later.

### Handling Multiple Inputs

In order to control several `input` elements, you can add names to these and control via `handleChanges()` and `event.target.name`.

```js
class Reservation extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isGoing: true,
      NumberOfGuests: 2
    };

    this.handleInputChange = this.handleInputChange.bind(this);
  }

  handleInputChange(event) {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState = ({[name]: value});
  }

  render () {
    return (
      <form>
        <label>
          Is Going:
          <input
            name="isGoing"
            type="checkbox"
            checked={this.state.isGoing}
            onChange={this.handleInputChange} />
        </label>
        <br />
        <label>
          Number Of Guests:
          <input
            name="NumberOfGuests"
            type="number"
            value={this.state.NumberOfGuests}
            onChange={this.handleInputChange} />
        </label>
      </form>
    );
  }
}
```

### Controlled Input Null Value

  Specifying the `value` prop on a controlled component prevents the user from changing the input unless you desire so. If you’ve specified a `value` but the input is still editable, you may have accidentally set value to `undefined` or `null`.


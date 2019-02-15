[< Previous abstract](React__11.md) | [Back To React Folder](https://github.com/Betra/Course-Abstract/tree/master/React) | [Hooks >](React__H1.md)
----------------------- | ----------------------------|-----------------------------

## Thinking In React

  React is, in our opinion, the premier way to build big, fast Web apps with JavaScript. It has scaled very well for us at Facebook and Instagram.

### Start With A Mock

Lets imagine we already have JSON API and mock design:

![God Design](https://reactjs.org/static/thinking-in-react-mock-1071fbcc9eed01fddc115b41e193ec11-4dd91.png)

JSON:

```json
[
  {category: "Sporting Goods", price: "$49.99", stocked: true, name: "Football"},
  {category: "Sporting Goods", price: "$9.99", stocked: true, name: "Baseball"},
  {category: "Sporting Goods", price: "$29.99", stocked: false, name: "Basketball"},
  {category: "Electronics", price: "$99.99", stocked: true, name: "iPod Touch"},
  {category: "Electronics", price: "$399.99", stocked: false, name: "iPhone 5"},
  {category: "Electronics", price: "$199.99", stocked: true, name: "Nexus 7"}
];
```

### Step 1: Break The UI Into Component Hierarchy

  The first thing you’ll want to do is to draw boxes around every component (and subcomponent) in the mock and give them all names. 

  But how do you know what should be its own component? Just use the same techniques for deciding if you should create a new function or object. One such technique is the single responsibility principle.

  Since you’re often displaying a JSON data model to a user, you’ll find that if your model was built correctly, your UI (and therefore your component structure) will map nicely. That’s because UI and data models tend to adhere to the same information architecture, which means the work of separating your UI into components is often trivial. Just break it up into components that represent exactly one piece of your data model.

![](https://reactjs.org/static/thinking-in-react-components-eb8bda25806a89ebdc838813bdfa3601-82965.png)

So we have 5 items:

1. **FilterableProductTable** (orange): *contains the entirety of the example*
2. **SearchBar** (blue): *receives all user input*
3. **ProductTable** (green): *displays and filters the data collection based on user input*
4. **ProductCategoryRow** (turquoise): *displays a heading for each category*
5. **ProductRow** (red): *displays a row for each product*

Or in parential way:

* **FilterableProductTable**
  * **SearchBar**
  * **ProductTable**
    * **ProductCategoryRow**
    * **ProductRow**
  
### Step 2: Build A Static Version in React

The best way is to build static ragdoll which renders UI at first. **Avoid state for now**.

### Step 3: Identify The Minimal (but complete) Representation Of UI State

  To build your app correctly, you first need to think of the minimal set of mutable state that your app needs. The key here is DRY: Don’t Repeat Yourself. Figure out the absolute minimal representation of the state your application needs and compute everything else you need on-demand. For example, if you’re building a TODO list, just keep an array of the TODO items around; don’t keep a separate state variable for the count. Instead, when you want to render the TODO count, simply take the length of the TODO items array.

Having:

* The original list of products
* The search text the user has entered
* The value of the checkbox
* The filtered list of products

To figure out which are state-appropriate ask 3 questions for each of these 5:

1. Is it passed in from a parent via props? If so, it probably isn’t state.
2. Does it remain unchanged over time? If so, it probably isn’t state.
3. Can you compute it based on any other state or props in your component? If so, it isn’t state.

So these are state:

* The search text the user has entered
* The value of the checkbox

### Step 4: Identify Where Your State Should Live

  OK, so we’ve identified what the minimal set of app state is. Next, we need to identify which component mutates, or owns, this state.

  Remember: React is all about one-way data flow down the component hierarchy. It may not be immediately clear which component should own what state. This is often the most challenging part for newcomers to understand, so follow these steps to figure it out:

For each state go through:

* Identify every component that renders something based on that state.
* Find a common owner component (a single component above all the components that need the state in the hierarchy).
* Either the common owner or another component higher up in the hierarchy should own the state.
* If you can’t find a component where it makes sense to own the state, create a new component simply for holding the state and add it somewhere in the hierarchy above the common owner component.

### Step 5: Add Inverse Data Flow

Goto: Lesson 10


And that's React Basics!
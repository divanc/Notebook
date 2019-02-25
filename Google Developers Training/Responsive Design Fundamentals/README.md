[Back to GDT Folder](https://github.com/Betra/Course-Abstract/tree/master/Google%20Developers%20Training)

# Responsive Web Design Fundamentals

1. [Lesson 1: Sites on Mobile](#1)
2. [Lesson 2: Digging in](#2)
3. [Lesson 3: Page Layout](#3)
4. [Lesson 4: Common Responsive Patterns](#4)
5. [Lesson 5: Optimization](#5)

<a name="1"></a>

# Lesson 1: Sites on Mobile

  RWD is art, not the science

Wanna get better — work more. There  is no right solution

This course is aimed to set thinkinkg in responsive web design.

Approaching RWD systematically increases speed of project comletion, as an afterthought it slows down

### Pan, Zoom, Touch, Ick

It is pain, I know. If users are gonna do that, you gonna lose them. That's how it works. Both on mobile and TV.

<a name="2"></a>

# Lesson 2: Digging in

## viewport

Viewport is an area to which browser can render content to.

Instead of reporting viewport in pxls, browser reports it in **device independent pixels** dips. Ignores Pixel density.

By default it is always set 980pxls. Not every device have that at all. Then it boosts the fonts, in order to save situation.

Of course, browser makes it worse, what did you expect?

Adding this, we manage problem to ourselves, browser is just a procrastinator.

```html
<meta type="viewport" content="width=device-width,initial-scale=1">
```

## CSS care

* CSS allow content to overflow blocks so things still can move beyond. Easy fix is `max-width`

```css
img {
  max-width: 100%;
}
```

* Buttons should be at least 48x48, as our fingers are much bigger, than mouse cursor

```css
button {
  width: 48px; /* BAD; Shame on you */
  size: 48px; /* BAD; Shame on you */
  
  min-width: 48px; /* Omg, He did it! */
  min-height: 48px; /*Yesssss*/
}
```

## Design small

You probably should start your design from the phone and gradually enlarge.

Starting small have to really prioritize what is really important for a user.

\#Perf Matters. Also helps to think how much data loads into a page.

### Tests

Let's examine how this string works:

```css
.top-news__item a {
 padding: 1.5em inherit;
}
```

Simply, up-down padding would give `1.5em` and left-right would give a parent value;

<a name="3"></a>

# Lesson 3: Page Layout

Your site should be different on different types of devices. The easiest way to control is with media queries

## @media

### HTML Way

```html
<link rel="stylesheet" media="screen and (min-width: 300px)" href="pattern.css"
```

That way we can add different CSSes for different sizes

### CSS Way

```css
@media screen and (min-width: 500px) {
  body {
    etc..
  }
}
```

or 

```css
@import url("no.css") only screen and (min-width: 500px);
```

Last one kills performance.

### Usage

Usually we want to use `min-width` or `max-width`

`max-width` case is when we want to activate @media *when screen is less, than amount specified*. And vice-versa

## Breakpoints

Place, where page changes layout is called **breakpoint**. You probably want to use many of those.

So how to pick those?

* Just do as small as you can get it
* Then stretch slowly looking, where spaces between re getting to big, or any solution may be applied
* create breakpoint
* Go on
* Goto()

So far we've used only simple media queries. if we want to, we can complicate things

```css
@media screen and (min-width:500px) and (max-width:600px) {

}
```

# Grids (horaaaaaaaaaay!)

But not CSS grids, sorry

# Flexbox

```html
<div>
  <div class="a"></div>
  <div class="b"></div>
  <div class="c"></div>
</div>
```

In normal layout this would give one-under-another positioning

```css
div:first-of-Type {
  display: flex; /* Changes everything*/
}
```

Now they are shown in a row, because that's standard flex-direction, fitting on a single line.

We can change that adding `flex-wrap: wrap`, which tells browser, it is ok to move to a next line.

```css
div:first-of-Type {
  display: flex; /* Changes everything*/
  flex-wrap: wrap;
}
```

Also we can change order of elements:

```css
@media screen and (min-width: 700px) {
  .a {  order: 3; }
  .b {  order: 1; }
  .c {  order: 2; }
}
```

<a name="4"></a>

# Lesson 4: Common Responsive Patterns

There are 4 of most common: 

* Fluid
* Column Drop
* Layout shifter
* Off canvas

## [Column Drop](https://jsfiddle.net/mbf6wg7s/12/)

At it's narrowest viewport, each element stacks vertically.

Becoming wider elemets expand until they hit first breakpoint. On breakpoint hit they are stacked by two in a row now.

And three.. And four... You get it

When viewport hits max width, we add margins on left & right. 

Let's code this

```html
<main class="container">
  <div class="box a"></div>
  <div class="box b"></div>
  <div class="box c"></div>
</main>
```

```css
.container {
  display: flex;
  flex-wrap: wrap;
  background: #bbb;
}

/*  For small screen they are 100%  */

.box { width: 100%; }

.a { background: orange; }

.b { background: blue; }

.c { background: magenta; }

/*  That's when we want to start stacking by two */

@media screen and (min-width: 600px) {
  div:first-of-type { width: 25%; }
  .b { width: 75%; }
}

@media screen and (min-width: 700px) {
  div:first-of-type,
  div:last-of-type { width: 25%; }
  .b { width: 50%; }
}
```

## [Mostly Fluid](https://jsfiddle.net/divance/egzuqh7k/15/)

Just like before, at narrow is stacked, but with more complex filling in

```html
<main class="container">
  <div class="box a"></div>
  <div class="box b"></div>
  <div class="box c"></div>
  <div class="box d"></div>
  <div class="box e"></div>
  
</main>
```

```css
.container {
  display: flex;
  flex-wrap: wrap;
  background: #bbb;
}

.box { width: 100%; }
.a { background: orange; }
.b { background: blue; }
.c { background: magenta; }
.d { background: violet; }
.e { background: purple; }

@media screen and (min-width: 500px) {
  .b { width: 50%; }
  .c { width: 50%; }
}

@media screen and (min-width:600px) {
  .a { width: 50%; }
  .b { width: 50%; }
  .c,.d,.e { width: 33.33%; }
}

@media screen and (min-width: 700px) {
  .container {
    width: 700px;
    margin: 0 auto;
  }
}
```

## [Layout Shifter](https://jsfiddle.net/divance/0w2b6Lks/14/)

Instead of reflowing, content moves

```html
<main>
  <div class="a"> </div>
  <main id="main2">
    <div class="b"> </div>
    <div class="c"> </div>
  </main>
  <div class="d"> </div>
</main>
```

```css
main {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
}

.box { width: 100%; }

.a { background: green; }
.b { background: darkblue; }
.c { background: red; }
.d { background: violet; }

@media screen and (min-width: 500px) {
  .a { width: 50%; }
  #main2 { width: 50%; }
}

@media screen and (min-width: 600px) {
  .a { width: 25%; order: 1; }
  #main2 { width: 50%; }
  .d { width: 25%; order: -1; }
}
```

## [Off Canvas](https://jsfiddle.net/divance/cok1059z/8/)

Turning off extra stuff on small screens

```html
  <body>
    <nav id="drawer">Menu</nav>
    <main>Not Menu</main>
  </body>
```

```css
html,
body,
main {
  width: 100%;
}

nav {
  width: 300px;
  height: 100%;
  position: absolute;
  transform: translate(-300px, 0);
  transition: transform 0.3s ease;
}

nav.open {
  transform: translate(100, 0);
}

@media screen and (min-width:600px) {
  nav {
    position: relative;
    transform: translate(0, 0);
  }

  body {
    display: flex;
    flex-flow: row nowrap;
  }

  main {
    width: auto;
    flex-grow: 1;
  }
}

```

<a name="5"></a>

## Lesson 5: Optimization

### Images

Images are heavy, they take 65% of every site

Wanna study? Not on this course, sweetheart

### Tables

There's a chance they are going to overflow the viewport

* Hidden columns
* No more tables
* Contained tables

#### Hidden columns 

Typically set the classes for less important info and hide it unless the monitor is big enough

#### No more tables

Under small screen we don't want our tables to act like tables `display: block` and program those

#### Contained tables

We can wrap it in div and make only div to hav horizontal overflow.

```css
div.withATable {
  width: 100%
  overflow-x: auto;
}
```

### Fonts

Too short or long lines are bad. 65 chars on the line seems to be ideal for the web.

Make sure fonts are big enough to read on smalls

### Minor Breakpoints

Just small changes: margins, font-size, icon size